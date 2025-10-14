package com.temporal.training.solution7;

import io.temporal.failure.ApplicationFailure;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class BankingActivitiesImpl implements BankingActivities {
    
    private static final Logger logger = LoggerFactory.getLogger(BankingActivitiesImpl.class);
    
    @Override
    public void withdraw(String accountId, double amount) {
        logger.info("Withdrawing ${} from account {}", amount, accountId);
        
        // Simulate invalid account data that requires manual retry
        if (accountId.contains("invalid")) {
            throw ApplicationFailure.newNonRetryableFailure("Invalid fromAccount ID: " + accountId, "InvalidAccount");
        }
        
        logger.info("Successfully withdrew ${} from account {}", amount, accountId);
    }
    
    @Override
    public void deposit(String accountId, double amount) {
        logger.info("Depositing ${} to account {}", amount, accountId);
        
        // Simulate invalid account data that requires manual retry
        if (accountId.contains("invalid")) {
            throw ApplicationFailure.newNonRetryableFailure("Invalid toAccount ID: " + accountId, "InvalidAccount");
        }
        
        logger.info("Successfully deposited ${} to account {}", amount, accountId);
    }
    
    @Override
    public void refund(String accountId, double amount) {
        logger.info("Refunding ${} to account {}", amount, accountId);
        logger.info("Successfully refunded ${} to account {}", amount, accountId);
    }
}
