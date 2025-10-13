# Exercise 3: Flow Control with Query Handlers (35 min)

## Learning Objectives
- Implement Query methods for workflow state inspection
- Track workflow status throughout execution
- Query workflow state from external clients
- Understand the difference between Signals and Queries

## Tasks

### 1. Complete MoneyTransferWorkflow Interface (5 min)
- Add `@QueryMethod` annotation to `getStatus()` method

### 2. Implement Query Handler in MoneyTransferWorkflowImpl (10 min)
- Implement `getStatus()` method to return current status
- Add status tracking throughout workflow execution:
  - Update status to `APPROVED` when approval received
  - Update status to `COMPLETED` when transfer succeeds
  - Update status to `CANCELLED` when rejected
  - Update status to `FAILED` on exceptions

### 3. Complete StartWorkflow with Query Demonstrations (20 min)
- Uncomment query calls to demonstrate status monitoring
- Query workflow status before sending approval signal
- Query workflow status after approval
- Display final status

## Key Concepts
- **Query Methods**: Read-only operations to inspect workflow state
- **Status Tracking**: Maintaining workflow state for external visibility
- **Signal vs Query**: Signals modify state, Queries only read state
- **Asynchronous Execution**: Starting workflows and querying while running

## Testing
1. Start the worker: `./gradlew execute -PmainClass=com.temporal.training.exercise3.StartWorker`
2. Run the workflow: `./gradlew execute -PmainClass=com.temporal.training.exercise3.StartWorkflow`
3. Observe status changes through query outputs

## Expected Output
- Initial status: PENDING
- Status after approval: APPROVED → COMPLETED
- Final status confirmation

## Next Steps
Exercise 4 will add visibility features like Search Attributes and User Metadata.
