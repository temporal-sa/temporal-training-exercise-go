# Exercise 7: Manual Activity Retry

## Objective
Learn how to implement manual retry patterns using signals when automatic retries are insufficient for handling invalid data scenarios.

## Key Concepts
- Disabling automatic activity retries
- Manual retry pattern using signals
- Interactive error correction
- Dynamic request updates during workflow execution

## What You'll Implement

### 1. Activity Configuration
Configure ActivityOptions to disable automatic retries by setting `MaximumAttempts: 1` in the RetryPolicy.

### 2. Manual Retry Logic
Wrap activity calls in a for loop that:
- Executes operations
- Catches errors and sets status to `StatusRetrying`
- Waits for retry signal before continuing

### 3. Signal Handlers
- `approve`: Handle approval/rejection signals
- `retry`: Handle retry signals with updated data (fromAccount, toAccount, amount)

### 4. Search Attributes
Upsert the `AccountId` search attribute for workflow visibility.

## Testing the Exercise

1. Start the worker:
```bash
go run exercise7/worker/main.go
```

2. Run the workflow:
```bash
go run exercise7/starter/main.go
```

3. Use Temporal CLI to interact with the workflow:

Send approval:
```bash
temporal workflow signal \
  --workflow-id money-transfer-workflow \
  --name approve \
  --input true
```

Retry with corrected data:
```bash
temporal workflow signal \
  --workflow-id money-transfer-workflow \
  --name retry \
  --input '{"Key":"toAccount","Value":"account-456"}'
```

## Expected Behavior
1. Workflow starts and attempts withdrawal (succeeds)
2. Workflow waits for approval signal
3. After approval, attempts deposit (may fail)
4. On failure, workflow enters RETRYING status
5. Send retry signal with corrected account data
6. Workflow completes successfully

## Key Learning Points
- When to use manual vs automatic retries
- Signal-based error correction patterns
- Interactive workflow debugging
- Handling invalid data scenarios gracefully
