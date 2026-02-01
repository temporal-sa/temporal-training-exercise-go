# Exercise 3 Solution: Flow Control with Query Handlers

## Key Implementation Points

### Query Handler
```go
err := workflow.SetQueryHandler(ctx, "getStatus", func() (TransferStatus, error) {
	return status, nil
})
if err != nil {
	return "", err
}
```

### Signal Handler
```go
approvalChannel := workflow.GetSignalChannel(ctx, "approve")
workflow.Go(ctx, func(ctx workflow.Context) {
	var value bool
	approvalChannel.Receive(ctx, &value)
	logger.Info("Approval received", "approved", value)
	approved = value
	approvalReceived = true
})
```

### Status Tracking
```go
var status TransferStatus = StatusPending

// After approval
if approved {
	status = StatusApproved
	err = workflow.ExecuteActivity(ctx, Deposit, request.ToAccount, request.Amount).Get(ctx, nil)
	if err != nil {
		status = StatusFailed
		return "", err
	}
	status = StatusCompleted
	return "Transfer completed successfully", nil
}

// After rejection
err = workflow.ExecuteActivity(ctx, Refund, request.FromAccount, request.Amount).Get(ctx, nil)
if err != nil {
	status = StatusFailed
	return "", err
}
status = StatusCancelled
return "Transfer rejected and refunded", nil
```

## Key Concepts
1. **Query Handlers**: Read-only workflow state inspection using `workflow.SetQueryHandler()`
2. **Status Tracking**: Maintaining workflow state throughout execution
3. **External Visibility**: Clients can monitor workflow progress
4. **Signal vs Query**: Signals use `workflow.SetSignalChannel()` to modify state, Queries use `workflow.SetQueryHandler()` to read state
