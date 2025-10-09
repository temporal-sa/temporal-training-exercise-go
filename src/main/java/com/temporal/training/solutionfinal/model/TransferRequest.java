package com.temporal.training.solutionfinal.model;

public record TransferRequest(
    String fromAccount,
    String toAccount,
    double amount,
    String transferId
) {}
