package com.temporal.training.exercise5;

public record TransferRequest(
    String fromAccount,
    String toAccount,
    double amount,
    String transferId
) {}