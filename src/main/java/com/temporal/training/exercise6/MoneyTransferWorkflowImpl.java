package com.temporal.training.exercise6;

import io.temporal.activity.ActivityOptions;
import io.temporal.common.SearchAttributeKey;
import io.temporal.workflow.Workflow;
import org.slf4j.Logger;

import java.time.Duration;

public class MoneyTransferWorkflowImpl implements MoneyTransferWorkflow {
    
    private static final Logger logger = Workflow.getLogger(MoneyTransferWorkflowImpl.class);
    
    private static class ActivityFactory {
        private static BankingActivities createActivity(String summary) {
            return Workflow.newActivityStub(
                BankingActivities.class,
                ActivityOptions.newBuilder()
                    .setStartToCloseTimeout(Duration.ofSeconds(5))
                    .setSummary(summary)
                    .build()
            );
        }
        
        static BankingActivities createWithdrawActivity(String accountId) {
            return createActivity("Withdrawing funds from account " + accountId);
        }
        
        static BankingActivities createDepositActivity(String accountId) {
            return createActivity("Depositing funds to account " + accountId);
        }
        
        static BankingActivities createRefundActivity(String accountId) {
            return createActivity("Refunding funds to account " + accountId);
        }
    }
    
    private TransferStatus status = TransferStatus.PENDING;
    private boolean approved = false;
    private boolean approvalReceived = false;
    
    @Override
    public String transfer(TransferRequest request) {
        logger.info("Starting transfer: {}", request);
        
        // Set search attributes for filtering
        Workflow.upsertTypedSearchAttributes(
            SearchAttributeKey.forKeyword("AccountId").valueSet(request.fromAccount())
        );
        
        try {
            // Step 1: Withdraw from source account
            ActivityFactory.createWithdrawActivity(request.fromAccount()).withdraw(request.fromAccount(), request.amount());
            logger.info("Withdrawal completed");
            
            // Step 2: Wait for approval
            logger.info("Waiting for approval...");
            Workflow.await(() -> approvalReceived);
            
            if (approved) {
                // Step 3a: Approved - deposit to target account
                status = TransferStatus.APPROVED;
                ActivityFactory.createDepositActivity(request.toAccount()).deposit(request.toAccount(), request.amount());
                status = TransferStatus.COMPLETED;
                logger.info("Transfer approved and completed");
                return "Transfer completed successfully";
            } else {
                // Step 3b: Not approved - refund to source account
                ActivityFactory.createRefundActivity(request.fromAccount()).refund(request.fromAccount(), request.amount());
                status = TransferStatus.CANCELLED;
                logger.info("Transfer rejected and refunded");
                return "Transfer rejected and refunded";
            }
        } catch (Exception e) {
            status = TransferStatus.FAILED;
            logger.error("Transfer failed: {}", e.getMessage());
            throw e;
        }
    }
    
    @Override
    public void approve(boolean approved) {
        logger.info("Approval received: {}", approved);
        this.approved = approved;
        this.approvalReceived = true;
    }
    
    @Override
    public TransferStatus getStatus() {
        return status;
    }
}