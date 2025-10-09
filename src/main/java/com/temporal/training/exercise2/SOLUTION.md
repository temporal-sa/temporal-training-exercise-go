# Exercise 2 Solution: Money Transfer Basics

## Key Implementation Points

### 1. BankingActivities Interface
```java
@ActivityMethod
void withdraw(String account, double amount);
```
- Add `@ActivityMethod` to all three methods
- Activities represent external operations that can fail and retry

### 2. BankingActivitiesImpl
```java
public void withdraw(String account, double amount) {
    logger.info("Withdrawing ${} from account {}", amount, account);
    // Simulate external API call
    if (Math.random() < 0.1) {
        throw new RuntimeException("Withdrawal failed - insufficient funds");
    }
}
```
- Add proper logging with SLF4J
- Include error simulation for realistic failures
- No return values needed for simple operations

### 3. MoneyTransferWorkflow Interface
```java
@SignalMethod
void approve(boolean approved);
```
- Add `@SignalMethod` annotation for external communication

### 4. MoneyTransferWorkflowImpl - Key Parts
```java
private final BankingActivities activities = Workflow.newActivityStub(
    BankingActivities.class,
    ActivityOptions.newBuilder()
        .setStartToCloseTimeout(Duration.ofSeconds(30))
        .build()
);

private boolean approved = false;
private boolean approvalReceived = false;
```

**Transfer Logic:**
```java
// 1. Withdraw
activities.withdraw(request.fromAccount(), request.amount());

// 2. Wait for approval
Workflow.await(() -> approvalReceived);

// 3. Deposit or refund based on approval
if (approved) {
    activities.deposit(request.toAccount(), request.amount());
    return "Transfer completed successfully";
} else {
    activities.refund(request.fromAccount(), request.amount());
    return "Transfer rejected and refunded";
}
```

**Signal Handler:**
```java
@Override
public void approve(boolean approved) {
    this.approved = approved;
    this.approvalReceived = true; // Unblocks Workflow.await()
}
```

### 5. StartWorker
```java
Worker worker = factory.newWorker(TASK_QUEUE);
worker.registerWorkflowImplementationTypes(MoneyTransferWorkflowImpl.class);
worker.registerActivitiesImplementations(new BankingActivitiesImpl());
factory.start();
```

### 6. StartWorkflow
```java
CompletableFuture<String> result = WorkflowClient.execute(workflow::transfer, request);
Thread.sleep(3000); // Wait before approval
workflow.approve(true); // Send signal
String transferResult = result.get(); // Get result
```

## Key Concepts Demonstrated

1. **Multiple Activities**: withdraw, deposit, refund operations
2. **Signal Handling**: External approval mechanism  
3. **Workflow.await()**: Blocking until condition is met
4. **Compensation Pattern**: Refund when transfer is rejected
5. **Asynchronous Execution**: Non-blocking workflow start

## Testing the Solution

1. **Start Worker**: `./gradlew run -PmainClass=com.temporal.training.solution2.StartWorker`
2. **Run Workflow**: `./gradlew run -PmainClass=com.temporal.training.solution2.StartWorkflow`

Expected output shows:
- Withdrawal completion
- Waiting for approval
- Approval signal received  
- Deposit completion
- Final transfer result