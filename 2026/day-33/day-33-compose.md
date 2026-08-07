# Day 33 - Docker Compose: Multi-Container Basics

## Objective

Learn how to manage single and multi-container applications using Docker Compose. Instead of creating containers, networks, and volumes manually, Docker Compose allows us to define everything in a single YAML file and manage the entire application with one command.

---

# Task 1 - Install & Verify Docker Compose

## Check Docker Compose

```bash
docker compose version
```

### Output

```bash
Docker Compose version v2.x.x
```

---

# Task 2 - First Docker Compose Project

## Project Structure

```
compose-basics/
│── docker-compose.yml
```

## docker-compose.yml

```yaml
services:
  nginx:
    image: nginx:latest
    container_name: nginx-compose
    ports:
      - "8080:80"
```

## Start Container

```bash
docker compose up
```

Or run in detached mode

```bash
docker compose up -d
```

## Verify Running Containers

```bash
docker ps
```

Open browser

```
http://localhost:8080
```

You should see the default Nginx page.

## Stop Everything

```bash
docker compose down
```

