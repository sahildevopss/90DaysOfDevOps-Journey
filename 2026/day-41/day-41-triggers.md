
# Day 41 – Triggers & Matrix Builds

## Task 1: Trigger on Pull Request

### Objective

Create a GitHub Actions workflow that runs when a pull request is opened or updated against the `main` branch.

---

## Workflow File

File:

`.github/workflows/pr-check.yml`

```yaml
name: Trigger on pull request

on:
  pull_request:
    branches:
      - main
    types:
      - opened
      - synchronize

jobs:
  pr-check:
    runs-on: ubuntu-latest

    steps:
      - name: Show PR branch
        run: echo "PR check running for branch is ${{ github.head_ref }}"'
```

---

## Implementation

### 1. Created the Pull Request Workflow

Created the following workflow file:

`.github/workflows/pr-check.yml`

The workflow was configured to trigger only when a pull request is opened or updated against the `main` branch.

### 2. Created a Feature Branch

Created a new branch:

```text
feature
```

### 3. Created a Pull Request

Created a pull request from:

```text
feature → main
```

### 4. Tested the Pull Request Trigger

The workflow was configured with two pull request events:

- `opened` – triggers when a new pull request is created.
- `synchronize` – triggers when new commits are pushed to the pull request branch.

The workflow prints the name of the source branch using:

```yaml
${{ github.head_ref }}
```

Expected output:

```text
PR check running for branch: feature
```

---

## Issue Encountered

Initially, the workflow failed with a YAML syntax error on the `run` line.

The original line was:

```yaml
run: echo "PR check running for branch: ${{ github.head_ref }}"
```

The problem was caused by the colon followed by a space in:

```text
branch: 
```

YAML interpreted it as mapping syntax.

### Fix

The complete command was enclosed in single quotes or we can remove the colon:

```yaml
run: 'echo "PR check running for branch: ${{ github.head_ref }}"'
```

After fixing the syntax and pushing the changes, the workflow executed successfully.

---

## Verification

The pull request workflow successfully ran against the pull request:

```text
feature → main
```

The workflow displayed:

```text
PR check running for branch: feature
```

This confirmed that the `pull_request` trigger was working correctly.

---

## Key Learning

A `push` trigger and a `pull_request` trigger respond to different GitHub events.

The existing `hello.yml` workflow uses:

```yaml
on: push
```

so it runs whenever a commit is pushed.

The new `pr-check.yml` workflow uses:

```yaml
on:
  pull_request:
```

so it runs when the specified pull request events occur.

I also learned that YAML values containing a colon followed by a space may need to be quoted.

---

## Screenshot

Add the GitHub Pull Request / Actions screenshot here:

![Pull Request Workflow](screenshots/PR-workflow.png)


