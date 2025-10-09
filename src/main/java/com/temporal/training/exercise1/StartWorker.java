package com.temporal.training.exercise1;

import io.temporal.client.WorkflowClient;
import io.temporal.serviceclient.WorkflowServiceStubs;
import io.temporal.worker.Worker;
import io.temporal.worker.WorkerFactory;

/**
 * Worker is responsible for executing workflow and activity code.
 * It polls task queues for work and executes the registered implementations.
 */
public class StartWorker {
    public static final String TASK_QUEUE = "hello-task-queue";

    public static void main(String[] args) {
        // Connect to Temporal service (local development server)
        WorkflowServiceStubs service = WorkflowServiceStubs.newLocalServiceStubs();
        WorkflowClient client = WorkflowClient.newInstance(service);
        WorkerFactory factory = WorkerFactory.newInstance(client);

        // Create a worker that listens to the task queue
        Worker worker = factory.newWorker(TASK_QUEUE);
        
        // TODO: Register the workflow implementation
        // Use: worker.registerWorkflowImplementationTypes(GreetingWorkflowImpl.class);
        
        // TODO: Register the activity implementation  
        // Use: worker.registerActivitiesImplementations(new GreetingActivityImpl());

        // Start the worker
        factory.start();
    }
}
