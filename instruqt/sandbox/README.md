# Sandbox Preset

This directory is an Instruqt **sandbox preset**, slug
`temporal-in-practice-sandbox-go`. It is a first-class Instruqt entity with
its own versioned release cycle, pushed and published independently of the
track.

```
sandbox/
├── config.yml                    Preset definition: name, slug, containers   [PUSHED]
├── scripts/
│   └── setup-workshop            All provisioning, git included              [PUSHED]
├── proxy/                        Network control panel, fetched via git clone
│   ├── toggle_addon.py           mitmproxy addon, re-reads state.json per request
│   ├── controlpanel.py           Flask UI + REST API on port 5000
│   ├── state.json                Default toggle state
│   └── static/index.html         The toggle UI
└── README.md                     This file
```

Only `config.yml` and `scripts/` are part of the preset payload. `proxy/` and
this README sit in the same directory for cohesion but are not uploaded; a
`sandbox pull` returns just those two. `proxy/` reaches the sandbox through the
git clone that `setup-workshop` performs.

## Why a preset

A preset shares one sandbox definition across tracks, prevents config drift,
and lets tracks **share a hot-start pool**. Because `scripts/setup-workshop`
ships with the preset rather than inside a git clone, it can install its own
dependencies, `git` among them. The base image is relied on for nothing but a
Debian userland, and `git clone` fetches only the exercise code.

The track therefore has **no `config.yml` and no `track_scripts/`**. It points
here from `track.yml`:

```yaml
sandbox_preset: temporal-in-practice-sandbox-go
```

## How this ships

| Change | Command | Takes effect |
|--------|---------|--------------|
| `config.yml`, `scripts/` | `instruqt sandbox push` then `instruqt sandbox publish` | Next sandbox **provision**, for every track on the preset |
| `proxy/` | `git push` to `WORKSHOP_REF` | Next sandbox provision |
| `../track/` (assignments, `track.yml`) | `instruqt track push` | Next lab start |

> [!WARNING]
> `instruqt sandbox publish` moves **every track using this preset** to the
> newly published version. There is no per-track version pinning. Do not
> publish during a live session.

## Updating the preset

### 1. Edit and test locally

`instruqt track validate` does not run this script, and a broken script
fails at provision time. With Hot Start that means a pool of broken sandboxes,
so test before pushing.

Boot the base image and run the script against your working tree. Seeding
`/root/workshop` first makes the script skip its clone, so it uses your
uncommitted code. Instruqt runs `linux/amd64`; drop `--platform` for a faster
native run on Apple silicon (the script handles both architectures):

```bash
docker rm -f instruqt-test 2>/dev/null
docker run -d --name instruqt-test --platform linux/amd64 \
  -p 8299:8233 -p 5099:5000 \
  ubuntu:24.04 sleep infinity

docker exec instruqt-test mkdir -p /opt/instruqt/bootstrap /root/workshop
docker exec instruqt-test touch /opt/instruqt/bootstrap/host-bootstrap-completed

tar --exclude='./.git' -cf - . \
  | docker exec -i instruqt-test tar -xf - -C /root/workshop

docker cp instruqt/sandbox/scripts/setup-workshop instruqt-test:/tmp/setup-workshop
docker exec instruqt-test /tmp/setup-workshop
```

Run those from the repository root. Temporal UI lands on
<http://localhost:8299>, the control panel on <http://localhost:5099>.

Then walk every challenge end to end. `env -u HOME` mimics a lifecycle script
that runs without a home directory, which is the case the scripts guard
against:

```bash
docker cp instruqt/track instruqt-test:/tmp/ch
docker exec instruqt-test env -u HOME bash -c '
  for d in /tmp/ch/0*; do
    echo "## $(basename $d)"
    $d/setup-workshop
    $d/check-workshop && echo "!! check passed before solve"
    $d/solve-workshop >/dev/null 2>&1
    $d/check-workshop || echo "!! check failed after solve"
    $d/cleanup-workshop
  done'

docker rm -f instruqt-test
```

To exercise the real clone path instead, skip the seeding step. `WORKSHOP_REF`
then has to name a ref you have already pushed:

```bash
docker exec -e WORKSHOP_REF=my-branch instruqt-test /tmp/setup-workshop
```

### 2. Push and publish

Run from this directory. `push` creates a remote **draft** (and the preset
itself, the first time); `publish` makes it live.

```bash
cd instruqt/sandbox
instruqt sandbox diff        # local vs remote, covers config.yml AND scripts/
instruqt sandbox push        # -> draft
instruqt sandbox publish --message="What changed"
```

`instruqt sandbox pull <slug>` fetches the latest version, preferring a draft
over the published one. Instruqt's serializer strips comments from
`config.yml`, so a pull looks different from what is committed here; `diff` is
semantic and reports no change.

Then commit, because `proxy/` and the script are both source of truth in git.

### 3. Re-warm the Hot Start pool

