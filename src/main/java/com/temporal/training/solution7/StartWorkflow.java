package com.temporal.training.solution7;

import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowOptions;
import io.temporal.serviceclient.WorkflowServiceStubs;

import java.util.UUID;

public class StartWorkflow {
    
    public static void main(String[] args) {
        WorkflowServiceStubs service = WorkflowServiceStubs.newLocalServiceStubs();
        WorkflowClient client = WorkflowClient.newInstance(service);
        
        String workflowId = "money-transfer-" + UUID.randomUUID();
        
        MoneyTransferWorkflow workflow = client.newWorkflowStub(
            MoneyTransferWorkflow.class,
            WorkflowOptions.newBuilder()
                .setWorkflowId(workflowId)
                .setTaskQueue("money-transfer-task-queue")
                .build()
        );
        
        TransferRequest request = new TransferRequest("account-123", "invalid-account-456", 100.0, UUID.randomUUID().toString());
        
        WorkflowClient.start(workflow::transfer, request);
    }
}
