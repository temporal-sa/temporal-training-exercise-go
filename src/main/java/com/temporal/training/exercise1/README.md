# Exercise 1: Hello Temporal (30 min)

## Learning Objectives
- Understand Temporal Workflow and Activity concepts
- Learn how to register workflows and activities with a Worker
- Execute your first Temporal workflow

## Key Concepts

### Workflow
- **@WorkflowInterface**: Marks an interface as a Temporal workflow
- **@WorkflowMethod**: Defines the main workflow entry point
- Workflows orchestrate activities and contain business logic

### Activity
- **@ActivityInterface**: Marks an interface as containing activities
- **@ActivityMethod**: Defines methods that perform actual work
- Activities handle external interactions (API calls, database operations, etc.)

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
1. Start worker: `./gradlew execute -PmainClass=com.temporal.training.exercise1.StartWorker`
2. Execute workflow: `./gradlew execute -PmainClass=com.temporal.training.exercise1.StartWorkflow`

## Expected Output
"Hello, Temporal!"
