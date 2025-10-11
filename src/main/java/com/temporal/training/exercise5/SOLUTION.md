# Exercise 5 Solution: User Metadata & Activity Summaries

## Solution Implementation

### 1. Activity Summaries in MoneyTransferWorkflowImpl.java

Replace the `ActivityFactory` methods with:

```java
private static class ActivityFactory {
    private static BankingActivities createActivity(String summary) {
        return Workflow.newActivityStub(
            BankingActivities.class,
            ActivityOptions.newBuilder()
                .setStartToCloseTimeout(Duration.ofSeconds(5))
                .setSummary(summary)
                .build()
        );
    }
    
    static BankingActivities createWithdrawActivity(String accountId) {
        return createActivity("Withdrawing funds from account " + accountId);
    }
    
    static BankingActivities createDepositActivity(String accountId) {
        return createActivity("Depositing funds to account " + accountId);
    }
    
    static BankingActivities createRefundActivity(String accountId) {
        return createActivity("Refunding funds to account " + accountId);
    }
}
```

### 2. Static Summary in StartWorkflow.java

Update the WorkflowOptions:

```java
String staticSummaryDetails = "Money transfer workflow for " + request.fromAccount() + " to " + request.toAccount();

WorkflowOptions options = WorkflowOptions.newBuilder()
    .setTaskQueue(StartWorker.TASK_QUEUE)
    .setWorkflowId("money-transfer-" + System.currentTimeMillis())
    .setStaticSummary(staticSummaryDetails)
    .build();
```

## Key Points

- **Activity Summary**: Provides runtime context for each activity execution
- **Static Summary**: Set once at workflow start, describes the overall workflow purpose
- **Observability**: Both summaries appear in Temporal Web UI for better monitoring
- **Best Practice**: Use descriptive, contextual summaries that help with debugging and monitoring