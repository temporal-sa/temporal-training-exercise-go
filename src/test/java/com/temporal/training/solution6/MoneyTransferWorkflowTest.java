package com.temporal.training.solution6;

import io.temporal.api.enums.v1.IndexedValueType;
import io.temporal.client.WorkflowOptions;
import io.temporal.testing.TestWorkflowRule;
import org.junit.After;
import org.junit.Rule;
import org.junit.Test;
import org.mockito.Mockito;

import static org.junit.Assert.assertEquals;
import static org.mockito.ArgumentMatchers.anyDouble;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.never;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.withSettings;

public class MoneyTransferWorkflowTest {

    @Rule
    public TestWorkflowRule testWorkflowRule = TestWorkflowRule.newBuilder()
            .setWorkflowTypes(MoneyTransferWorkflowImpl.class)
            .setDoNotStart(true)
            .build();

    private void registerSearchAttribute() {
        testWorkflowRule.getTestEnvironment().registerSearchAttribute(
            "AccountId", IndexedValueType.INDEXED_VALUE_TYPE_TEXT
        );
    }

    @After
    public void tearDown() {
        testWorkflowRule.getTestEnvironment().shutdown();
    }

    @Test
    public void testSuccessfulTransfer() {
        BankingActivities mockActivities = Mockito.mock(
                BankingActivities.class,
                withSettings().withoutAnnotations()
        );
        
        registerSearchAttribute();
        testWorkflowRule.getWorker().registerActivitiesImplementations(mockActivities);
        testWorkflowRule.getTestEnvironment().start();

        MoneyTransferWorkflow workflow = testWorkflowRule
                .getWorkflowClient()
                .newWorkflowStub(
                        MoneyTransferWorkflow.class,
                        WorkflowOptions.newBuilder().setTaskQueue(testWorkflowRule.getTaskQueue()).build()
                );
        
        TransferRequest request = new TransferRequest("account-123", "account-456", 100.0, "transfer-1");

        testWorkflowRule.getTestEnvironment().registerDelayedCallback(
            java.time.Duration.ofSeconds(1),
            () -> workflow.approve(true)
        );

        String result = workflow.transfer(request);

        assertEquals("Transfer completed successfully", result);
        assertEquals(TransferStatus.COMPLETED, workflow.getStatus());
        
        verify(mockActivities).withdraw("account-123", 100.0);
        verify(mockActivities).deposit("account-456", 100.0);
        verify(mockActivities, never()).refund(anyString(), anyDouble());
    }

    @Test
    public void testRejectedTransfer() {
        BankingActivities mockActivities = Mockito.mock(
                BankingActivities.class,
                withSettings().withoutAnnotations()
        );
        
        registerSearchAttribute();
        testWorkflowRule.getWorker().registerActivitiesImplementations(mockActivities);
        testWorkflowRule.getTestEnvironment().start();

        MoneyTransferWorkflow workflow = testWorkflowRule
                .getWorkflowClient()
                .newWorkflowStub(
                        MoneyTransferWorkflow.class,
                        WorkflowOptions.newBuilder().setTaskQueue(testWorkflowRule.getTaskQueue()).build()
                );
        
        TransferRequest request = new TransferRequest("account-123", "account-456", 100.0, "transfer-2");

        testWorkflowRule.getTestEnvironment().registerDelayedCallback(
            java.time.Duration.ofSeconds(1),
            () -> workflow.approve(false)
        );

        String result = workflow.transfer(request);

        assertEquals("Transfer rejected and refunded", result);
        assertEquals(TransferStatus.CANCELLED, workflow.getStatus());
        
        verify(mockActivities).withdraw("account-123", 100.0);
        verify(mockActivities).refund("account-123", 100.0);
        verify(mockActivities, never()).deposit(anyString(), anyDouble());
    }

    @Test
    public void testQueryStatus() {
        BankingActivities mockActivities = Mockito.mock(
                BankingActivities.class,
                withSettings().withoutAnnotations()
        );
        
        registerSearchAttribute();
        testWorkflowRule.getWorker().registerActivitiesImplementations(mockActivities);
        testWorkflowRule.getTestEnvironment().start();

        MoneyTransferWorkflow workflow = testWorkflowRule
                .getWorkflowClient()
                .newWorkflowStub(
                        MoneyTransferWorkflow.class,
                        WorkflowOptions.newBuilder().setTaskQueue(testWorkflowRule.getTaskQueue()).build()
                );
        
        TransferRequest request = new TransferRequest("account-123", "account-456", 100.0, "transfer-3");

        testWorkflowRule.getTestEnvironment().registerDelayedCallback(
            java.time.Duration.ofMillis(500),
            () -> assertEquals(TransferStatus.PENDING, workflow.getStatus())
        );

        testWorkflowRule.getTestEnvironment().registerDelayedCallback(
            java.time.Duration.ofSeconds(1),
            () -> workflow.approve(true)
        );

        workflow.transfer(request);
        assertEquals(TransferStatus.COMPLETED, workflow.getStatus());
    }
}