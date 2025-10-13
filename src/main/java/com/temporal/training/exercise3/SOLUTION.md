# Exercise 3 Solution: Flow Control with Query Handlers

## Key Implementation Points

### 1. MoneyTransferWorkflow Interface
```java
@QueryMethod
TransferStatus getStatus();
```
- Add `@QueryMethod` to query method
- Queries are read-only and don't modify workflow state

### 2. MoneyTransferWorkflowImpl - Status Tracking
```java
private TransferStatus status = TransferStatus.PENDING;

@Override
public String transfer(TransferRequest request) {
    this.request = request;
    
    try {
        activities.withdraw(request.fromAccount(), request.amount());
        
        Workflow.await(() -> approvalReceived);
        
        if (approved) {
            status = TransferStatus.APPROVED;
            activities.deposit(request.toAccount(), request.amount());
            status = TransferStatus.COMPLETED;
            return "Transfer completed successfully";
        } else {
            activities.refund(request.fromAccount(), request.amount());
            status = TransferStatus.CANCELLED;
            return "Transfer rejected and refunded";
        }
    } catch (Exception e) {
        status = TransferStatus.FAILED;
        throw e;
    }
}
```

### 3. Query Handler Implementation
```java
@Override
public TransferStatus getStatus() {
    return status;
}
```
- Simple getter that returns current workflow state
- Can be called anytime during workflow execution

### 4. StartWorkflow with Query Usage
```java
CompletableFuture<String> result = WorkflowClient.execute(workflow::transfer, request);

// Query before approval
System.out.println("Current status: " + workflow.getStatus());

workflow.approve(true);

// Query after approval
System.out.println("Status after approval: " + workflow.getStatus());
```

## Key Concepts Demonstrated

1. **Query Methods**: Read-only workflow state inspection
2. **Status Tracking**: Maintaining workflow state throughout execution
3. **External Visibility**: Clients can monitor workflow progress
4. **Signal vs Query**: Signals modify state, queries only read
5. **Asynchronous Monitoring**: Query running workflows anytime

## Testing the Solution

1. **Start Worker**: `./gradlew execute -PmainClass=com.temporal.training.solution3.StartWorker`
2. **Run Workflow**: `./gradlew execute -PmainClass=com.temporal.training.solution3.StartWorkflow`

Expected output shows:
- Initial status: PENDING
- Status progression: PENDING → APPROVED → COMPLETED
- Query responses throughout execution
