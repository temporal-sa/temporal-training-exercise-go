---
slug: signals
id: 8ux0rdosx7m8
type: challenge
title: 'Exercise 2: Money Transfer with Signals'
teaser: A transfer that stops and waits for a human to approve it, then deposits or
  refunds.
notes:
- type: text
  contents: |-
    # A Workflow is waiting on a human. Who pays for the wait?

    The money left the source account. Now a person has to approve the
    deposit. That person is at lunch.

    No polling loop. No cron job. No row in a "pending approvals" table.
    The Workflow just waits, and the wait costs you nothing.
- type: text
  contents: |-
    # Did you know?

    A Workflow blocked in `workflow.Await` holds no goroutine on your
    Worker and no memory. Temporal evicts it and brings it back when the
    Signal arrives. A wait of ten seconds and a wait of ten months cost
    the same.
tabs:
- id: pciscnxhsolh
  title: Code Editor
  type: code
  hostname: workshop
  path: /root/workshop/exercise2
- id: pmvo21k8ohhg
  title: Worker
  type: terminal
  hostname: workshop
  workdir: /root/workshop
- id: hsmvmmwbscvz
  title: Terminal
  type: terminal
  hostname: workshop
  workdir: /root/workshop
- id: bt3qzyx1hgae
  title: Temporal UI
  type: service
  hostname: workshop
  path: /
  port: 8233
- id: s0aj0y53lllr
  title: Solution
  type: code
  hostname: workshop
  path: /root/workshop/solution2
- id: 5mr88fuvxh7j
  title: Network Control Panel
  type: service
  hostname: workshop
  path: /
  port: 5000
difficulty: basic
timelimit: 2700
enhanced_loading: null
---

# Money Transfer with Signals

Withdraw. Wait for a human. Deposit if they said yes, refund if they said no.

# The Code

Open the [button label="Code Editor" background="#444CE7"](tab-0). Four files
carry TODOs. `models.go` is already done.

### `activities.go`

`Deposit` and `Refund` are written. Finish `Withdraw` the same way `Deposit`
does it: log, then fail 10% of the time with `rand.Float64() < 0.1`.

The random failure is deliberate. It is what makes Temporal's automatic
Activity retries visible in the next step.

### `workflow.go`

- `ActivityOptions` with a 5 second `StartToCloseTimeout`.
- A Signal channel named `approve`, drained by a goroutine started with
  `workflow.Go`. Never a bare `go` statement inside a Workflow.
- Execute `Withdraw`.
- Block until the approval arrives.
- Execute `Deposit` on approval, `Refund` on rejection, and return a message.

The Signal half looks like this:

```go
var approved, approvalReceived bool

approvalChannel := workflow.GetSignalChannel(ctx, "approve")
workflow.Go(ctx, func(ctx workflow.Context) {
    approvalChannel.Receive(ctx, &approved)
    approvalReceived = true
})

// ... Withdraw ...

workflow.Await(ctx, func() bool { return approvalReceived })
```

### `worker/main.go`

Register the Workflow and all three Activities.

### `starter/main.go`

Uncomment the two lines that wait two seconds and send the `approve` Signal.

> [!IMPORTANT]
> Delete each `// TODO` comment as you finish it. **Check** looks for leftover
> TODOs before it looks at anything else.

# Start the Worker

Click the [button label="Worker" background="#444CE7"](tab-1):

```bash,run
go run ./exercise2/worker
```

# Run It

Click the [button label="Terminal" background="#444CE7"](tab-2):

```bash,run
go run ./exercise2/starter
```

```bash,nocopy
Transfer completed successfully
```

Watch the Worker tab. If an Activity hit its simulated failure you will see
the error, then the same Activity running again. You did not write that retry.

# Signal It Yourself

The starter is not special. Any client that can reach the server can start a
Workflow and move it along, including the CLI. Start a transfer with no
starter at all:

```bash,run
temporal workflow start \
  --type Transfer \
  --task-queue money-transfer-task-queue \
  --workflow-id money-transfer-workflow \
  --input '{"FromAccount":"account-123","ToAccount":"account-456","Amount":250,"TransferID":"transfer-cli"}'
```

It withdraws and then waits. Nobody is going to approve it, so reject it:

```bash,run
temporal workflow signal \
  --workflow-id money-transfer-workflow \
  --name approve \
  --input false
```

```bash,run
temporal workflow result --workflow-id money-transfer-workflow
```

```bash,nocopy
Result  "Transfer rejected and refunded"
```

> [!NOTE]
> Every exercise reuses the Workflow ID `money-transfer-workflow`. If a run is
> stuck waiting and you want a clean start, terminate it:
> `temporal workflow terminate --workflow-id money-transfer-workflow`.
> Otherwise the starter attaches to the run that is still going instead of
> starting a new one.

# Look at the Wait

Click the [button label="Temporal UI" background="#444CE7"](tab-3). Open your
Workflow, then the **Event History**. Find `WorkflowExecutionSignaled`. Above
it, the Workflow was doing nothing at all, and it was not consuming a Worker
while it did.

Click **Check** when a transfer completes.

# Key Takeaways

- `workflow.GetSignalChannel` plus `workflow.Go` receives Signals into
  Workflow state.
- `workflow.Await` blocks on that state without burning a Worker.
- Failed Activities retry on their own. The default retry policy is already on.
- The Signal is recorded in history, so replay makes the same decision.
