package com.temporal.training.exercise4;

import io.temporal.activity.ActivityOptions;
import io.temporal.workflow.Workflow;
import org.slf4j.Logger;

import java.time.Duration;
// TODO: Import SearchAttributeKey for custom search attributes

public class MoneyTransferWorkflowImpl implements MoneyTransferWorkflow {
    
    private static final Logger logger = Workflow.getLogger(MoneyTransferWorkflowImpl.class);
    
    private final BankingActivities activities = Workflow.newActivityStub(
        BankingActivities.class,
        ActivityOptions.newBuilder()
            .setStartToCloseTimeout(Duration.ofSeconds(30))
            .build()
    );
    
    private TransferStatus status = TransferStatus.PENDING;
    private TransferRequest request;
    private boolean approved = false;
    private boolean approvalReceived = false;
    
    @Override
    public String transfer(TransferRequest request) {
        this.request = request;
        logger.info("Starting transfer: {}", request);
        
        // TODO: Set search attributes for filtering
        // Use Workflow.upsertTypedSearchAttributes() with SearchAttributeKey.forKeyword("AccountId")
        // Set the value to request.fromAccount()
        
        try {
            // Step 1: Withdraw from source account
            activities.withdraw(request.fromAccount(), request.amount());
            logger.info("Withdrawal completed");
            
            // Step 2: Wait for approval
            logger.info("Waiting for approval...");
            Workflow.await(() -> approvalReceived);
            
            if (approved) {
                // Step 3a: Approved - deposit to target account
                status = TransferStatus.APPROVED;
                activities.deposit(request.toAccount(), request.amount());
                status = TransferStatus.COMPLETED;
                logger.info("Transfer approved and completed");
                return "Transfer completed successfully";
            } else {
                // Step 3b: Not approved - refund to source account
                activities.refund(request.fromAccount(), request.amount());
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