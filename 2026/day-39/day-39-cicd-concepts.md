# Day 39 – CI/CD Concepts

## Objective

Understand why CI/CD exists, how CI differs from Continuous Delivery and Continuous Deployment, and how a CI/CD pipeline works before writing any pipeline configuration.

---

# Task 1 – The Problem

## What Can Go Wrong?

Imagine a team of 5 developers working on the same repository and manually deploying applications to production.

Problems that can occur:

- Developers can introduce conflicting changes.
- Bugs may be discovered only after deployment.
- Manual deployment steps can be missed or performed incorrectly.
- Different developers may use different versions of dependencies or tools.
- The application may work on one developer's machine but fail in another environment.
- Testing may be inconsistent or skipped.
- Deployments become slower and more stressful as the team grows.
- It becomes difficult to know exactly what was deployed and when.
- Repeating the same deployment steps manually increases the chance of human error.
- Rollbacks can become difficult if the deployment process is not standardized.

## What Does "It Works on My Machine" Mean?

"It works on my machine" means that an application works correctly in one developer's local environment but fails in another environment.

This can happen because of differences in:

- Operating system
- Programming language version
- Dependency versions
- Environment variables
- Configuration files
- Installed software
- Database versions
- Network configuration

CI/CD helps reduce these problems by creating a consistent and automated process for building, testing, and deploying applications.

## How Many Times a Day Can a Team Safely Deploy Manually?

There is no fixed number of manual deployments that is considered safe.

The real problem is that manual deployment does not scale well.

A team might manually deploy a few times per day, but as deployment frequency increases, repetitive manual steps create more opportunities for:

- Human error
- Missed steps
- Configuration mistakes
- Inconsistent deployments
- Longer deployment times

CI/CD allows teams to make deployments more frequent, repeatable, and consistent.

---

# Task 2 – CI vs CD

## Continuous Integration (CI)

Continuous Integration means developers frequently integrate their code into a shared repository.

Whenever code is pushed or a pull request is created, automated processes can build the application and run tests.

CI helps detect integration problems, build failures, and bugs early.

### Real-World Example

```text
Developer
    ↓
git push
    ↓
GitHub
    ↓
Build
    ↓
Automated Tests
    ↓
Pass / Fail
```

If the tests fail, the team can fix the problem before the code moves further through the delivery process.

---

## Continuous Delivery

Continuous Delivery extends CI by keeping the application in a deployable state.

After the code passes automated validation, the application can be packaged and prepared for deployment.

Production deployment may still require manual approval.

### Real-World Example

```text
Developer
    ↓
Push to GitHub
    ↓
Build
    ↓
Test
    ↓
Build Docker Image
    ↓
Deploy to Staging
    ↓
Manual Approval
    ↓
Production
```

The important idea is that the software is always ready to be released.

---

## Continuous Deployment

Continuous Deployment takes automation one step further.

Every change that successfully passes the required pipeline checks is automatically deployed to production without requiring a manual approval.

### Real-World Example

```text
Developer
    ↓
Push to GitHub
    ↓
Build
    ↓
Test
    ↓
Build Docker Image
    ↓
Deploy to Production
```

If the pipeline passes all required checks, the deployment happens automatically.

---

# CI vs Continuous Delivery vs Continuous Deployment

| Concept | Main Purpose | Production Deployment |
|---|---|---|
| Continuous Integration | Integrate and validate code frequently | Not necessarily |
| Continuous Delivery | Keep software ready to release | Usually requires approval |
| Continuous Deployment | Automatically release validated changes | Automatic |

### Simple Way to Remember

```text
CI
↓
Integrate + Test

Continuous Delivery
↓
Integrate + Test + Keep Ready to Deploy

Continuous Deployment
↓
Integrate + Test + Automatically Deploy
```

---

# Task 3 – Pipeline Anatomy

## Trigger

A trigger is the event that starts a pipeline.

Common triggers include:

- Push to a branch
- Pull request
- Git tag
- Manual execution
- Scheduled execution

Example:

```text
Developer pushes code
        ↓
Pipeline Triggered
```

---

## Stage

A stage is a logical phase of a pipeline.

Common stages include:

- Build
- Test
- Package
- Deploy

Example:

```text
Build Stage
     ↓
Test Stage
     ↓
Deploy Stage
```

---

## Job

A job is a unit of work executed as part of a pipeline.

A stage can contain one or more jobs depending on the CI/CD platform.

Example:

```text
Test Stage
    ↓
Test Job
```

A test job could install dependencies and run automated tests.

---

## Step

A step is a single command or action inside a job.

Example:

```text
Test Job
    ↓
Step 1 → Install dependencies
Step 2 → Run tests
Step 3 → Generate test report
```

---

## Runner

A runner is the machine or execution environment that runs the jobs.

The runner executes the commands defined in the pipeline.

Example:

```text
Pipeline Job
     ↓
GitHub Actions Runner
     ↓
Commands Execute
```

A runner can be hosted by the CI/CD platform or managed by the organization itself.

---

## Artifact

An artifact is an output produced by a job.

Examples include:

- Build files
- Test reports
- Compiled binaries
- Application packages
- Logs
- Docker image metadata

