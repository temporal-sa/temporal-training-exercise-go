package com.temporal.training.exercise3;

public record TransferRequest(
    String fromAccount,
    String toAccount,
    double amount,
    String transferId
) {}