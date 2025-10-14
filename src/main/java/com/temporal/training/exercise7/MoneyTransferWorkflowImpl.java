package com.temporal.training.exercise7;

import io.temporal.activity.ActivityOptions;
import io.temporal.common.SearchAttributeKey;
import io.temporal.workflow.Workflow;
import org.slf4j.Logger;

import java.time.Duration;

public class MoneyTransferWorkflowImpl implements MoneyTransferWorkflow {
    
    private static final Logger logger = Workflow.getLogger(MoneyTransferWorkflowImpl.class);
    
    private final BankingActivities activities = Workflow.newActivityStub(
        BankingActivities.class,
        ActivityOptions.newBuilder()
            .setStartToCloseTimeout(Duration.ofSeconds(5))
            .build()
    );
    
    private TransferStatus status = TransferStatus.PENDING;
    private boolean approved = false;
    private boolean approvalReceived = false;
    private boolean retryRequested = false;
    private TransferRequest updatedRequest;
    
    @Override
    public String transfer(TransferRequest request) {
        logger.info("Starting transfer: {}", request);
        
        Workflow.upsertTypedSearchAttributes(
            SearchAttributeKey.forKeyword("AccountId").valueSet(request.fromAccount())
        );
        
        try {
            updatedRequest = request;
            
            // TODO: Step 1: use executeWithRetry to run Withdraw with manual retry capability
            
            // Step 2: Wait for approval
            logger.info("Waiting for approval...");
            Workflow.await(() -> approvalReceived);
            
            if (approved) {
                status = TransferStatus.APPROVED;
                // TODO: Step 3a: Deposit with manual retry capability
                status = TransferStatus.COMPLETED;
                logger.info("Transfer completed successfully");
                return "Transfer completed successfully";
            } else {
                // Step 3b: Refund
                activities.refund(updatedRequest.fromAccount(), updatedRequest.amount());
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
    
    private void executeWithRetry(Runnable operation) {
        while (true) {
            try {
                operation.run();
                break;
            } catch (Exception e) {
                status = TransferStatus.RETRYING;
                retryRequested = false;
                
                Workflow.await(() -> retryRequested);
            }
        }
    }
    
    @Override
    public void approve(boolean approved) {
        logger.info("Approval received: {}", approved);
        this.approved = approved;
        this.approvalReceived = true;
    }
    
    @Override
    public void retry(RetryUpdate update) {
        // TODO: Implement retry signal handler that updates the request based on the key/value
        // Handle: "fromAccount", "toAccount", "amount"
    }
    
    @Override
    public TransferStatus getStatus() {
        return status;
    }
}