Example:

```text
Build Job
    ↓
Application Build
    ↓
Build Artifact
```

---

# Task 4 – CI/CD Pipeline Diagram

## Scenario

A developer pushes code to GitHub.

The application is tested, built into a Docker image, and deployed to a staging server.

```text
                    Developer
                        |
                        | git push
                        v
                +---------------+
                |    GitHub     |
                +---------------+
                        |
                     Trigger
                        |
                        v
              +-------------------+
              |   Stage 1: Test   |
              +-------------------+
                        |
                  Test Job
                        |
              +-------------------+
              | Checkout Code     |
              | Install Depends   |
              | Run Tests         |
              +-------------------+
                        |
                      PASS
                        |
                        v
              +-------------------+
              |  Stage 2: Build   |
              +-------------------+
                        |
                   Build Job
                        |
              +-------------------+
              | Build Application |
              | Build Docker Image|
              +-------------------+
                        |
                   Docker Image
                        |
                        v
              +-------------------+
              |  Stage 3: Deploy  |
              +-------------------+
                        |
                  Deploy Job
                        |
                        v
              +-------------------+
              |  Staging Server    |
              +-------------------+
```

## Pipeline Flow

```text
Code Push
    ↓
Trigger
    ↓
Test
    ↓
Build
    ↓
Docker Image
    ↓
Deploy
    ↓
Staging Server
```

---

# Task 5 – Explore in the Wild

## Repository

Repository explored:

FastAPI

GitHub Repository:

https://github.com/fastapi/fastapi

Workflow directory:

```text
.github/workflows/
```

I selected one of the workflow YAML files from the repository and inspected its configuration.

## What Triggers the Workflow?

The trigger is defined using the `on:` section of the GitHub Actions workflow.

Typical GitHub Actions triggers include:

```yaml
on:
  push:
  pull_request:
```

This means the workflow can be triggered when code is pushed or when a pull request is created or updated.

## How Many Jobs Does It Have?

The number of jobs can be identified under:

```yaml
jobs:
```

Each top-level entry under `jobs` represents a separate job.

For example:

```yaml
jobs:
  test:
    ...

  build:
    ...
```

In this example, there are two jobs:

- `test`
- `build`

## What Does the Workflow Do?

The workflow automates tasks that would otherwise need to be performed manually.

The general process is:

```text
GitHub Event
     ↓
Runner Starts
     ↓
Checkout Repository
     ↓
Install Dependencies
     ↓
Run Checks / Tests
     ↓
Report Result
```

---

# Important CI/CD Concepts

## CI/CD Is a Practice, Not a Tool

CI/CD is not a single software product.

It is a set of practices used to automate and improve the software development and delivery process.

Tools such as:

- GitHub Actions
- Jenkins
- GitLab CI/CD
- CircleCI

can be used to implement CI/CD practices.

---

# Why CI/CD Matters

Without automation:

```text
Code
 ↓
Manual Testing
 ↓
Manual Build
 ↓
Manual Deployment
 ↓
Production
```

With CI/CD:

```text
Code
 ↓
Automated Build
 ↓
Automated Tests
 ↓
Automated Packaging
 ↓
Automated / Controlled Deployment
 ↓
Production or Staging
```

CI/CD provides:

- Faster feedback
- Repeatable processes
- Consistent builds
- Automated testing
- Reduced manual errors
- Faster deployments
- Better visibility into failures
- More reliable software delivery

---

# Pipeline Failure Is Not Always a Bad Thing

A failed pipeline does not necessarily mean CI/CD is broken.

A pipeline failure can mean that automation successfully detected a problem.

For example:

```text
Developer Push
      ↓
Pipeline
      ↓
Tests
      ↓
FAIL
      ↓
Developer Fixes Bug
      ↓
Push Again
      ↓
Pipeline
      ↓
PASS
```

The purpose of CI/CD is not to make every pipeline pass.

The purpose is to provide fast and reliable feedback so problems are detected before they reach users or production.

---

# Key Takeaways

- CI/CD is a software delivery practice, not just a tool.
- Continuous Integration focuses on integrating and validating code frequently.
- Continuous Delivery keeps software in a deployable state.
- Continuous Deployment automatically releases validated changes.
- A trigger starts a pipeline.
- A stage represents a logical phase of the pipeline.
- A job is a unit of work.
- A step is an individual action or command.
- A runner executes the job.
- An artifact is an output produced by the pipeline.
- Automation makes software delivery more repeatable and consistent.
- A failed pipeline can be useful because it provides early feedback.
- GitHub Actions, Jenkins, GitLab CI/CD, and CircleCI are tools used to implement CI/CD.

---

# Day 39 Summary

The main lesson from today is that CI/CD is about creating a reliable path from code change to validated software and, when appropriate, deployment.

The basic flow is:

```text
Developer
    ↓
Git Push
    ↓
Trigger
    ↓
Build
    ↓
Test
    ↓
Package
    ↓
Deploy
    ↓
Staging / Production
```

Before creating a CI/CD pipeline, understanding this flow is important because the YAML configuration is simply a way of expressing these steps to an automation system.
