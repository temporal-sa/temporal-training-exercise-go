# Exercise 2 Solution: Money Transfer Basics

## Key Implementation Points

### 1. activities.go
```go
func Withdraw(ctx context.Context, account string, amount float64) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Withdrawing", "amount", amount, "account", account)
	if rand.Float64() < 0.1 {
		return fmt.Errorf("withdrawal failed - insufficient funds")
	}
	return nil
}
```

### 2. workflow.go
```go
func (w *MoneyTransferWorkflow) Transfer(ctx workflow.Context, request TransferRequest) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting transfer", "request", request)

	ao := workflow.ActivityOptions{StartToCloseTimeout: 5 * time.Second}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var approved, approvalReceived bool

	approvalChannel := workflow.GetSignalChannel(ctx, "approve")
	workflow.Go(ctx, func(ctx workflow.Context) {
		var value bool
		approvalChannel.Receive(ctx, &value)
		logger.Info("Approval received", "approved", value)
		approved = value
		approvalReceived = true
	})

	err := workflow.ExecuteActivity(ctx, Withdraw, request.FromAccount, request.Amount).Get(ctx, nil)
	if err != nil {
		return "", err
	}
	logger.Info("Withdrawal completed")

	logger.Info("Waiting for approval...")
	workflow.Await(ctx, func() bool { return approvalReceived })

	if approved {
		err = workflow.ExecuteActivity(ctx, Deposit, request.ToAccount, request.Amount).Get(ctx, nil)
		if err != nil {
			return "", err
		}
		logger.Info("Transfer approved and completed")
		return "Transfer completed successfully", nil
	}

	err = workflow.ExecuteActivity(ctx, Refund, request.FromAccount, request.Amount).Get(ctx, nil)
	if err != nil {
		return "", err
	}
	logger.Info("Transfer rejected and refunded")
	return "Transfer rejected and refunded", nil
}
```

### 3. worker/main.go
```go
w.RegisterWorkflow((&exercise2.MoneyTransferWorkflow{}).Transfer)
w.RegisterActivity(exercise2.Withdraw)
w.RegisterActivity(exercise2.Deposit)
w.RegisterActivity(exercise2.Refund)
```

### 4. starter/main.go
```go
time.Sleep(2 * time.Second)
err = c.SignalWorkflow(context.Background(), we.GetID(), we.GetRunID(), "approve", true)
```

## Key Concepts Demonstrated
1. **Multiple Activities**: withdraw, deposit, refund operations
2. **Signal Handling**: External approval mechanism  
3. **workflow.Await()**: Blocking until condition is met
4. **Compensation Pattern**: Refund when transfer is rejected
