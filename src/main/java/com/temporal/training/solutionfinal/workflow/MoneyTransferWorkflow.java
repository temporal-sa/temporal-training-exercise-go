package com.temporal.training.solutionfinal.workflow;

import com.temporal.training.solutionfinal.model.TransferRequest;
import com.temporal.training.solutionfinal.model.TransferStatus;

import io.temporal.workflow.QueryMethod;
import io.temporal.workflow.SignalMethod;
import io.temporal.workflow.WorkflowInterface;
import io.temporal.workflow.WorkflowMethod;

@WorkflowInterface
public interface MoneyTransferWorkflow {
    
    @WorkflowMethod
    void transfer(TransferRequest request);
    
    @SignalMethod
    void approve();
    
    @SignalMethod
    void reject();
    
    @QueryMethod
    TransferStatus getStatus();
    

}
