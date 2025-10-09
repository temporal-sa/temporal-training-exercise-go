package com.temporal.training.exercise2;

/**
 * Transfer request data using Java 17 record.
 */
public record TransferRequest(
    String fromAccount,
    String toAccount,
    double amount,
    String transferId
) {}