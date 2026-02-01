# Exercise 7 Solution: Manual Activity Retry

## Key Implementation Points

### 1. Disable Automatic Retries
```go
ao := workflow.ActivityOptions{
	StartToCloseTimeout: 5 * workflow.Second,
	RetryPolicy: &workflow.RetryPolicy{
		MaximumAttempts: 1,
	},
}
ctx = workflow.WithActivityOptions(ctx, ao)
```

### 2. Retry Signal Handler
```go
approvalChannel := workflow.GetSignalChannel(ctx, "approve")
workflow.Go(ctx, func(ctx workflow.Context) {
	var value bool
	approvalChannel.Receive(ctx, &value)
	logger.Info("Approval received", "approved", value)
	approved = value
	approvalReceived = true
})

retryChannel := workflow.GetSignalChannel(ctx, "retry")
workflow.Go(ctx, func(ctx workflow.Context) {
	for {
		var update RetryUpdate
		retryChannel.Receive(ctx, &update)
		logger.Info("Retry update received", "key", update.Key, "value", update.Value)
		switch update.Key {
		case "fromAccount":
			updatedRequest.FromAccount = update.Value
		case "toAccount":
			updatedRequest.ToAccount = update.Value
		case "amount":
			if amt, err := strconv.ParseFloat(update.Value, 64); err == nil {
				updatedRequest.Amount = amt
			}
		}
		retryRequested = true
	}
})
```

### 3. Manual Retry Loop
```go
for {
	err = workflow.ExecuteActivity(ctx, Withdraw, updatedRequest.FromAccount, updatedRequest.Amount).Get(ctx, nil)
	if err == nil {
		break
	}
	status = StatusRetrying
	retryRequested = false
	workflow.Await(ctx, func() bool { return retryRequested })
}
```

### 4. Complete Workflow Pattern
```go
// Withdraw with retry
for {
	err = workflow.ExecuteActivity(ctx, Withdraw, updatedRequest.FromAccount, updatedRequest.Amount).Get(ctx, nil)
	if err == nil {
		break
	}
	status = StatusRetrying
	retryRequested = false
	workflow.Await(ctx, func() bool { return retryRequested })
}

// Wait for approval
workflow.Await(ctx, func() bool { return approvalReceived })

if approved {
	status = StatusApproved
	// Deposit with retry
	for {
		err = workflow.ExecuteActivity(ctx, Deposit, updatedRequest.ToAccount, updatedRequest.Amount).Get(ctx, nil)
		if err == nil {
			break
		}
		status = StatusRetrying
		retryRequested = false
		workflow.Await(ctx, func() bool { return retryRequested })
	}
	status = StatusCompleted
	return "Transfer completed successfully", nil
}
```

## Key Concepts
- Manual retry loops for handling correctable errors
- Signal-based data correction
- Dynamic request updates during execution
- Status tracking for retry states

## Testing
1. Start worker: `go run exercise7/worker/main.go`
2. Run workflow: `go run exercise7/starter/main.go`
3. Send signals to correct data and retry operations
