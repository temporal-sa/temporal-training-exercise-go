# Exercise 2: Money Transfer Basics (45 min)

## Learning Objectives
- Implement multiple activities (withdraw, deposit, refund)
- Use signals for workflow interaction
- Implement await mechanism for approval
- Handle basic error scenarios with compensation

## Tasks

### 1. Implement BankingActivities (10 min)
- Implement `Withdraw()`: Add logging and error simulation
- Implement `Deposit()`: Add logging and error simulation  
- Implement `Refund()`: Add logging for compensation
- Use rand.Float64() to simulate network failures
- Return errors for failures

### 2. Implement MoneyTransferWorkflow (15 min)
- Configure ActivityOptions with 5 second timeout
- Set up signal channel for "approve" signal
- Implement transfer workflow logic:
  1. Call withdraw activity
  2. Wait for approval signal using `workflow.Await()`
  3. If approved: call deposit activity
  4. If not approved: call refund activity
  5. Return appropriate status message

### 3. Complete StartWorker (5 min)
- Register MoneyTransferWorkflow and all activities

### 4. Complete StartWorkflow (5 min)
- Send approval signal after a short delay

## Key Concepts
- **Activities**: External operations that can fail and be retried
- **Signals**: Asynchronous messages sent to running workflows
- **workflow.Await()**: Wait for conditions to be met
- **Compensation**: Undoing operations when things go wrong

## Testing
1. Start the worker: `go run exercise2/worker/main.go`
2. Run the workflow: `go run exercise2/starter/main.go`
3. Observe the workflow waiting for approval, then completing the transfer

## Next Steps
Exercise 3 will add query methods for better observability.
