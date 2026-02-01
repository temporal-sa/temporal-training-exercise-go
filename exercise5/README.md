# Exercise 5: User Metadata & Activity Summaries

## Learning Objectives
- Add descriptive summaries to activities for better observability
- Understand how metadata enhances workflow visibility in Temporal Web UI

## Background
Temporal provides metadata capabilities to make workflows and activities more observable:
- **Activity Summary**: Descriptive text that appears in the Temporal Web UI for each activity execution

## Your Task

### Add Activity Summaries
In `workflow.go`, add Summary field to ActivityOptions for each activity:

- Withdraw: `"Withdrawing funds from account {fromAccount}"`
- Deposit: `"Depositing funds to account {toAccount}"`  
- Refund: `"Refunding funds to account {fromAccount}"`

## Expected Behavior
After implementing the summaries:
1. Activities will show descriptive text in the Temporal Web UI
2. Monitoring and debugging become easier with meaningful metadata

## Testing Your Implementation

1. Start the Temporal server:
```bash
temporal server start-dev --search-attribute AccountId=Keyword
```

Alternatively, add search attribute to existing server:
```bash
temporal operator search-attribute create --name AccountId --type Keyword
```

2. Run the worker:
```bash
go run exercise5/worker/main.go
```

3. Execute the workflow:
```bash
go run exercise5/starter/main.go
```

4. Check the Temporal Web UI at http://localhost:8233 to see the summaries in action

## Key Concepts
- **Activity Summary**: Runtime metadata that describes what an activity is doing
- **Observability**: Making workflows easier to monitor and debug through descriptive metadata
