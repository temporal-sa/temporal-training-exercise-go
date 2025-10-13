# Exercise 2: Money Transfer Basics (45 min)

## Learning Objectives
- Implement multiple activities (withdraw, deposit, refund)
- Use signals for workflow interaction
- Implement await mechanism for approval
- Handle basic error scenarios with compensation

## Tasks

### 1. Complete BankingActivities Interface (5 min)
- Add `@ActivityMethod` annotations to all methods
- Review the three activity methods: withdraw, deposit, refund

### 2. Implement BankingActivitiesImpl (10 min)
- Implement `withdraw()`: Add logging and error simulation
- Implement `deposit()`: Add logging and error simulation  
- Implement `refund()`: Add logging for compensation
- Use Math.random() to simulate network failures
- Throw RuntimeException for failures

### 3. Complete MoneyTransferWorkflow Interface (5 min)
- Add `@SignalMethod` annotation to the `approve()` method

### 4. Implement MoneyTransferWorkflowImpl (15 min)
- Create activity stub with proper ActivityOptions (30 second timeout)
- Implement transfer workflow logic:
  1. Call withdraw activity
  2. Wait for approval signal using `Workflow.await()`
  3. If approved: call deposit activity
  4. If not approved: call refund activity
  5. Return appropriate status message

### 5. Complete StartWorker (5 min)
- Set up WorkflowServiceStubs, WorkflowClient, WorkerFactory
- Register MoneyTransferWorkflowImpl and BankingActivitiesImpl
- Start the worker

### 6. Complete StartWorkflow (5 min)
- Create and start workflow with a sample TransferRequest
- Send approval signal after a short delay
- Print the workflow result

## Key Concepts
- **Activities**: External operations that can fail and be retried
- **Signals**: Asynchronous messages sent to running workflows
- **Workflow.await()**: Wait for conditions to be met
- **Compensation**: Undoing operations when things go wrong

## Testing
1. Start the worker: `./gradlew execute -PmainClass=com.temporal.training.exercise2.StartWorker`
2. Run the workflow: `./gradlew execute -PmainClass=com.temporal.training.exercise2.StartWorkflow`
3. Observe the workflow waiting for approval, then completing the transfer

## Next Steps
Exercise 3 will add timeout handling and query methods for better observability.
