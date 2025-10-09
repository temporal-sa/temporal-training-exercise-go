package com.temporal.training.exercise1;

import io.temporal.activity.ActivityInterface;
import io.temporal.activity.ActivityMethod;

/**
 * Activity interface defines the contract for activities.
 * Activities perform the actual work (external API calls, database operations, etc.)
 * The @ActivityInterface annotation marks this as an activity interface.
 */
@ActivityInterface
public interface GreetingActivity {
    
    /**
     * The @ActivityMethod annotation marks methods that can be called as activities.
     * 
     * @param name The name to create a greeting for
     * @return A greeting message
     */
    @ActivityMethod
    String createGreeting(String name);
}