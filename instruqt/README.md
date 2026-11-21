# Temporal in Practice: Instruqt Track (Go)

The seven exercises in this repo, packaged as an Instruqt lab. Attendees get a
container sandbox with a Temporal dev server already running, Go 1.24 with the
SDK pre-compiled into the build cache, the code editor on the exercise
package, and the finished solution one tab away.

```
instruqt/
├── track/       7 challenge directories + track.yml   ->  instruqt track push
└── sandbox/     config.yml + scripts/                 ->  instruqt sandbox push
                 proxy/                                ->  git push
```

Ported from the Python track in
[temporal-training-exercise-python](https://github.com/temporal-sa/temporal-training-exercise-python/tree/main/instruqt),
which has eight challenges. This repo has no versioning exercise, so the Go
track stops at manual retry.

## Three ways a change ships

| What you edited | How it ships | Live at |
|---|---|---|
| `track/` — assignments, lifecycle scripts, `track.yml` | `instruqt track push` | next lab start |
| `sandbox/config.yml`, `sandbox/scripts/` | `instruqt sandbox push` then `publish` | next sandbox provision |
| `sandbox/proxy/`, `exercise*/`, `solution*/` | `git push` to `WORKSHOP_REF` | next sandbox provision |

The track and the sandbox preset are separate Instruqt entities:
`instruqt track push` does not touch the preset, and `instruqt sandbox push`
does not touch the track. The track carries no `config.yml` and no
`track_scripts/` — `track.yml` names the preset instead:

```yaml
sandbox_preset: temporal-in-practice-sandbox-go
```

> [!IMPORTANT]
> The third row is the one that catches people. Exercise code and the proxy
> files are part of neither push. They reach the sandbox through the
> `git clone` that `sandbox/scripts/setup-workshop` runs at provision time, so
> an edit is live only once it is pushed to the ref `WORKSHOP_REF` names
> (default `main`). Pushing the track ten times will not pick it up.

Authenticate once with `instruqt auth login`.

## First publish

Neither entity exists on Instruqt yet. The preset has to exist before the
track can be validated, because `instruqt track validate` resolves
`sandbox_preset` against the remote. Order matters:

```bash
git push                                  # instruqt/sandbox/proxy/ must be on WORKSHOP_REF
cd instruqt/sandbox
instruqt sandbox push                     # creates the preset as a draft
instruqt sandbox publish --message="Initial Go sandbox"
cd ../track
instruqt track validate
instruqt track push                       # creates the track, writes ids into the YAML
git add . && git commit -m "Track: record Instruqt ids"
instruqt track test
```

`track.yml` and the assignments ship without `id:` and `checksum` fields.
The first push assigns them and writes them back into the files; commit that.

## Pushing the track

Once it exists, a push updates it in place. Run from `instruqt/track/`:

```bash
cd instruqt/track
instruqt track validate    # local files plus the remote preset lookup, no push
instruqt track push
instruqt track pull        # ids for any new challenge or tab, then commit
instruqt track test        # runs every solve script in order, end to end
```

`instruqt track push` rewrites `track.yml` and every `assignment.md`
**frontmatter** in place: comments stripped, keys reordered, block scalars
collapsed, `checksum` injected. Commit that as-is rather than prettifying it,
and keep durable explanation in this README instead of in YAML comments.

> [!IMPORTANT]
> Never edit the track in the web UI while editing locally — it corrupts the
> `.remote` cache. Some settings are web-UI authoritative no matter what you
> push: `idle_timeout`, Hot Start, and invite configuration.

Two things that break a push or silently reorder the lab:

- **Every file in a challenge directory is parsed as a lifecycle script.** A
  stray diagram next to `assignment.md` breaks the track. Shared assets belong
  under `sandbox/` or inline in the markdown.
- **Challenge order comes from directory sort order**, so the `NN-` prefixes
  have to stay sequential and `track.yml` must have no `challenges:` block.
  Lifecycle script names carry the preset's hostname (`-workshop`).

## How the checks tell exercises apart

Every exercise from 2 onward starts the same Workflow ID
(`money-transfer-workflow`), with the same Workflow type (`Transfer`), on the
same task queue. A check that only asked "did a transfer complete?" would pass
exercise 5 on the strength of a run left over from exercise 2.

So each `setup-workshop` writes the current time to
`/tmp/workshop-start-NN`, and each `check-workshop` counts only executions
with `StartTime` after it. Setup also terminates any run still waiting on a
Signal: Go's `ExecuteWorkflow` returns the existing run when the Workflow ID is
already running, so a stuck run from the previous chapter would otherwise
swallow the next starter.

Every check also runs `go build` on the exercise package before it talks to
Temporal, so a compile error is reported as one, not as "no Workflow
completed".

| Challenge | Beyond TODOs gone and the package compiling, the check requires |
|---|---|
| 01 | `greeting-workflow` completed and returned `Hello, Temporal!` |
| 02 | `GetSignalChannel` in the Workflow; a completed transfer whose history has a Signal |
| 03 | A `getStatus` Query handler; a completed transfer |
| 04 | `UpsertSearchAttributes`; an execution with `AccountId="account-123"` |
| 05 | Three `Summary:` fields; every scheduled Activity in the latest run has a summary |
| 06 | A `_test.go` using `go.temporal.io/sdk/testsuite`; `go test ./exercise6` passes with at least two tests |
| 07 | `MaximumAttempts: 1` and a `retry` Signal channel; a completed transfer |

Solve scripts copy `solutionN/` into `exerciseN/` with every `solutionN`
reference rewritten, since the package name and import path both carry it.
The exercise 7 solve script drives the full manual retry path from the CLI:
deposit to `invalid-account`, wait for `RETRYING`, send the retry Signal.

## Pushing the sandbox preset

`instruqt track validate` never runs `sandbox/scripts/`, and a broken setup
script fails at provision time — which, with Hot Start, means a pool of broken
sandboxes. Test it in Docker first; `sandbox/README.md` has the container loop.

```bash
cd instruqt/sandbox
instruqt sandbox diff        # local vs remote, covers config.yml AND scripts/
instruqt sandbox push        # -> draft
instruqt sandbox publish --message="What changed"
git add instruqt/sandbox/ && git commit -m "Sandbox: <what changed>" && git push
```

`push` creates a draft; `publish` makes it live. `push` refuses with `nothing
to push` when local and remote match (`--force` overrides), and `publish`
refuses with `no draft to publish` when there is none.

> [!WARNING]
> `instruqt sandbox publish` moves **every track on this preset** to the newly
> published version. There is no per-track pinning. Do not publish during a
> live session.

## After any change: re-warm Hot Start

A warm pool holds sandboxes built from the preset version and `WORKSHOP_REF`
that were current when it was warmed. They keep serving that state no matter
what you publish or push afterwards.

Re-warm the pool in the web UI after any preset publish, exercise change, or
`WORKSHOP_REF` bump. Then launch one pooled sandbox and click through a
challenge: an upstream blip during warm-up (apt, `go.dev`, `proxy.golang.org`,
`temporal.download`) produces broken pooled sandboxes silently.

Hot Start is a web-UI setting with no CLI equivalent, and it is the supported
mechanism above roughly five simultaneous starts.

## Live workshop pre-flight

- [ ] **`WORKSHOP_REF` pinned** to a tag or commit in
      `sandbox/scripts/setup-workshop`. On the default `main`, pools warmed at
      different times contain different code.
- [ ] **Exercise and `proxy/` changes pushed** to that ref.
- [ ] **Preset published**, if `config.yml` or `scripts/` changed.
- [ ] **`maintenance: false`** in `track.yml`, pushed. It ships `true`, so only
      owners and authors can launch.
- [ ] **Hot Start enabled**, pool sized to attendees plus 10 to 20 percent,
      warmed well ahead of the session and re-warmed after the last change.
- [ ] **`idle_timeout` set in the web UI** (Track, Settings, Additional
      Settings) to workshop length plus 20 minutes. The web UI wins after track
      creation, so a pushed `3600` can still show as `600`.
- [ ] **`feedback_recap_enabled` and `feedback_tab_enabled` set to `false`** in
      `track.yml` if feedback is collected out of band. Both ship `true`, which
      suits a self-paced run.
- [ ] **`instruqt track test` passes.** It catches regressions a manual
      click-through misses.

Afterwards, disable Hot Start if the pool was sized for this session — idle
pre-provisioned sandboxes cost money.

No `secrets:` are declared, so there is nothing to populate under Track
Settings, Secrets.
