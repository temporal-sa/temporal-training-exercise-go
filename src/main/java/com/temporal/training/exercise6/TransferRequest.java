package com.temporal.training.exercise6;

public record TransferRequest(
    String fromAccount,
    String toAccount,
    double amount,
    String transferId
) {}