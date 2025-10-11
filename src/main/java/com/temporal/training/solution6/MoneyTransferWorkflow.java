package com.temporal.training.solution6;

import io.temporal.workflow.WorkflowInterface;
import io.temporal.workflow.WorkflowMethod;
import io.temporal.workflow.SignalMethod;
import io.temporal.workflow.QueryMethod;

@WorkflowInterface
public interface MoneyTransferWorkflow {
    
    @WorkflowMethod
    String transfer(TransferRequest request);
    
    @SignalMethod
    void approve(boolean approved);
    
    @QueryMethod
    TransferStatus getStatus();
}