---
slug: unit-testing
id: 0pm9icmnyh52
type: challenge
title: 'Exercise 6: Unit Testing Workflows'
teaser: Test a Workflow that waits for a human, with time skipping and mocked Activities.
  No dev server involved.
notes:
- type: text
  contents: |-
    # How do you unit test a Workflow that waits two days for approval?

    You do not wait two days. You do not mock out the wait either.

    The test environment skips the clock forward the instant nothing is
    left to run, so a two-day timer resolves in microseconds and the
    Workflow never knows.
- type: text
  contents: |-
    # No server needed

    `testsuite.WorkflowTestSuite` runs your Workflow in-process, with
    mocked Activities and a fake clock. These tests run in CI with nothing
    installed, and they finish in well under a second.
tabs:
- id: yl8ewaxmm0ow
  title: Code Editor
  type: code
  hostname: workshop
  path: /root/workshop/exercise6
- id: p0huzci8jpqx
  title: Worker
  type: terminal
  hostname: workshop
  workdir: /root/workshop
- id: 20pvcaiq6bei
  title: Terminal
  type: terminal
  hostname: workshop
  workdir: /root/workshop
- id: chk0pam7lb9q
  title: Temporal UI
  type: service
  hostname: workshop
  path: /
  port: 8233
- id: wnaibxyzxxe7
  title: Solution
  type: code
  hostname: workshop
  path: /root/workshop/solution6
- id: leigs0xlwlsv
  title: Network Control Panel
  type: service
  hostname: workshop
  path: /
  port: 5000
difficulty: intermediate
timelimit: 2700
enhanced_loading: null
---

# Unit Testing Workflows

No Worker to start. No dev server to talk to. No TODOs either: `exercise6/` is
the finished transfer Workflow from exercises 1 to 5, and it has no tests. You
write them.

> [!NOTE]
> This is the only exercise that does not need the Temporal dev server. Run
> everything from the [button label="Terminal" background="#444CE7"](tab-2).

# Create the Test File

In the [button label="Code Editor" background="#444CE7"](tab-0), create
`workflow_test.go` next to `workflow.go`. Start with a testify suite that
embeds Temporal's `WorkflowTestSuite`:

```go
package exercise6

import (
    "testing"
    "time"

    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/suite"
    "go.temporal.io/sdk/testsuite"
)

type WorkflowTestSuite struct {
    suite.Suite
    testsuite.WorkflowTestSuite
}

func TestWorkflowTestSuite(t *testing.T) {
    suite.Run(t, new(WorkflowTestSuite))
}
```

# Test 1: Approved Transfer

Each test gets a fresh environment, mocks the Activities it expects, and
schedules the approval Signal on the fake clock:

```go
func (s *WorkflowTestSuite) TestSuccessfulTransfer() {
    env := s.NewTestWorkflowEnvironment()
    env.RegisterWorkflow((&MoneyTransferWorkflow{}).Transfer)

    env.OnActivity(Withdraw, mock.Anything, "account-123", 100.0).Return(nil)
    env.OnActivity(Deposit, mock.Anything, "account-456", 100.0).Return(nil)

    env.RegisterDelayedCallback(func() {
        env.SignalWorkflow("approve", true)
    }, time.Second)

    env.ExecuteWorkflow((&MoneyTransferWorkflow{}).Transfer, TransferRequest{
        FromAccount: "account-123",
        ToAccount:   "account-456",
        Amount:      100.0,
        TransferID:  "transfer-1",
    })

    s.True(env.IsWorkflowCompleted())
    s.NoError(env.GetWorkflowError())

    var result string
    s.NoError(env.GetWorkflowResult(&result))
    s.Equal("Transfer completed successfully", result)
}
```

The mocks replace the real Activities, so their random failures never fire.
`mock.Anything` stands in for the Activity's `context.Context`.

# Test 2: Rejected Transfer

Write this one yourself. Same shape, with three differences:

- Signal `approve` with `false`.
- Mock `Refund` against `account-123` instead of `Deposit`.
- Assert the result is `"Transfer rejected and refunded"`.

# Test 3: Query the State Mid-Flight

Register a second delayed callback that fires *before* the approval and
queries the running Workflow:

```go
env.RegisterDelayedCallback(func() {
    val, err := env.QueryWorkflow("getStatus")
    s.NoError(err)
    var status TransferStatus
    s.NoError(val.Get(&status))
    s.Equal(StatusPending, status)
}, 500*time.Millisecond)
```

# Run the Tests

Click the [button label="Terminal" background="#444CE7"](tab-2):

```bash,run
go test ./exercise6 -v
```

```bash,nocopy
--- PASS: TestWorkflowTestSuite (0.05s)
    --- PASS: TestWorkflowTestSuite/TestQueryStatus (0.00s)
    --- PASS: TestWorkflowTestSuite/TestRejectedTransfer (0.00s)
    --- PASS: TestWorkflowTestSuite/TestSuccessfulTransfer (0.00s)
```

Look at the time. The Workflow waited a full second for approval in every
test, on a clock that never actually ticked.

Run one test while you work on it:

```bash
go test ./exercise6 -v -run 'TestWorkflowTestSuite/TestRejectedTransfer'
```

Click **Check** when at least the approved and rejected tests pass.

# Key Takeaways

- `testsuite.WorkflowTestSuite` runs a Workflow in-process, with no server.
- Time skipping fast-forwards timers and delayed callbacks, so testing a long
  wait costs nothing.
- `env.OnActivity` mocks an Activity for one test. Mock the Activity, not its
  source of randomness.
- `env.QueryWorkflow` inside a delayed callback asserts on intermediate state,
  not just the return value.
