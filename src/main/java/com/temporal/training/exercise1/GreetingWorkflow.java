package com.temporal.training.exercise1;

import io.temporal.workflow.WorkflowInterface;
import io.temporal.workflow.WorkflowMethod;

/**
 * Workflow interface defines the contract for a Temporal workflow.
 * The @WorkflowInterface annotation marks this as a workflow.
 */
@WorkflowInterface
public interface GreetingWorkflow {
    
    /**
     * The @WorkflowMethod annotation marks the main entry point of the workflow.
     * This method will be called when the workflow is started.
     * 
     * @param name The name to greet
     * @return A greeting message
     */
    @WorkflowMethod
    String greet(String name);
}