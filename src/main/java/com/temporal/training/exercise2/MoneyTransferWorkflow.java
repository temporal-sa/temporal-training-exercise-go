package com.temporal.training.exercise2;

import io.temporal.workflow.WorkflowInterface;
import io.temporal.workflow.WorkflowMethod;
import io.temporal.workflow.SignalMethod;

/**
 * Money transfer workflow interface.
 * TODO: Add signal method for approval.
 */
@WorkflowInterface
public interface MoneyTransferWorkflow {
    
    /**
     * Main workflow method to process money transfer.
     * TODO: Implement the workflow logic.
     * 
     * @param request The transfer request details
     * @return Transfer result or status
     */
    @WorkflowMethod
    String transfer(TransferRequest request);
    
    /**
     * Signal method for approving the transfer.
     * TODO: Add @SignalMethod annotation and implement.
     * 
     * @param approved Whether the transfer is approved
     */
    // TODO: Add annotation
    void approve(boolean approved);
}