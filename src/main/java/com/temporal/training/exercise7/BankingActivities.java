package com.temporal.training.exercise7;

import io.temporal.activity.ActivityInterface;

@ActivityInterface
public interface BankingActivities {
    
    void withdraw(String accountId, double amount);
    
    void deposit(String accountId, double amount);
    
    void refund(String accountId, double amount);
}