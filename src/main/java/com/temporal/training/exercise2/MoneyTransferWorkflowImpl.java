package com.temporal.training.exercise2;

import io.temporal.activity.ActivityOptions;
import io.temporal.workflow.Workflow;

import java.time.Duration;

/**
 * Money transfer workflow implementation.
 * TODO: Implement the complete transfer workflow with approval mechanism.
 */
public class MoneyTransferWorkflowImpl implements MoneyTransferWorkflow {
    
    /**
     * Activity stub for banking operations.
     * TODO: Create activity stub with proper options.
     */
    private final BankingActivities activities = null; // TODO: Implement this
    
    /**
     * Flag to track approval status.
     * TODO: Use this for approval mechanism.
     */
    private boolean approved = false;
    
    @Override
    public String transfer(TransferRequest request) {
        // TODO: Implement the transfer workflow:
        // 1. Withdraw from source account
        // 2. Wait for approval signal
        // 3. If approved: deposit to target account
        // 4. If not approved: refund to source account
        // 5. Return appropriate status
        
        return null; // Replace with actual implementation
    }
    
    @Override
    public void approve(boolean approved) {
        // TODO: Implement approval signal handler
        // Set the approved flag
    }
}