---
slug: hello-temporal
id: khy45mzgj60a
type: challenge
title: 'Exercise 1: Hello Temporal'
teaser: Write your first Workflow, Activity and Worker in Go. Run it and read the
  event history.
notes:
- type: text
  contents: |-
    # What does a Workflow actually run on?

    You write a function. You never call it. Something else picks it up,
    runs it, and records every step.

    That something is a Worker, and it is polling a task queue right now
    waiting for code you have not written yet.
- type: text
  contents: |-
    # Already running in your sandbox

    A Temporal dev server on 127.0.0.1:7233, with the Web UI on port 8233.
    The AccountId search attribute is already registered, so exercise 4
    works when you get there.

    Go 1.24 is installed and the Temporal SDK is already compiled into the
    build cache, so `go run` starts in seconds.

    The Solution tab has finished code for every exercise. Use it when you
    are stuck, not before.
tabs:
- id: yh2x7h5hmiwh
  title: Code Editor
  type: code
  hostname: workshop
  path: /root/workshop/exercise1
- id: ykzyserbdx0h
  title: Worker
  type: terminal
  hostname: workshop
  workdir: /root/workshop
- id: 13x381eycgnt
  title: Terminal
  type: terminal
  hostname: workshop
  workdir: /root/workshop
- id: m1x45s1vsgr1
  title: Temporal UI
  type: service
  hostname: workshop
  path: /
  port: 8233
- id: 6krwcnaqmy0s
  title: Solution
  type: code
  hostname: workshop
  path: /root/workshop/solution1
- id: wunyjp9xsd4t
  title: Network Control Panel
  type: service
  hostname: workshop
  path: /
  port: 5000
difficulty: basic
timelimit: 1800
enhanced_loading: null
---

# Hello Temporal

Three files. Five TODOs. A greeting that survives anything.

> [!NOTE]
> **Sandbox Notes:**
> - [button label="Code Editor" background="#444CE7"](tab-0) edits the files in `exercise1/`
> - [button label="Worker" background="#444CE7"](tab-1) runs the Worker
> - [button label="Terminal" background="#444CE7"](tab-2) starts Workflows
> - [button label="Temporal UI" background="#444CE7"](tab-3) is the Temporal Web UI
> - [button label="Solution" background="#444CE7"](tab-4) has the finished code
> - [button label="Network Control Panel" background="#444CE7"](tab-5) toggles outbound HTTP
>
> The blue buttons above are clickable. Click any to jump to that tab.
> Instruqt saves editor changes for you. There is no save step.

# The Code

Open the [button label="Code Editor" background="#444CE7"](tab-0). Three files
carry the TODOs:

| File | What goes in it |
|------|-----------------|
| `activity.go` | `CreateGreeting` returns `"Hello, <name>!"` |
| `workflow.go` | `ActivityOptions` with a 10 second `StartToCloseTimeout`, then execute `CreateGreeting` and return its result |
| `worker/main.go` | Register the Workflow and the Activity on `hello-task-queue` |

`starter/main.go` is already done. Read it anyway: it is the client that asks
the server to run `Greet`.

The Activity is where side effects live. The Workflow only orchestrates. The
Worker is the process that runs both.

In `workflow.go` the Activity call looks like this:

```go
ao := workflow.ActivityOptions{StartToCloseTimeout: 10 * time.Second}
ctx = workflow.WithActivityOptions(ctx, ao)

var result string
err := workflow.ExecuteActivity(ctx, CreateGreeting, name).Get(ctx, &result)
```

> [!IMPORTANT]
> Delete each `// TODO` comment as you finish it. **Check** looks for leftover
> TODOs before it looks at anything else.

# Start the Worker

Click the [button label="Worker" background="#444CE7"](tab-1) terminal:

```bash,run
go run ./exercise1/worker
```

It stays in the foreground and polls. Leave it running. If you edit the
Workflow or the Activity later, stop it with `Ctrl+C` and start it again: a Go
Worker runs the code it was compiled with.

# Run It

Click the [button label="Terminal" background="#444CE7"](tab-2):

```bash,run
go run ./exercise1/starter
```

```bash,nocopy
Hello, Temporal!
```

# Read the History

Click the [button label="Temporal UI" background="#444CE7"](tab-3). Click
`greeting-workflow` at the top of the list. Hit refresh if it is not there yet.

Open the **Event History**. Look for `ActivityTaskScheduled`,
`ActivityTaskStarted` and `ActivityTaskCompleted`. Those three events are the
whole point. Temporal wrote down that the Activity ran and what it returned.
Replay reads that back instead of calling the Activity again.

The CLI reads the same history:

```bash,run
temporal workflow show --workflow-id greeting-workflow
```

Click **Check** when the greeting prints.

# Key Takeaways

- A Workflow is a function that takes `workflow.Context` first. An Activity is
  a plain Go function.
- Activities do the side effects. Workflows orchestrate and stay deterministic.
- A Worker only runs what you register with it, on the task queue you name.
- Every Activity result lands in the event history, which is what makes replay work.
