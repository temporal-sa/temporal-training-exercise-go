---
slug: activity-summaries
id: 8wukhk3fa2ae
type: challenge
title: 'Exercise 5: Activity Summaries and Metadata'
teaser: Make the event history readable. Attach a one-line summary to every Activity.
notes:
- type: text
  contents: |-
    # Three Activities named Withdraw, Deposit and Refund. Which account?

    You are on a call. Someone shares an event history. It says
    `Withdraw`, `Deposit`, `Refund`.

    It does not say whose money, or how much. That detail was in the
    payload, and nobody wants to click into three payloads on a call.
- type: text
  contents: |-
    # Did you know?

    Activity summaries are user metadata. They render in the Web UI beside
    the event, they cost nothing at runtime, and they are not part of
    Workflow logic. Changing a summary string does not break replay.
tabs:
- id: sxpztdz6oky4
  title: Code Editor
  type: code
  hostname: workshop
  path: /root/workshop/exercise5
- id: watqzlnzfp4a
  title: Worker
  type: terminal
  hostname: workshop
  workdir: /root/workshop
- id: i3msxhhx60np
  title: Terminal
  type: terminal
  hostname: workshop
  workdir: /root/workshop
- id: 5pru0plgshgt
  title: Temporal UI
  type: service
  hostname: workshop
  path: /
  port: 8233
- id: ud0gdn6pcg9n
  title: Solution
  type: code
  hostname: workshop
  path: /root/workshop/solution5
- id: pgo9u7ltgkcu
  title: Network Control Panel
  type: service
  hostname: workshop
  path: /
  port: 5000
difficulty: basic
timelimit: 1800
enhanced_loading: null
---

# Activity Summaries and Metadata

The Workflow already works. This exercise makes it legible to whoever debugs
it at 3am.

# The Code

Open the [button label="Code Editor" background="#444CE7"](tab-0). Every TODO
is in `workflow.go`.

Each Activity already gets its own context from `workflow.WithActivityOptions`.
Add a `Summary` to each of the three `ActivityOptions` and interpolate the
account into it:

```go
withdrawCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
    StartToCloseTimeout: 5 * time.Second,
    Summary:             fmt.Sprintf("Withdrawing funds from account %s", request.FromAccount),
})
```

Do the same for the deposit (`"Depositing funds to account %s"`, to
`request.ToAccount`) and the refund (`"Refunding funds to account %s"`, to
`request.FromAccount`). `fmt` is already imported.

Delete every `// TODO` comment when you are done.

# Run It

Click the [button label="Worker" background="#444CE7"](tab-1):

```bash,run
go run ./exercise5/worker
```

Click the [button label="Terminal" background="#444CE7"](tab-2):

```bash,run
go run ./exercise5/starter
```

> [!IMPORTANT]
> If the Worker was already running from before your edit, stop it with
> `Ctrl+C` and start it again. A Go Worker runs the binary it was compiled
> from, and an old binary schedules Activities without summaries.

# See the Difference

Click the [button label="Temporal UI" background="#444CE7"](tab-3). Click your
Workflow at the top of the list. Refresh if it is not there.

Open the **Event History** and look at the Activity rows: each one now says
which account it touched and which direction the money went. Compare that
against a run from exercise 4 further down the list.

# Go Further: Summarise the Whole Execution

Activities are not the only thing that can carry a summary. In
`starter/main.go`, add `StaticSummary` to the `StartWorkflowOptions`:

```go
options := client.StartWorkflowOptions{
    ID:            "money-transfer-workflow",
    TaskQueue:     "money-transfer-task-queue",
    StaticSummary: fmt.Sprintf("Money transfer from %s to %s", request.FromAccount, request.ToAccount),
}
```

Run the starter again. The Workflows list now shows a summary next to the
execution, and nobody has to open it to know what it is.

Click **Check** when the summaries show up.

# Key Takeaways

- `Summary` on `ActivityOptions` labels one Activity in the UI.
- `StaticSummary` on `StartWorkflowOptions` labels the whole execution.
- Both are user metadata. They do not affect Workflow logic or replay.
- Keep summaries short and specific. They are read at a glance, under pressure.
