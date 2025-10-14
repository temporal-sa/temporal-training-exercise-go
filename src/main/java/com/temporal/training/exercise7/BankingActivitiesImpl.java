package com.temporal.training.exercise7;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class BankingActivitiesImpl implements BankingActivities {
    
    private static final Logger logger = LoggerFactory.getLogger(BankingActivitiesImpl.class);
    
    @Override
    public void withdraw(String accountId, double amount) {
        logger.info("Withdrawing ${} from account {}", amount, accountId);
        
        if (accountId.contains("invalid")) {
            // TODO: Throw ApplicationFailure.newNonRetryableFailure here
        }
        
        logger.info("Successfully withdrew ${} from account {}", amount, accountId);
    }
    
    @Override
    public void deposit(String accountId, double amount) {
        logger.info("Depositing ${} to account {}", amount, accountId);
        
        if (accountId.contains("invalid")) {
            // TODO: Throw ApplicationFailure.newNonRetryableFailure here
        }
        
        logger.info("Successfully deposited ${} to account {}", amount, accountId);
    }
    
    @Override
    public void refund(String accountId, double amount) {
        logger.info("Refunding ${} to account {}", amount, accountId);
        logger.info("Successfully refunded ${} to account {}", amount, accountId);
    }
}
