package com.temporal.training.solutionfinal;

import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowOptions;
import io.temporal.serviceclient.WorkflowServiceStubs;
import java.util.UUID;

import com.temporal.training.solutionfinal.model.TransferRequest;
import com.temporal.training.solutionfinal.workflow.MoneyTransferWorkflow;

public class StartWorkflow {
    
    public static void main(String[] args) {
        WorkflowServiceStubs service = WorkflowServiceStubs.newLocalServiceStubs();
        WorkflowClient client = WorkflowClient.newInstance(service);
        
        String workflowId = "transfer-" + UUID.randomUUID();
        
        MoneyTransferWorkflow workflow = client.newWorkflowStub(
            MoneyTransferWorkflow.class,
            WorkflowOptions.newBuilder()
                .setWorkflowId(workflowId)
                .setTaskQueue(StartWorker.TASK_QUEUE)
                .build()
        );
        
        TransferRequest request = new TransferRequest(
            "account-123",
            "account-456", 
            100.0,
            UUID.randomUUID().toString()
        );
        
        // Start workflow
        WorkflowClient.start(workflow::transfer, request);
        System.out.println("Started transfer workflow: " + workflowId);
        
        // Query status
        System.out.println("Status: " + workflow.getStatus());
        
        // Approve transfer
        workflow.approve();
        System.out.println("Transfer approved");
    }
}
