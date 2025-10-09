package com.temporal.training.solution1;

import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowOptions;
import io.temporal.serviceclient.WorkflowServiceStubs;

public class StartWorkflow {
    public static void main(String[] args) {
        WorkflowServiceStubs service = WorkflowServiceStubs.newLocalServiceStubs();
        WorkflowClient client = WorkflowClient.newInstance(service);

        WorkflowOptions options = WorkflowOptions.newBuilder()
            .setTaskQueue(StartWorker.TASK_QUEUE)
            .setWorkflowId("greeting-workflow")
            .build();

        GreetingWorkflow workflow = client.newWorkflowStub(GreetingWorkflow.class, options);
        String result = workflow.greet("Temporal");
        System.out.println(result);
    }
}
