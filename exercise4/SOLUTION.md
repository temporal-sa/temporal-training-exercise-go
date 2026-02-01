# Exercise 4 Solution: Visibility & Monitoring with Custom Search Attributes

## Key Implementation

### Upsert Search Attributes
```go
err := workflow.UpsertSearchAttributes(ctx, map[string]interface{}{
	"AccountId": request.FromAccount,
})
if err != nil {
	return "", err
}
```

## Key Points
- Place early in workflow execution
- Enables filtering workflows by account
- Immediately visible in Temporal Web UI
- Can be updated multiple times during workflow execution

## Testing
1. Start worker: `go run exercise4/worker/main.go`
2. Run workflow: `go run exercise4/starter/main.go`
3. Verify in Web UI: Open http://localhost:8233 and check workflow details
4. Filter workflows by AccountId using CLI:
```bash
temporal workflow list --query 'AccountId="account-123"'
```

## Expected Results
- Workflow executes with all previous functionality intact
- AccountId search attribute visible in Temporal Web UI
- Can filter workflows using: `AccountId = "account-123"`
