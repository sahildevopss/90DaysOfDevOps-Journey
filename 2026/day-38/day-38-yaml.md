# Day 38 – YAML Basics

## Objective

Learn YAML syntax, indentation, lists, nested objects, multi-line strings, and YAML validation.

---

## Task 1: Key-Value Pairs

Create `person.yaml`:

```yaml
name: Sahil Sanadi
role: DevOps Engineer
experience_years: 4
learning: true
```

### Important Points

- `name` is a string.
- `role` is a string.
- `experience_years` is a number.
- `learning: true` is a boolean.
- YAML does not require quotes for simple strings.

---

## Task 2: Lists

Add the following to `person.yaml`:

```yaml
tools:
  - Docker
  - Git
  - AWS
  - Linux
  - Terraform

hobbies: [learning, fitness, technology]
```

### Two Ways to Write Lists in YAML

#### 1. Block Style

```yaml
tools:
  - Docker
  - Git
  - AWS
```

#### 2. Inline Style

```yaml
tools: [Docker, Git, AWS]
```

### Answer

The two common ways to write a list in YAML are:

1. Block style using `-`
2. Inline style using `[item1, item2]`

---

## Complete `person.yaml`

```yaml
name: Sahil Sanadi
role: DevOps Engineer
experience_years: 4
learning: true

tools:
  - Docker
  - Git
  - AWS
  - Linux
  - Terraform

hobbies: [learning, fitness, technology]
```

---

## Task 3: Nested Objects

Create `server.yaml`:

```yaml
server:
  name: devops-server
  ip: 192.168.1.100
  port: 8080

database:
  host: localhost
  name: devopsdb
  credentials:
    user: devops
    password: secret123
```

### Structure

```text
server
├── name
├── ip
└── port

database
├── host
├── name
└── credentials
    ├── user
    └── password
```

### Tab vs Spaces

Try replacing spaces with a tab:

```yaml
server:
	name: devops-server
	ip: 192.168.1.100
```

Validate the file:

```bash
yamllint server.yaml
```

### What Happens?

YAML does not allow tabs for indentation. The validator will report an error because a tab character was used instead of spaces.

A typical error may look similar to:

```text
syntax error: found character '\t' that cannot start any token
```

### Answer

When I used a tab instead of spaces for indentation, YAML validation failed because YAML does not allow tabs for indentation. YAML uses spaces to define structure and nesting.

---

## Task 4: Multi-line Strings

Add a `startup_script` field to `server.yaml`.

### 1. Block Style `|`

The `|` style preserves newlines.

```yaml
startup_script: |
  #!/bin/bash
  echo "Starting server"
  systemctl start nginx
  echo "Server started"
```

Use `|` when line breaks need to be preserved.

Common use cases:

- Shell scripts
- Configuration files
- Multi-line commands

### 2. Fold Style `>`

The `>` style folds multiple lines into a single line.

```yaml
startup_message: >
  This is a startup message
  that will be folded
  into a single line.
```

The resulting value is conceptually:

```text
This is a startup message that will be folded into a single line.
```

Use `>` when multiple lines should be treated as one continuous line.

Common use cases:

- Long sentences
- Paragraphs
- Long descriptions

### Answer: `|` vs `>`

Use `|` when line breaks need to be preserved. Use `>` when multiple lines should be folded into a single line.

---

## Task 5: Validate Your YAML

### Install `yamllint`

On Ubuntu/WSL:

```bash
sudo apt update
sudo apt install yamllint -y
```

Check the installation:

```bash
yamllint --version
```

### Validate `person.yaml`

```bash
yamllint person.yaml
```

### Validate `server.yaml`

```bash
yamllint server.yaml
```

### Validate All YAML Files

```bash
yamllint *.yaml
```

### Intentionally Break the Indentation

Incorrect YAML:

```yaml
server:
  name: devops-server
    ip: 192.168.1.100
```

The indentation of `ip` is incorrect.

Correct version:

```yaml
server:
  name: devops-server
  ip: 192.168.1.100
```

Validate again:

```bash
yamllint server.yaml
```

After fixing the indentation, the YAML should validate successfully, assuming there are no other errors.

---

## Task 6: Spot the Difference

### Block 1 – Correct

```yaml
name: devops
tools:
  - docker
  - kubernetes
```

### Block 2 – Broken

```yaml
name: devops
tools:
- docker
  - kubernetes
```

### What Is Wrong?

The list items are not consistently indented.

The recommended version is:

```yaml
name: devops
tools:
  - docker
  - kubernetes
```

YAML uses indentation to understand the structure of the data.

---

## Final `server.yaml`

```yaml
server:
  name: devops-server
  ip: 192.168.1.100
  port: 8080

database:
  host: localhost
  name: devopsdb
  credentials:
    user: devops
    password: secret123

startup_script: |
  #!/bin/bash
  echo "Starting server"
  systemctl start nginx
  echo "Server started"

startup_message: >
  This is a startup message
  that will be folded
  into a single line.
```

---

## What I Learned

### 1. Indentation Matters

YAML uses indentation to define relationships between keys and values. Spaces should be used instead of tabs.

### 2. YAML Supports Different List Styles

Lists can be written using block style:

```yaml
tools:
  - Docker
  - Git
  - AWS
```

Or inline style:

```yaml
tools: [Docker, Git, AWS]
```

### 3. Multi-line Strings Have Different Behaviors

The `|` operator preserves line breaks, while `>` folds line breaks into spaces.

---

## Important YAML Rules

- Use spaces, never tabs, for indentation.
- Two spaces is the common indentation style.
- Indentation defines the structure of YAML.
- `true` and `false` are boolean values.
- `"true"` is a string.
- Lists can use `-` or inline `[ ]` syntax.
- Use `|` when newlines need to be preserved.
- Use `>` when lines should be folded together.
- Validate YAML before using it in CI/CD pipelines.

---

## Verification Commands

```bash
cat person.yaml
cat server.yaml

yamllint person.yaml
yamllint server.yaml
```

---

## Expected Directory Structure

```text
2026/
└── day-38/
    ├── person.yaml
    ├── server.yaml
    └── day-38-yaml.md
---

---

## Key Takeaway

YAML looks simple, but indentation controls its structure. A small indentation mistake or an accidental tab can cause a YAML file to become invalid.

This becomes especially important when working with CI/CD tools such as GitHub Actions, where YAML is used to define automation workflows.
