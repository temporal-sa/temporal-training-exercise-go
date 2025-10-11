package com.temporal.training.solution5;

public record TransferRequest(
    String fromAccount,
    String toAccount,
    double amount,
    String transferId
) {}