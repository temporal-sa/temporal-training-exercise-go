package com.temporal.training.solution7;

import io.temporal.workflow.QueryMethod;
import io.temporal.workflow.SignalMethod;
import io.temporal.workflow.WorkflowInterface;
import io.temporal.workflow.WorkflowMethod;

@WorkflowInterface
public interface MoneyTransferWorkflow {
    
    @WorkflowMethod
    String transfer(TransferRequest request);
    
    @SignalMethod
    void approve(boolean approved);
    
    @SignalMethod
    void retry(RetryUpdate update);
    
    @QueryMethod
    TransferStatus getStatus();
}
