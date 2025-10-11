package com.temporal.training.solution6;

import io.temporal.activity.ActivityInterface;
import io.temporal.activity.ActivityMethod;

@ActivityInterface
public interface BankingActivities {
    
    @ActivityMethod
    void withdraw(String account, double amount);
    
    @ActivityMethod
    void deposit(String account, double amount);
    
    @ActivityMethod
    void refund(String account, double amount);
}