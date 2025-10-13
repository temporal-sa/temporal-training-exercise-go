# Exercise 5: User Metadata & Activity Summaries

## Learning Objectives
- Add descriptive summaries to activities for better observability
- Set static summaries for workflows to improve monitoring
- Understand how metadata enhances workflow visibility in Temporal Web UI

## Background
Temporal provides metadata capabilities to make workflows and activities more observable:
- **Activity Summary**: Descriptive text that appears in the Temporal Web UI for each activity execution
- **Static Summary**: Workflow-level description that helps identify the purpose of a workflow instance

## Your Task

### 1. Add Activity Summaries
In `MoneyTransferWorkflowImpl.java`, update the `ActivityFactory` methods to include descriptive summaries:

- `createWithdrawActivity()`: Add `.setSummary("Withdrawing funds from account " + accountId)`
- `createDepositActivity()`: Add `.setSummary("Depositing funds to account " + accountId)`  
- `createRefundActivity()`: Add `.setSummary("Refunding funds to account " + accountId)`

### 2. Add Workflow Static Summary
In `StartWorkflow.java`, add a static summary to the workflow options:
- Create a descriptive string: `"Money transfer workflow for " + request.fromAccount() + " to " + request.toAccount()`
- Add `.setStaticSummary(staticSummaryDetails)` to the WorkflowOptions builder

## Expected Behavior
After implementing the summaries:
1. Activities will show descriptive text in the Temporal Web UI
2. The workflow will have a clear description visible in the workflow list
3. Monitoring and debugging become easier with meaningful metadata

## Testing Your Implementation

1. Start the Temporal server:
```bash
temporal server start-dev --search-attribute AccountId=Text
```

2. Run the worker:
```bash
./gradlew execute -PmainClass=com.temporal.training.exercise5.StartWorker
```

3. Execute the workflow:
```bash
./gradlew execute -PmainClass=com.temporal.training.exercise5.StartWorkflow
```

4. Check the Temporal Web UI at http://localhost:8233 to see the summaries in action

## Key Concepts
- **Activity Summary**: Runtime metadata that describes what an activity is doing
- **Static Summary**: Workflow-level description set at workflow start time
- **Observability**: Making workflows easier to monitor and debug through descriptive metadata
