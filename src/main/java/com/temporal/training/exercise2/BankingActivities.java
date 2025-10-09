package com.temporal.training.exercise2;

import io.temporal.activity.ActivityInterface;
import io.temporal.activity.ActivityMethod;

/**
 * Banking activities interface for money transfer operations.
 * TODO: Implement withdraw, deposit, and refund activities.
 */
@ActivityInterface
public interface BankingActivities {
    
    @ActivityMethod
    void withdraw(String account, double amount);
    
    @ActivityMethod
    void deposit(String account, double amount);
    
    @ActivityMethod
    void refund(String account, double amount);
}