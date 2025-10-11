package com.temporal.training.solution6;

public record TransferRequest(
    String fromAccount,
    String toAccount,
    double amount,
    String transferId
) {}