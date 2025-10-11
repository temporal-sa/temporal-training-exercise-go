# Exercise 6: Testing Strategy & Time Skipping

## Learning Objectives
- Write comprehensive unit tests for Temporal workflows
- Use TestWorkflowRule for isolated testing
- Mock activities using Mockito with proper configuration
- Register search attributes in test environment
- Use time skipping for fast test execution
- Test different workflow scenarios (success, failure, queries)

## Background
Testing Temporal workflows requires special considerations:
- **TestWorkflowRule**: Provides isolated test environment with time control
- **Activity Mocking**: Use Mockito with `withSettings().withoutAnnotations()` to avoid annotation conflicts
- **Search Attributes**: Must be registered in test environment before use
- **Time Skipping**: Use `registerDelayedCallback()` for fast signal/timer testing
- **Scenario Testing**: Test happy path, error cases, and query functionality

## Your Task

Complete the unit tests in `MoneyTransferWorkflowTest.java` by implementing the TODO sections:

### 1. Register Search Attribute
In `registerSearchAttribute()` method:
```java
testWorkflowRule.getTestEnvironment().registerSearchAttribute(
    "AccountId", IndexedValueType.INDEXED_VALUE_TYPE_TEXT
);
```

### 2. Create Mock Activities
Use Mockito to create activity mocks:
```java
BankingActivities mockActivities = Mockito.mock(
    BankingActivities.class,
    withSettings().withoutAnnotations()
);
```

### 3. Setup Test Environment
For each test:
- Call `registerSearchAttribute()`
- Register mock activities with worker
- Start test environment

### 4. Create Workflow Stub
```java
MoneyTransferWorkflow workflow = testWorkflowRule
    .getWorkflowClient()
    .newWorkflowStub(
        MoneyTransferWorkflow.class,
        WorkflowOptions.newBuilder().setTaskQueue(testWorkflowRule.getTaskQueue()).build()
    );
```

### 5. Use Time Skipping for Signals
```java
testWorkflowRule.getTestEnvironment().registerDelayedCallback(
    java.time.Duration.ofSeconds(1),
    () -> workflow.approve(true)
);
```

### 6. Verify Activity Calls
```java
verify(mockActivities).withdraw("account-123", 100.0);
verify(mockActivities).deposit("account-456", 100.0);
verify(mockActivities, never()).refund(anyString(), anyDouble());
```

## Test Scenarios to Implement

### testSuccessfulTransfer()
- Mock activities and start test environment
- Execute workflow with approval signal after 1 second
- Verify successful completion and correct activity calls

### testRejectedTransfer()
- Similar setup but send `approve(false)`
- Verify rejection, refund, and no deposit

### testQueryStatus()
- Test query functionality at different workflow stages
- Verify status changes from PENDING to COMPLETED

## Expected Behavior
After completing the implementation:
1. All 3 tests should pass
2. Tests should execute quickly (under 10 seconds total)
3. Activity interactions should be properly verified
4. Different workflow paths should be covered

## Running Your Tests

```bash
gradle test --tests "com.temporal.training.exercise6.MoneyTransferWorkflowTest"
```

## Key Testing Concepts
- **Isolation**: Each test runs in its own environment
- **Determinism**: Time skipping makes tests predictable
- **Coverage**: Test success, failure, and query scenarios
- **Verification**: Mock verification ensures correct activity calls
- **Fast Execution**: No real time delays in tests

## Common Pitfalls
- Forgetting `withSettings().withoutAnnotations()` for activity mocks
- Not registering search attributes before starting test environment
- Incorrect callback timing for signals
- Missing activity verification assertions