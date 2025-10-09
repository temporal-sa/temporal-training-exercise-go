package com.temporal.training.solution4;

public record TransferRequest(
    String fromAccount,
    String toAccount,
    double amount,
    String transferId
) {}