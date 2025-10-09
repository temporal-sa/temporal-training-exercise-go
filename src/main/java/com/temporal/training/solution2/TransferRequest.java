package com.temporal.training.solution2;

public record TransferRequest(
    String fromAccount,
    String toAccount,
    double amount,
    String transferId
) {}