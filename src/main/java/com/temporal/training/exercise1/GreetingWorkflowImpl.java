package com.temporal.training.exercise1;

import io.temporal.activity.ActivityOptions;
import io.temporal.workflow.Workflow;

import java.time.Duration;

/**
 * Workflow implementation contains the actual workflow logic.
 * This class implements the GreetingWorkflow interface.
 */
public class GreetingWorkflowImpl implements GreetingWorkflow {
    
    /**
     * Activity stub - this is how workflows call activities.
     * Workflow.newActivityStub() creates a proxy to call activity methods.
     * 
     * TODO: Create an activity stub for GreetingActivity with:
     * - ActivityOptions with setStartToCloseTimeout of 10 seconds
     */
    private final GreetingActivity activity = null; // TODO: Implement this
    
    @Override
    public String greet(String name) {
        // TODO: Call the activity's createGreeting method and return the result
        return null; // Replace with activity call
    }
}