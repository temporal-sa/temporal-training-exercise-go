package com.temporal.training.exercise3;

import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowOptions;
import io.temporal.serviceclient.WorkflowServiceStubs;

import java.util.concurrent.CompletableFuture;

public class StartWorkflow {
    
    public static void main(String[] args) throws Exception {
        WorkflowServiceStubs service = WorkflowServiceStubs.newLocalServiceStubs();
        WorkflowClient client = WorkflowClient.newInstance(service);
        
        WorkflowOptions options = WorkflowOptions.newBuilder()
            .setTaskQueue(StartWorker.TASK_QUEUE)
            .setWorkflowId("money-transfer-" + System.currentTimeMillis())
            .build();
        
        MoneyTransferWorkflow workflow = client.newWorkflowStub(MoneyTransferWorkflow.class, options);
        
        TransferRequest request = new TransferRequest(
            "account-123",
            "account-456", 
            100.0,
            "transfer-" + System.currentTimeMillis()
        );
        
        System.out.println("Starting money transfer workflow...");
        CompletableFuture<String> result = WorkflowClient.execute(workflow::transfer, request);
        
        // TODO: Query status before approval
        // System.out.println("Current status: " + workflow.getStatus());
        
        // Wait a bit, then send approval
        Thread.sleep(3000);
        System.out.println("Sending approval signal...");
        workflow.approve(true);
        
        // TODO: Query status after approval
        // System.out.println("Status after approval: " + workflow.getStatus());
        
        String transferResult = result.get();
        System.out.println("Transfer result: " + transferResult);
        
        // TODO: Query final status
        // System.out.println("Final status: " + workflow.getStatus());
    }
}