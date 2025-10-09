package com.temporal.training.exercise1;

import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowOptions;
import io.temporal.serviceclient.WorkflowServiceStubs;

/**
 * Starter class creates and executes a workflow.
 * This is the client that starts workflow executions.
 */
public class StartWorkflow {
    public static void main(String[] args) {
        // Connect to Temporal service
        WorkflowServiceStubs service = WorkflowServiceStubs.newLocalServiceStubs();
        WorkflowClient client = WorkflowClient.newInstance(service);

        // Configure workflow options
        WorkflowOptions options = WorkflowOptions.newBuilder()
            .setTaskQueue(StartWorker.TASK_QUEUE)  // Must match worker's task queue
            .setWorkflowId("greeting-workflow")     // Unique identifier for this workflow
            .build();

        // Create workflow stub and execute
        GreetingWorkflow workflow = client.newWorkflowStub(GreetingWorkflow.class, options);
        String result = workflow.greet("Temporal");
        System.out.println(result);
    }
}
