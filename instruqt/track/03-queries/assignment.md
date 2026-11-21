---
slug: queries
type: challenge
title: 'Exercise 3: Money Transfer with Queries'
teaser: Ask a running Workflow what it is doing right now, without touching a database.
notes:
- type: text
  contents: |-
    # Where do you look to find out what a Workflow is doing?

    Not a status table. Not a log aggregator. Not a metrics dashboard.

    The Workflow itself holds the state, in ordinary Go variables. A
    Query reads those variables out of a running execution.
- type: text
  contents: |-
    # Signals in, Queries out

    A Signal changes Workflow state and gets written to history. A Query
    reads state and is never written to history at all.

    That is why a Query handler must not mutate anything and must not
    call an Activity. Break that rule and replay stops matching.
tabs:
- title: Code Editor
  type: code
  hostname: workshop
  path: /root/workshop/exercise3
- title: Worker
  type: terminal
  hostname: workshop
  workdir: /root/workshop
- title: Terminal
  type: terminal
  hostname: workshop
  workdir: /root/workshop
- title: Temporal UI
  type: service
  hostname: workshop
  path: /
  port: 8233
- title: Solution
  type: code
  hostname: workshop
  path: /root/workshop/solution3
- title: Network Control Panel
  type: service
  hostname: workshop
  path: /
  port: 5000
difficulty: basic
timelimit: 2100
enhanced_loading: null
---

# Money Transfer with Queries

Same transfer as exercise 2. This time it can answer questions while it runs.

# The Code

Open the [button label="Code Editor" background="#444CE7"](tab-0). Every TODO
is in `workflow.go`. The Activities, the Worker and the starter are done, and
the starter already queries `getStatus` before it sends the approval.

### Register the Query handler

Register a handler named `getStatus` that returns the `status` variable. It is
the first call in `Transfer` that returns an error, so it declares `err`:

```go
err := workflow.SetQueryHandler(ctx, "getStatus", func() (TransferStatus, error) {
    return status, nil
})
if err != nil {
    return "", err
}
```

### Move the status along

The `TransferStatus` constants are in `models.go`. Set `status` at each step:

| When | Status |
|------|--------|
| Start | `StatusPending` (already set) |
| Approval received | `StatusApproved` |
| Deposit succeeded | `StatusCompleted` |
| Refunded | `StatusCancelled` |
| Any Activity error | `StatusFailed` |

> [!WARNING]
> A Query handler runs during replay. Do not mutate state in one, do not
> execute an Activity from one, and do not block in one. Read a variable and
> return it.

> [!IMPORTANT]
> Delete each `// TODO` comment as you finish it. **Check** looks for leftover
> TODOs before it looks at anything else.

# Start the Worker

Click the [button label="Worker" background="#444CE7"](tab-1):

```bash,run
go run ./exercise3/worker
```

# Run It

Click the [button label="Terminal" background="#444CE7"](tab-2):

```bash,run
go run ./exercise3/starter
```

```bash,nocopy
Status: PENDING
Transfer completed successfully
```

`PENDING` was read out of the running Workflow one second after it started,
while it sat waiting for approval.

# Query From Outside

Start a transfer from the CLI and leave it waiting:

```bash,run
temporal workflow start \
  --type Transfer \
  --task-queue money-transfer-task-queue \
  --workflow-id money-transfer-workflow \
  --input '{"FromAccount":"account-123","ToAccount":"account-456","Amount":75,"TransferID":"transfer-cli"}'
```

Ask it what it is doing:

```bash,run
temporal workflow query --workflow-id money-transfer-workflow --name getStatus
```

Approve it, then ask again:

```bash,run
temporal workflow signal --workflow-id money-transfer-workflow --name approve --input true
sleep 2
temporal workflow query --workflow-id money-transfer-workflow --name getStatus
```

```bash,nocopy
QueryResult  "COMPLETED"
```

A completed Workflow still answers Queries, because the Worker rebuilds its
state by replaying the history.

> [!NOTE]
> A Query is answered by a Worker, not by the server. Leave the Worker running
> in the [button label="Worker" background="#444CE7"](tab-1) tab or the Query
> times out.

Click **Check** when a transfer completes.

# Key Takeaways

- `workflow.SetQueryHandler` exposes Workflow state to any client. Read-only.
- Query results never enter the event history, which is why mutation is banned.
- The state lives in plain Go variables. Temporal rebuilds them by replay.
- Signals write, Queries read. Reach for a Signal when you need to change something.
