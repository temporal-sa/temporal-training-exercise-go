package com.temporal.training.solution7;

public record TransferRequest(
    String fromAccount,
    String toAccount,
    double amount,
    String transferId
) {
    public TransferRequest withFromAccount(String fromAccount) {
        return new TransferRequest(fromAccount, toAccount, amount, transferId);
    }
    
    public TransferRequest withToAccount(String toAccount) {
        return new TransferRequest(fromAccount, toAccount, amount, transferId);
    }
    
    public TransferRequest withAmount(double amount) {
        return new TransferRequest(fromAccount, toAccount, amount, transferId);
    }
}