**After any preset publish, exercise change, or `WORKSHOP_REF` bump, re-warm
the pool in the Instruqt web UI.** Then launch one sandbox from the pool and
click through a challenge.

## What setup-workshop provisions

Nine steps, in order. This script owns every binary dependency; nothing is
inherited from the base image.

1. System packages: `git`, `curl`, `jq`, `unzip`, `ca-certificates`, editors.
2. `git clone --depth 1` of the exercise repo to `/root/workshop`.
3. Go from the official tarball to `/usr/local/go`, symlinked into
   `/usr/local/bin`, with `GOTOOLCHAIN=local` written into `$GOROOT/go.env`.
4. The Temporal CLI, symlinked to `/usr/local/bin/temporal`.
5. `go mod download`, then `go build ./solution...` and `go test ./solution6`
   to compile the SDK and testify into the build cache.
6. uv, a Python for the proxy only, mitmproxy and Flask in `/opt/proxy-venv`,
   and the mitmproxy CA cert trusted system-wide.
7. `proxy/` staged to `/root/proxy/` from the clone.
8. Shell environment written to `/etc/profile.d/workshop.sh`.
9. mitmproxy, the control panel, and `temporal server start-dev` started, then
   a health wait that fails the provision if the server never comes up.

Measured at roughly a minute on `ubuntu:24.04`. Hot Start absorbs it.

### Adding a dependency

Edit step 1. Not the base image, and not the track.

### Changing the Go version

`GO_VERSION` at the top of the script. Keep it at or above the `toolchain`
line in `go.mod`. With `GOTOOLCHAIN=local`, an older install does not quietly
download a newer toolchain through the proxy; it fails with a clear version
error instead.

### Changing the base image

`config.yml` in this directory. It needs only a Debian or Ubuntu userland with
`apt-get` and `dpkg`; the script installs everything else.

## Things that are easy to break

**Toolchain on the default PATH.** `go`, `gofmt` and `temporal` are symlinked
into `/usr/local/bin` on purpose. Check and solve scripts run
non-interactively and never source a profile, so a `PATH` export in a profile
is invisible to them. They also set `HOME` if it is missing, because Go needs
it to find `GOPATH` and `GOCACHE`.

**A warm build cache.** Step 5 is what makes a learner's first `go run` take
a second instead of a minute. It lives in `/root/.cache/go-build` and
`/root/go/pkg/mod`. Running anything as a user other than root throws it away.

**Proxy and cert variables.** They live in `/etc/profile.d/workshop.sh`, with
`/etc/bash.bashrc` sourcing it for interactive non-login shells, which is what
Instruqt's terminal tabs open. Putting them only in `/etc/bash.bashrc` does not
work: Ubuntu's `/root/.bashrc` returns early for non-interactive shells.

**`NO_PROXY`.** grpc-go honours `HTTPS_PROXY`. Without
`NO_PROXY=127.0.0.1,localhost`, the SDK's connection to `localhost:7233` is
sent to mitmproxy and every Worker and starter fails to connect.

## Network control panel

Ported from [temporal-sa/ai-agents-webinar](https://github.com/temporal-sa/ai-agents-webinar).
mitmproxy listens on `127.0.0.1:8888`, every shell routes outbound HTTP and
HTTPS through it, and the Flask panel on port 5000 flips toggles by writing
`state.json`. The addon re-reads that file on every request, so a toggle takes
effect immediately with no restart.

| Toggle | Hosts | Effect when off |
|--------|-------|-----------------|
| `goproxy` | `proxy.golang.org`, `sum.golang.org` | `go get` and module downloads fail |
| `github` | `github.com`, `raw.githubusercontent.com`, `codeload.github.com` | Git and release downloads fail |
| `httpbin` | `httpbin.org` | A request-echo endpoint for proving the proxy is in the path |
| Kill switch | everything | HTTP 503 on all outbound traffic |

The seven exercises call no external HTTP APIs of their own, and every module
they need is already in the module cache, so the toggles do not disrupt any
Workflow as shipped. The panel works as an instructor demo now: flip `goproxy`
off, run `go get github.com/google/uuid@v1.5.0` in a scratch module, watch it
fail with the proxy's 503, flip it back.

To add a toggle: add the key and its hostnames to `SERVICE_HOSTS` in
`toggle_addon.py`, add the key to `state.json`, and add a display label to
`SERVICE_LABELS` in `controlpanel.py`. The UI renders itself from
`SERVICE_LABELS`, so there is no HTML to edit. These files ship on a
`git push`, not a `sandbox push`.

## Logs

| File | What |
|------|------|
| `/tmp/proxy.log` | mitmproxy |
| `/tmp/controlpanel.log` | Flask control panel |
| `/tmp/temporal-server.log` | Temporal dev server |
| `/tmp/mitm-init.log` | The throwaway run that generates the CA cert |
| `/tmp/solve-NN-worker.log` | The Worker a solve script started |

The provisioning output itself goes to the sandbox's lifecycle-script log,
visible in the Instruqt web UI.
