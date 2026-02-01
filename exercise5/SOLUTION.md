# Exercise 5 Solution: User Metadata & Activity Summaries

## Solution Implementation

### Activity Summaries in workflow.go

```go
withdrawCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
	StartToCloseTimeout: 5 * time.Second,
	Summary:             fmt.Sprintf("Withdrawing funds from account %s", request.FromAccount),
})
err = workflow.ExecuteActivity(withdrawCtx, Withdraw, request.FromAccount, request.Amount).Get(ctx, nil)

// For deposit
depositCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
	StartToCloseTimeout: 5 * time.Second,
	Summary:             fmt.Sprintf("Depositing funds to account %s", request.ToAccount),
})
err = workflow.ExecuteActivity(depositCtx, Deposit, request.ToAccount, request.Amount).Get(ctx, nil)

// For refund
refundCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
	StartToCloseTimeout: 5 * time.Second,
	Summary:             fmt.Sprintf("Refunding funds to account %s", request.FromAccount),
})
err = workflow.ExecuteActivity(refundCtx, Refund, request.FromAccount, request.Amount).Get(ctx, nil)
```

## Key Points
- **Activity Summary**: Provides runtime context for each activity execution
- **Observability**: Summaries appear in Temporal Web UI for better monitoring
- **Best Practice**: Use descriptive, contextual summaries that help with debugging
