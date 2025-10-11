package com.temporal.training.exercise6;

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
        // TODO: Register the "AccountId" search attribute with TEXT type
        // Hint: Use testWorkflowRule.getTestEnvironment().registerSearchAttribute()
    }

    @After
    public void tearDown() {
        testWorkflowRule.getTestEnvironment().shutdown();
    }

    @Test
    public void testSuccessfulTransfer() {
        // TODO: Create a mock BankingActivities using Mockito
        // Hint: Use Mockito.mock() with withSettings().withoutAnnotations()
        BankingActivities mockActivities = null;
        
        // TODO: Register search attribute, activities, and start test environment
        
        // TODO: Create workflow stub using testWorkflowRule
        MoneyTransferWorkflow workflow = null;
        
        TransferRequest request = new TransferRequest("account-123", "account-456", 100.0, "transfer-1");

        // TODO: Register delayed callback to approve the transfer after 1 second
        // Hint: Use testWorkflowRule.getTestEnvironment().registerDelayedCallback()

        // TODO: Execute the workflow and verify results
        String result = null;

        // TODO: Add assertions to verify:
        // - Result message is "Transfer completed successfully"
        // - Final status is COMPLETED
        // - withdraw() was called with correct parameters
        // - deposit() was called with correct parameters  
        // - refund() was never called
    }

    @Test
    public void testRejectedTransfer() {
        // TODO: Implement test for rejected transfer scenario
        // Similar to testSuccessfulTransfer but:
        // - Send approval(false) instead of approval(true)
        // - Verify result is "Transfer rejected and refunded"
        // - Verify status is CANCELLED
        // - Verify withdraw() and refund() were called, but not deposit()
    }

    @Test
    public void testQueryStatus() {
        // TODO: Implement test for query functionality
        // - Register callback to check status is PENDING after 500ms
        // - Register callback to approve after 1 second
        // - Verify final status is COMPLETED
    }
}