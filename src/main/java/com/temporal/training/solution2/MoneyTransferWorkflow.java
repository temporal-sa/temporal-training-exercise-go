package com.temporal.training.solution2;

import io.temporal.workflow.WorkflowInterface;
import io.temporal.workflow.WorkflowMethod;
import io.temporal.workflow.SignalMethod;

@WorkflowInterface
public interface MoneyTransferWorkflow {
    
    @WorkflowMethod
    String transfer(TransferRequest request);
    
    @SignalMethod
    void approve(boolean approved);
}