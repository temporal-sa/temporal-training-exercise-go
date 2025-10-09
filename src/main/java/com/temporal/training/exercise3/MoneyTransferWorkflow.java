package com.temporal.training.exercise3;

import io.temporal.workflow.WorkflowInterface;
import io.temporal.workflow.WorkflowMethod;
import io.temporal.workflow.SignalMethod;
import io.temporal.workflow.QueryMethod;

/**
 * Money transfer workflow interface with query support.
 * TODO: Add @QueryMethod annotations for status queries.
 */
@WorkflowInterface
public interface MoneyTransferWorkflow {
    
    @WorkflowMethod
    String transfer(TransferRequest request);
    
    @SignalMethod
    void approve(boolean approved);
    
    /**
     * Query method to get current transfer status.
     * TODO: Add @QueryMethod annotation.
     */
    // TODO: Add annotation
    TransferStatus getStatus();
    

}