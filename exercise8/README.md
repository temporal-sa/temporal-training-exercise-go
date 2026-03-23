# Exercise 8: Versioning

## Learning Objectives
- Understand Temporal Workflow and Activity concepts
- Learn how to register workflows and activities with a Worker
- Execute your first Temporal workflow

## Exercises

### Workflow Type Versioning
These must be run from the `01-workflow-type-versioning` directory

1. Start the Worker: `go run ./worker`
2. Start a `Greet` workflow from the UI, or via the CLI: `temporal workflow start --task-queue hello-task-queue --type Greet --input "\"Temporal\""`

#### Introduce a non-deterministic error
1. Make a non-deterministic to the Workflow. 
2. Restart the Worker. 
3. View the non-deterministic error in the Temporal UI.

#### Rollback the change
1. Comment out the non-deterministic change to the Workflow.
2. Restart the Worker.
3. View the Workflow in the UI. The error will resolve.

#### Create a new Workflow Type
1. Create a copy of the Workflow function `Greet` with the name `GreetV2`.
2. Add the non-deterministic change to `GreetV2`.
3. Register `GreetV2` on the worker.
2. Restart the worker.
3. Start a `GreetV2` workflow from the UI, or via the CLI: `temporal workflow start --task-queue hello-task-queue --type GreetV2 --input "\"Temporal\""`


### Task Queue Versioning
These must be run from the `02-task-queue-versioning` directory

1. Start the Worker: `go run ./worker`
2. Start a `Greet` workflow from the UI, or via the CLI: `temporal workflow start --task-queue hello-task-queue --type Greet --input "\"Temporal\""`

#### Introduce a non-deterministic error
1. Make a non-deterministic to the Workflow. 
2. Restart the Worker. 
3. View the non-deterministic error in the Temporal UI.

#### Rollback the change
1. Comment out the non-deterministic change to the Workflow.
2. Restart the Worker.
3. View the Workflow in the UI. The error will resolve.

#### Create a new Worker and Workflow
1. Create a copy of the Workflow function `Greet` with the name `GreetV2`.
2. Create a copy of the Worker directory and name it `./Worker2`
3. Change `./Worker2`'s task queue to `"hello-task-queue-v2"`.
4. Start `Worker2` in a new terminal.
5. Start a `Greet` workflow from the UI, or via the CLI: `temporal workflow start --task-queue hello-task-queue-v2 --type GreetV2 --input "\"Temporal\""`

 
### Patching API
These must be run from the `03-patching` directory

1. Start the Worker: `go run ./worker`
2. Start a `Greet` workflow from the UI, or via the CLI: `temporal workflow start --task-queue hello-task-queue --type Greet --input "\"Temporal\""`
3. (Optional) Repeat the steps from earlier to introduce a non-deterministic change and roll it back.
4. Use the patching API to add the `Lowercase` activity to the workflow.
5. Start a `Greet` workflow from the UI, or via the CLI: `temporal workflow start --task-queue hello-task-queue --type Greet --input "\"Temporal\""`


### Worker Versioning
Instructions coming soon!


### Worker Controller
Instructions coming soon!