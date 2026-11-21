---
slug: manual-retry
id: kyiiqy5jspdi
type: challenge
title: 'Exercise 7: Manual Activity Retry'
teaser: An Activity fails on bad data. Automatic retry cannot help. Fix the data with
  a Signal and let it continue.
notes:
- type: text
  contents: |-
    # The account number is wrong. How many times should Temporal retry?

    Automatic retry is the right answer when the failure is transient.
    A network blip, a rate limit, a restarted database.

    A typo in an account number is not transient. Retrying it a hundred
    times gets you the same error a hundred times.
- type: text
  contents: |-
    # Did you know?

    `RetryPolicy{MaximumAttempts: 1}` tells Temporal to run an Activity
    once and hand any failure straight back to the Workflow.

    The Workflow can then park, wait for a human to send corrected data,
    and run the Activity again. The withdrawal that already succeeded
    stays succeeded.
tabs:
- id: e4urosmakvht
  title: Code Editor
  type: code
  hostname: workshop
  path: /root/workshop/exercise7
- id: lonidraxsso2
  title: Worker
  type: terminal
  hostname: workshop
  workdir: /root/workshop
- id: eynlg9rztnvr
  title: Terminal
  type: terminal
  hostname: workshop
  workdir: /root/workshop
- id: kyk10ackog5x
  title: Temporal UI
  type: service
  hostname: workshop
  path: /
  port: 8233
- id: 48fdcjskwynl
  title: Solution
  type: code
  hostname: workshop
  path: /root/workshop/solution7
- id: lie2qx3ismhx
  title: Network Control Panel
  type: service
  hostname: workshop
  path: /
  port: 5000
- id: v5mcxyoaitph
  title: Signals
  type: terminal
  hostname: workshop
  workdir: /root/workshop
difficulty: intermediate
timelimit: 2400
enhanced_loading: null
---

# Manual Activity Retry

The transfer stops on a bad account number and waits for you to fix it. The
withdrawal that already went through does not get undone.

> [!IMPORTANT]
> This exercise has a seventh tab. Use
> [button label="Signals" background="#444CE7"](tab-6) to send Signals while
> the starter is still waiting in the
> [button label="Terminal" background="#444CE7"](tab-2).

# The Code

Open the [button label="Code Editor" background="#444CE7"](tab-0). Every TODO
is in `workflow.go`. `RetryUpdate` and `StatusRetrying` are already in
`models.go`.

### Turn off automatic retry

```go
ao := workflow.ActivityOptions{
    StartToCloseTimeout: 5 * time.Second,
    RetryPolicy: &temporal.RetryPolicy{
        MaximumAttempts: 1,
    },
}
```

`MaximumAttempts: 1` is the whole trick. Without it Temporal retries the bad
account forever and the Workflow never gets a chance to intervene.

### Listen for the retry Signal

A `retry` Signal carries a `RetryUpdate` with a `Key` and a `Value`. Receive
it in a loop, because a transfer can need more than one correction:

```go
retryChannel := workflow.GetSignalChannel(ctx, "retry")
workflow.Go(ctx, func(ctx workflow.Context) {
    for {
        var update RetryUpdate
        retryChannel.Receive(ctx, &update)
        switch update.Key {
        case "fromAccount":
            updatedRequest.FromAccount = update.Value
        case "toAccount":
            updatedRequest.ToAccount = update.Value
        case "amount":
            if amt, err := strconv.ParseFloat(update.Value, 64); err == nil {
                updatedRequest.Amount = amt
            }
        }
        retryRequested = true
    }
})
```

### Wrap Withdraw and Deposit in a retry loop

Run the Activity. On success, break. On failure, park until a retry Signal
arrives, then go round again with the corrected request:

```go
for {
    err = workflow.ExecuteActivity(ctx, Deposit, updatedRequest.ToAccount, updatedRequest.Amount).Get(ctx, nil)
    if err == nil {
        break
    }
    status = StatusRetrying
    retryRequested = false
    workflow.Await(ctx, func() bool { return retryRequested })
}
```

Do the same for `Withdraw`. Delete every `// TODO` comment when you are done.

# Break It On Purpose

The shipped Activities only fail at random, 10% of withdrawals and 5% of
deposits. To see the retry path every time, make a bad account fail for sure.
In `activities.go`, replace the random check in `Deposit`:

```go
if account == "invalid-account" {
    return fmt.Errorf("deposit failed - invalid account")
}
```

Then in `starter/main.go`, change the target:

```go
ToAccount:   "invalid-account",
```

# Start the Worker

Click the [button label="Worker" background="#444CE7"](tab-1):

```bash,run
go run ./exercise7/worker
```

# Run It

Click the [button label="Terminal" background="#444CE7"](tab-2):

```bash,run
go run ./exercise7/starter
```

The starter sends the approval after two seconds. Withdraw succeeds. Deposit
hits `invalid-account` and fails once, with no automatic retry. The starter is
now waiting for a result that will not arrive until you fix the data.

# Check Where It Parked

Click the [button label="Signals" background="#444CE7"](tab-6):

```bash,run
temporal workflow query --workflow-id money-transfer-workflow --name getStatus
```

```bash,nocopy
QueryResult  "RETRYING"
```

That is the Workflow parked mid-transfer, holding the state of a withdrawal
that already succeeded.

# Fix It

Still in the [button label="Signals" background="#444CE7"](tab-6), send the
corrected account:

```bash,run
temporal workflow signal \
  --workflow-id money-transfer-workflow \
  --name retry \
  --input '{"Key":"toAccount","Value":"account-456"}'
```

Back in the [button label="Terminal" background="#444CE7"](tab-2), the starter
prints `Transfer completed successfully`. The deposit ran against the
corrected account. Withdraw never ran twice.

Open the [button label="Temporal UI" background="#444CE7"](tab-3) and look at
the event history: one `ActivityTaskFailed` for the deposit, then the retry
Signal, then a fresh `ActivityTaskScheduled` for the same Activity.

Click **Check** when the transfer completes.

# Key Takeaways

- `MaximumAttempts: 1` stops automatic retry for failures that a data change,
  not time, will fix.
- A Workflow can catch the failure, park on `workflow.Await`, and resume.
- Work already completed stays completed. That is the difference from
  restarting the process.
- Signals carry the corrected data in, so no redeploy and no manual database edit.
