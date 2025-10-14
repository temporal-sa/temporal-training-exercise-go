# Temporal Training Exercise - Java Solution

Complete Java implementation for progressive Temporal training exercises.

## Prerequisites
- Java 17+
- Temporal CLI

## Setup

1. Start Temporal dev server with AccountId search attribute:
```bash
temporal server start-dev --search-attribute AccountId=Text
```

2. Build project:
```bash
./gradlew build
```

## Exercise Progression

### Exercise 1: Hello Temporal (30 min)
- Basic Workflow and Activity setup
- Worker registration and execution

### Exercise 2: Money Transfer Basics (45 min)
- Multiple activities (withdraw, deposit, refund)
- Signal-based approval mechanism
- Workflow.await() for conditional waiting
- Basic error handling and compensation

### Exercise 3: Query Handlers (35 min)
- Query methods for workflow state inspection
- Status tracking throughout execution
- External workflow monitoring
- Signal vs Query differences

### Exercise 4: Visibility & Monitoring (30 min)
- Custom Search Attributes (AccountId)
- Upsert Search Attributes from workflows
- Workflow filtering and discovery

### Exercise 5: User Metadata & Activity Summaries (30 min)
- Activity summaries for better observability
- Static workflow summaries
- Enhanced monitoring in Temporal Web UI

### Exercise 6: Testing Strategy (45 min)
- Unit tests with TestWorkflowRule
- Time skipping for fast tests
- Activity mocking with Mockito
- Search attribute registration in tests

### Exercise 7: Manual Activity Retry (40 min)
- Manual retry pattern using signals
- Invalid data handling scenarios
- Disabling automatic retries
- Interactive retry commands

## Running Exercises

### Exercise 1 (Hello Temporal)
```bash
# Start worker
./gradlew execute -PmainClass=com.temporal.training.exercise1.StartWorker
# Run workflow (in another terminal)
./gradlew execute -PmainClass=com.temporal.training.exercise1.StartWorkflow
```

### Exercise 2 (Money Transfer Basics)
```bash
# Start worker
./gradlew execute -PmainClass=com.temporal.training.exercise2.StartWorker
# Run workflow (in another terminal)
./gradlew execute -PmainClass=com.temporal.training.exercise2.StartWorkflow
```

### Exercise 3 (Query Handlers)
```bash
# Start worker
./gradlew execute -PmainClass=com.temporal.training.exercise3.StartWorker
# Run workflow (in another terminal)
./gradlew execute -PmainClass=com.temporal.training.exercise3.StartWorkflow
```

### Exercise 4 (Visibility & Monitoring)
```bash
# Start worker
./gradlew execute -PmainClass=com.temporal.training.exercise4.StartWorker
# Run workflow (in another terminal)
./gradlew execute -PmainClass=com.temporal.training.exercise4.StartWorkflow
```

### Exercise 5 (User Metadata & Activity Summaries)
```bash
# Start worker
./gradlew execute -PmainClass=com.temporal.training.exercise5.StartWorker
# Run workflow (in another terminal)
./gradlew execute -PmainClass=com.temporal.training.exercise5.StartWorkflow
```

### Exercise 6 (Testing Strategy)
```bash
# Run tests
./gradlew test --tests "com.temporal.training.exercise6.MoneyTransferWorkflowTest"
```

### Exercise 7 (Manual Activity Retry)
```bash
# Start worker
./gradlew execute -PmainClass=com.temporal.training.exercise7.StartWorker
# Run workflow (in another terminal)
./gradlew execute -PmainClass=com.temporal.training.exercise7.StartWorkflow
# Send retry signal with corrected data
temporal workflow signal \
  --workflow-id money-transfer-workflow \
  --name retry \
  --input '{"key":"fromAccount","value":"account-123"}'
```

## Running Solutions

Replace `exercise` with `solution` in the class names above to run complete implementations.

## Key Features Implemented

- **Java 17 Records**: Modern data classes
- **Activity Retry**: Configurable retry policies
- **Signal/Query**: Workflow interaction patterns
- **Search Attributes**: Custom filtering
- **User Metadata**: Activity context
- **Compensation**: Saga pattern for rollbacks
- **Testing**: Time-skipping unit tests
- **Error Handling**: Distinguishes activity vs workflow failures
- **Manual Retry**: Signal-based retry for failed activities
