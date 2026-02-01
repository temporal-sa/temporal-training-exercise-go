# Exercise 3: Flow Control with Query Handlers (35 min)

## Learning Objectives
- Implement Query handlers for workflow state inspection
- Track workflow status throughout execution
- Query workflow state from external clients
- Understand the difference between Signals and Queries

## Tasks

### 1. Implement Query Handler (10 min)
- Add `workflow.SetQueryHandler(ctx, "getStatus", ...)` for "getStatus" query
- Return current status and nil error from the handler function
- Handle any errors from SetQueryHandler

### 2. Add Status Tracking (10 min)
- Update status to `StatusApproved` when approval received
- Update status to `StatusCompleted` when transfer succeeds
- Update status to `StatusCancelled` when rejected
- Update status to `StatusFailed` on exceptions

## Key Concepts
- **Query Handlers**: Read-only operations to inspect workflow state
- **Status Tracking**: Maintaining workflow state for external visibility
- **Signal vs Query**: Signals modify state, Queries only read state

## Testing
1. Start the worker: `go run exercise3/worker/main.go`
2. Run the workflow: `go run exercise3/starter/main.go`
3. Observe status changes through query outputs

### Query Using CLI
You can also query the workflow status using the Temporal CLI:
```bash
temporal workflow query \
  --workflow-id money-transfer-workflow \
  --type getStatus
```

## Expected Output
- Initial status: PENDING
- Status after approval: APPROVED → COMPLETED
- Final status confirmation
