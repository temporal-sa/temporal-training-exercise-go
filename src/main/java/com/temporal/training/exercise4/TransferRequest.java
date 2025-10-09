package com.temporal.training.exercise4;

public record TransferRequest(
    String fromAccount,
    String toAccount,
    double amount,
    String transferId
) {}