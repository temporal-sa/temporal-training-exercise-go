package com.temporal.training.solution3;

public record TransferRequest(
    String fromAccount,
    String toAccount,
    double amount,
    String transferId
) {}