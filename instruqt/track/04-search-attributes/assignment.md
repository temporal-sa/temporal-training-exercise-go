---
slug: search-attributes
id: iqkve574uwgd
type: challenge
title: 'Exercise 4: Money Transfer with Search Attributes'
teaser: Tag a Workflow with a business identifier, then find it by that identifier
  instead of its Workflow ID.
notes:
- type: text
  contents: |-
    # Support asks about account-123. You have 40,000 Workflows. Now what?

    You do not have the Workflow ID. You have an account number.

    One call makes that account number queryable across every execution
    the cluster has ever seen.
- type: text
  contents: |-
    # One call. Fifteen minutes.

    The shortest exercise in the workshop. Custom search attributes have to
    be registered on the cluster before a Workflow can set one, so
    `AccountId=Text` is already registered on your dev server.

    Try it without that registration and the Workflow Task fails outright.
tabs:
- id: g8bnywke3imj
  title: Code Editor
  type: code
  hostname: workshop
  path: /root/workshop/exercise4
- id: gfo7emtguf22
  title: Worker
  type: terminal
  hostname: workshop
  workdir: /root/workshop
- id: yrw6rrgdonyt
  title: Terminal
  type: terminal
  hostname: workshop
  workdir: /root/workshop
- id: wvo2quu1t9vw
  title: Temporal UI
  type: service
  hostname: workshop
  path: /
  port: 8233
- id: ygwefptri5hr
  title: Solution
  type: code
  hostname: workshop
  path: /root/workshop/solution4
- id: y9sxxaw99o2w
  title: Network Control Panel
  type: service
  hostname: workshop
  path: /
  port: 5000
difficulty: basic
timelimit: 1800
enhanced_loading: null
---

# Money Transfer with Search Attributes

One call. A Workflow you can find by account number.

# The Code

Open the [button label="Code Editor" background="#444CE7"](tab-0) and find the
TODOs in `workflow.go`, at the top of `Transfer`.

Set the `AccountId` search attribute to `request.FromAccount`. Do it before
the first Activity so the tag is on the execution from the start:

```go
err := workflow.UpsertSearchAttributes(ctx, map[string]interface{}{
    "AccountId": request.FromAccount,
})
if err != nil {
    return "", err
}
```

This call declares `err`, so change the later `err := workflow.SetQueryHandler(...)`
to `err = ...`.

> [!IMPORTANT]
> `AccountId=Text` is already registered on the dev server in this sandbox.
> Setting an unregistered attribute fails the Workflow Task, which shows up
> as a Workflow that starts and then goes nowhere.

Delete both `// TODO` comments when you are done.

# Run It

Click the [button label="Worker" background="#444CE7"](tab-1):

```bash,run
go run ./exercise4/worker
```

Click the [button label="Terminal" background="#444CE7"](tab-2):

```bash,run
go run ./exercise4/starter
```

The starter uses `account-123` as the source account.

# Find It By Account

From the [button label="Terminal" background="#444CE7"](tab-2):

```bash,run
temporal workflow list --query 'AccountId="account-123"'
```

Now filter on an account that never sent anything:

```bash,run
temporal workflow list --query 'AccountId="account-999"'
```

Empty. The filter is doing real work.

# See It In the UI

Click the [button label="Temporal UI" background="#444CE7"](tab-3). Click your
Workflow, then look at the summary panel. `AccountId` is listed there with its
value. The Event History also has an `UpsertWorkflowSearchAttributes` event.

Back on the Workflows list, put this in the search box:

```bash,nocopy
AccountId="account-123"
```

Click **Check** when the query returns your Workflow.

# Key Takeaways

- `workflow.UpsertSearchAttributes` tags an execution with business data.
- Custom attributes must be registered on the cluster first. The dev server
  takes `--search-attribute Name=Type`.
- Search attributes are for finding Workflows, not for storing Workflow state.
  Keep the payload small.
- The same query syntax works in the CLI, the Web UI, and the client API.
