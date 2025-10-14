# Exercise 7 Solution: Manual Activity Retry

## Overview
This exercise demonstrates how to implement manual retry patterns for activities that fail due to invalid data, requiring human intervention before retry.

## Key Concepts

### Manual Retry Pattern
- Disable automatic retries for specific failure types
- Use signals to trigger manual retries after data correction
- Handle non-retryable failures gracefully

### Implementation Details

#### 1. Non-Retryable Failures
Activities throw `ApplicationFailure.newNonRetryableFailure()` for invalid data:
```java
if (accountId.contains("invalid")) {
    throw ApplicationFailure.newNonRetryableFailure(
        "Invalid account ID: " + accountId, 
        "InvalidAccount"
    );
}
```

#### 2. Execute with Retry Pattern
Workflow uses a retry loop that waits for manual intervention:
```java
private void executeWithRetry(Runnable operation) {
    while (true) {
        try {
            operation.run();
            break;
        } catch (Exception e) {
            status = TransferStatus.RETRYING;
            retryRequested = false;
            Workflow.await(() -> retryRequested);
        }
    }
}
```

#### 3. Signal Handler for Updates
Retry signal updates request data dynamically:
```java
@Override
public void retry(RetryUpdate update) {
    updatedRequest = switch (update.key()) {
        case "fromAccount" -> updatedRequest.withFromAccount(update.value());
        case "toAccount" -> updatedRequest.withToAccount(update.value());
        case "amount" -> updatedRequest.withAmount(Double.parseDouble(update.value()));
        default -> updatedRequest;
    };
    this.retryRequested = true;
}
```

## Running the Exercise

1. Start worker:
```bash
./gradlew execute -PmainClass=com.temporal.training.exercise7.StartWorker
```

2. Run workflow with invalid account:
```bash
./gradlew execute -PmainClass=com.temporal.training.exercise7.StartWorkflow
```

3. Send retry signal with corrected data:
```bash
temporal workflow signal \
  --workflow-id <workflow-id> \
  --name retry \
  --input '{"key":"toAccount","value":"account-456"}'
```

## Expected Behavior

1. Workflow starts with invalid account data
2. Activity fails with non-retryable error
3. Workflow waits for manual intervention
4. Admin corrects data and sends retry signal
5. Workflow retries with corrected data and completes

## Key Benefits

- Handles data validation failures gracefully
- Allows human intervention for complex error scenarios
- Maintains workflow state during correction process
- Provides clear audit trail of retry attempts
