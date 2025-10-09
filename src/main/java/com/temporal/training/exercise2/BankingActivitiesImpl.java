package com.temporal.training.exercise2;

/**
 * Implementation of banking activities.
 * TODO: Complete the implementation of all banking operations.
 */
public class BankingActivitiesImpl implements BankingActivities {
    
    @Override
    public void withdraw(String account, double amount) {
        // TODO: Implement withdraw logic
        // Add logging and simulate external API call
        // Use Math.random() < 0.1 to simulate occasional network failures
        // Throw RuntimeException for failures
    }
    
    @Override
    public void deposit(String account, double amount) {
        // TODO: Implement deposit logic
        // Add logging and simulate external API call
        // Use Math.random() < 0.05 to simulate occasional network failures
        // Throw RuntimeException for failures
    }
    
    @Override
    public void refund(String account, double amount) {
        // TODO: Implement refund logic
        // Add logging for the compensation operation
    }
}
