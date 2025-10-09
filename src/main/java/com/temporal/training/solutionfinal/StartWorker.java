package com.temporal.training.solutionfinal;

import com.temporal.training.solutionfinal.activity.BankingActivitiesImpl;
import com.temporal.training.solutionfinal.workflow.MoneyTransferWorkflowImpl;

import io.temporal.client.WorkflowClient;
import io.temporal.serviceclient.WorkflowServiceStubs;
import io.temporal.worker.Worker;
import io.temporal.worker.WorkerFactory;

public class StartWorker {
    
    public static final String TASK_QUEUE = "money-transfer-queue";
    
    public static void main(String[] args) {
        WorkflowServiceStubs service = WorkflowServiceStubs.newLocalServiceStubs();
        WorkflowClient client = WorkflowClient.newInstance(service);
        WorkerFactory factory = WorkerFactory.newInstance(client);
        
        Worker worker = factory.newWorker(TASK_QUEUE);
        worker.registerWorkflowImplementationTypes(MoneyTransferWorkflowImpl.class);
        worker.registerActivitiesImplementations(new BankingActivitiesImpl());
        
        factory.start();
        System.out.println("Worker started for task queue: " + TASK_QUEUE);
    }
}
