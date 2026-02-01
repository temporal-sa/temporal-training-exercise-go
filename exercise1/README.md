# Exercise 1: Hello Temporal (30 min)

## Learning Objectives
- Understand Temporal Workflow and Activity concepts
- Learn how to register workflows and activities with a Worker
- Execute your first Temporal workflow

## Key Concepts

### Workflow
- Workflows orchestrate activities and contain business logic
- Defined as methods on a struct
- Must be deterministic and use Temporal SDK APIs

### Activity
- Activities handle external interactions (API calls, database operations, etc.)
- Can fail and be retried automatically
- Defined as regular Go functions

### Worker
- Polls task queues for work
- Executes workflow and activity code
- Must register workflow implementations and activity implementations

## Tasks
1. Complete the workflow implementation
2. Complete the activity implementation  
3. Complete the worker setup
4. Run the workflow

## Running the Exercise
1. Start worker: `go run exercise1/worker/main.go`
2. Execute workflow: `go run exercise1/starter/main.go`

## Expected Output
"Hello, Temporal!"
