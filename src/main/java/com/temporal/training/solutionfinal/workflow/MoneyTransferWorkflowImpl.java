package com.temporal.training.solutionfinal.workflow;

import io.temporal.activity.ActivityOptions;
import io.temporal.common.RetryOptions;
import io.temporal.common.SearchAttributeKey;
import io.temporal.workflow.Async;
import io.temporal.workflow.Workflow;
import org.slf4j.Logger;
import java.time.Duration;

import com.temporal.training.solutionfinal.activity.BankingActivities;
import com.temporal.training.solutionfinal.model.TransferRequest;
import com.temporal.training.solutionfinal.model.TransferStatus;

public class MoneyTransferWorkflowImpl implements MoneyTransferWorkflow {
    
    private static final Logger logger = Workflow.getLogger(MoneyTransferWorkflowImpl.class);
    
    private final BankingActivities activities = Workflow.newActivityStub(
        BankingActivities.class,
        ActivityOptions.newBuilder()
            .setStartToCloseTimeout(Duration.ofSeconds(5))
            .setRetryOptions(RetryOptions.newBuilder()
                .setMaximumAttempts(3)
                .build())
            .build()
    );
    
    private TransferStatus status = TransferStatus.PENDING;
    private TransferRequest request;
    private boolean approved = false;
    private boolean rejected = false;
    
    @Override
    public void transfer(TransferRequest request) {
        this.request = request;
        logger.info("Starting transfer: {}", request);
        
        // Set search attributes
        Workflow.upsertTypedSearchAttributes(
            SearchAttributeKey.forKeyword("AccountId").valueSet(request.fromAccount())
        );
        
        // Wait for approval with timeout
        logger.info("Waiting for approval...");
        Workflow.await(Duration.ofMinutes(10), () -> approved || rejected);
        
        if (rejected) {
            logger.info("Transfer rejected");
            status = TransferStatus.CANCELLED;
            return;
        }
        
        if (!approved) {
            logger.info("Transfer timed out without approval");
            status = TransferStatus.FAILED;
            return;
        }
        
        status = TransferStatus.APPROVED;
        logger.info("Transfer approved, executing...");
        
        try {
            // Execute transfer
            activities.withdraw(request.fromAccount(), request.amount());
            activities.deposit(request.toAccount(), request.amount());
            status = TransferStatus.COMPLETED;
            logger.info("Transfer completed successfully");
        } catch (Exception e) {
            // Compensate on failure
            logger.info("Transfer failed, initiating compensation");
            Async.procedure(activities::refund, request.fromAccount(), request.amount());
            status = TransferStatus.FAILED;
            throw e;
        }
    }
    
    @Override
    public void approve() {
        logger.info("Approval signal received");
        approved = true;
    }
    
    @Override
    public void reject() {
        logger.info("Rejection signal received");
        rejected = true;
    }
    
    @Override
    public TransferStatus getStatus() {
        return status;
    }
    

}
