# Temporal Training Exercise - Go Solution

Complete Go implementation for progressive Temporal training exercises.

## Prerequisites
- Go 1.21+
- Temporal CLI

### Installing Temporal CLI

**macOS (Homebrew):**
```bash
brew install temporal
```

**Linux:**
```bash
curl -sSf https://temporal.download/cli.sh | sh
```

**Windows:**
```powershell
scoop install temporal-cli
```

**Or download directly:**
Visit https://github.com/temporalio/cli/releases and download the appropriate binary for your platform.

## Setup

1. Start Temporal dev server with AccountId search attribute:
```bash
temporal server start-dev --search-attribute AccountId=Text
```

2. Install dependencies:
```bash
go mod download
```

## Exercise Progression

### [Exercise 1: Hello Temporal](exercise1/) (30 min)
- Basic Workflow and Activity setup
- Worker registration and execution

### [Exercise 2: Money Transfer Basics](exercise2/) (45 min)
- Multiple activities (withdraw, deposit, refund)
- Signal-based approval mechanism
- workflow.Await() for conditional waiting
- Basic error handling and compensation

### [Exercise 3: Query Handlers](exercise3/) (35 min)
- Query methods for workflow state inspection
- Status tracking throughout execution
- External workflow monitoring
- Signal vs Query differences

### [Exercise 4: Visibility & Monitoring](exercise4/) (30 min)
- Custom Search Attributes (AccountId)
- Upsert Search Attributes from workflows
- Workflow filtering and discovery

### [Exercise 5: User Metadata & Activity Summaries](exercise5/) (30 min)
- Activity summaries for better observability
- Enhanced monitoring in Temporal Web UI

### [Exercise 6: Testing Strategy](exercise6/) (45 min)
- Unit tests with testsuite
- Time skipping for fast tests
- Activity mocking with testify/mock
- Search attribute registration in tests

### [Exercise 7: Manual Activity Retry](exercise7/) (40 min)
- Manual retry pattern using signals
- Invalid data handling scenarios
- Disabling automatic retries
- Interactive retry commands

## Running Exercises

### Exercise 1 (Hello Temporal)
```bash
# Start worker
go run solution1/worker/main.go
# Run workflow (in another terminal)
go run solution1/starter/main.go
```

### Exercise 2 (Money Transfer Basics)
```bash
# Start worker
go run solution2/worker/main.go
# Run workflow (in another terminal)
go run solution2/starter/main.go
```

### Exercise 3 (Query Handlers)
```bash
# Start worker
go run solution3/worker/main.go
# Run workflow (in another terminal)
go run solution3/starter/main.go
```

### Exercise 4 (Visibility & Monitoring)
```bash
# Start worker
go run solution4/worker/main.go
# Run workflow (in another terminal)
go run solution4/starter/main.go
```

### Exercise 5 (User Metadata & Activity Summaries)
```bash
# Start worker
go run solution5/worker/main.go
# Run workflow (in another terminal)
go run solution5/starter/main.go
```

### Exercise 6 (Testing Strategy)
```bash
# Run tests
go test ./solution6 -v
```

### Exercise 7 (Manual Activity Retry)
```bash
# Start worker
go run solution7/worker/main.go
# Run workflow (in another terminal)
go run solution7/starter/main.go
# Or send retry signal manually
temporal workflow signal \
  --workflow-id money-transfer-workflow \
  --name retry \
  --input '{"Key":"fromAccount","Value":"account-123"}'
```

## Key Features Implemented

- **Go Structs**: Modern data structures
- **Activity Retry**: Configurable retry policies
- **Signal/Query**: Workflow interaction patterns
- **Search Attributes**: Custom filtering
- **Activity Summaries**: Enhanced observability
- **Compensation**: Saga pattern for rollbacks
- **Testing**: Time-skipping unit tests with testify
- **Error Handling**: Distinguishes activity vs workflow failures
- **Manual Retry**: Signal-based retry for failed activities

## Project Structure

```
.
├── solution1/          # Hello Temporal
├── solution2/          # Money Transfer Basics
├── solution3/          # Query Handlers
├── solution4/          # Visibility & Monitoring
├── solution5/          # Activity Summaries
├── solution6/          # Testing Strategy
├── solution7/          # Manual Activity Retry
└── go.mod
```

Each solution contains:
- `workflow.go` - Workflow implementation
- `activities.go` - Activity implementations
- `models.go` - Data structures
- `worker/main.go` - Worker startup
- `starter/main.go` - Workflow starter
- `workflow_test.go` - Unit tests (solution6)
