package com.temporal.training.exercise3;

import io.temporal.activity.ActivityOptions;
import io.temporal.workflow.Workflow;
import org.slf4j.Logger;

import java.time.Duration;

/**
 * Money transfer workflow implementation with query support.
 * TODO: Implement query handlers and status tracking.
 */
public class MoneyTransferWorkflowImpl implements MoneyTransferWorkflow {
    
    private static final Logger logger = Workflow.getLogger(MoneyTransferWorkflowImpl.class);
    
    private final BankingActivities activities = Workflow.newActivityStub(
        BankingActivities.class,
        ActivityOptions.newBuilder()
            .setStartToCloseTimeout(Duration.ofSeconds(5))
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
        
        // Step 1: Withdraw from source account
        activities.withdraw(request.fromAccount(), request.amount());
        logger.info("Withdrawal completed");
        
        // Step 2: Wait for approval
        logger.info("Waiting for approval...");
        Workflow.await(() -> approvalReceived);
        
        if (approved) {
            // Step 3a: Approved - deposit to target account
            // TODO: Update status to APPROVED
            activities.deposit(request.toAccount(), request.amount());
            // TODO: Update status to COMPLETED
            logger.info("Transfer approved and completed");
            return "Transfer completed successfully";
        } else {
            // Step 3b: Not approved - refund to source account
            activities.refund(request.fromAccount(), request.amount());
            // TODO: Update status to CANCELLED
            logger.info("Transfer rejected and refunded");
            return "Transfer rejected and refunded";
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
        // TODO: Return current transfer status
        return null;
    }
}
