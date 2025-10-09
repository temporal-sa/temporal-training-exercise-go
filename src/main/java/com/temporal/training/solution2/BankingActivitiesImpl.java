package com.temporal.training.solution2;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class BankingActivitiesImpl implements BankingActivities {
    
    private static final Logger logger = LoggerFactory.getLogger(BankingActivitiesImpl.class);
    
    @Override
    public void withdraw(String account, double amount) {
        logger.info("Withdrawing ${} from account {}", amount, account);
        // Simulate external API call
        if (Math.random() < 0.1) {
            throw new RuntimeException("Withdrawal failed - insufficient funds");
        }
    }
    
    @Override
    public void deposit(String account, double amount) {
        logger.info("Depositing ${} to account {}", amount, account);
        // Simulate external API call
        if (Math.random() < 0.05) {
            throw new RuntimeException("Deposit failed - account not found");
        }
    }
    
    @Override
    public void refund(String account, double amount) {
        logger.info("Refunding ${} to account {}", amount, account);
    }
}