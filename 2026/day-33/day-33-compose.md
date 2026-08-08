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
---

# Task 3 - WordPress + MySQL Multi-Container Setup

## Project Structure

```
wordpress-compose/
│── docker-compose.yml
```

## docker-compose.yml

```yaml
services:
  db:
    image: mysql:8.0
    container_name: mysql-db
    restart: always

    environment:
      MYSQL_ROOT_PASSWORD: root123
      MYSQL_DATABASE: wordpress
      MYSQL_USER: wpuser
      MYSQL_PASSWORD: password123

    volumes:
      - mysql_data:/var/lib/mysql

  wordpress:
    image: wordpress:latest
    container_name: wordpress

    restart: always

    depends_on:
      - db

    ports:
      - "8081:80"

    environment:
      WORDPRESS_DB_HOST: db:3306
      WORDPRESS_DB_USER: wpuser
      WORDPRESS_DB_PASSWORD: password123
      WORDPRESS_DB_NAME: wordpress

volumes:
  mysql_data:
```

## Start Project

```bash
docker compose up -d
```

Verify

```bash
docker ps
```

Open browser

```
http://localhost:8081
```

Complete the WordPress installation.

---

## Verify Data Persistence

Stop everything

```bash
docker compose down
```

Start again

```bash
docker compose up -d
```

Visit

```
http://localhost:8081
```

The WordPress installation and database remain available because MySQL uses the named volume:

```
mysql_data
```

---

# Task 4 - Common Docker Compose Commands

## Start in Detached Mode

```bash
docker compose up -d
```

---

## View Running Services

```bash
docker compose ps
```

---

## View Logs of All Services

```bash
docker compose logs
```

Follow logs

```bash
docker compose logs -f
```

---

## View Logs of a Specific Service

```bash
docker compose logs wordpress
```

or

```bash
docker compose logs db
```

---

## Stop Services Without Removing

```bash
docker compose stop
```

Restart

```bash
docker compose start
```

---

## Remove Containers and Network

```bash
docker compose down
```

Remove everything including volumes

```bash
docker compose down -v
```

---

## Rebuild Images

```bash
docker compose up --build
```

or

```bash
docker compose up --build -d
```

---

# Task 5 - Environment Variables

## docker-compose.yml

```yaml
services:
  db:
    image: mysql:8.0

    environment:
      MYSQL_ROOT_PASSWORD: ${MYSQL_ROOT_PASSWORD}
      MYSQL_DATABASE: ${MYSQL_DATABASE}
      MYSQL_USER: ${MYSQL_USER}
      MYSQL_PASSWORD: ${MYSQL_PASSWORD}
```

## .env

```
MYSQL_ROOT_PASSWORD=root123
MYSQL_DATABASE=wordpress
MYSQL_USER=wpuser
MYSQL_PASSWORD=password123
```

Docker Compose automatically reads the .env file.

Verify

```bash
docker compose config
```

This command displays the final configuration with variables substituted.

---

# Important Docker Compose Commands

| Command | Purpose |
|----------|----------|
| docker compose up | Start services |
| docker compose up -d | Start in background |
| docker compose down | Remove containers and network |
| docker compose stop | Stop containers |
| docker compose start | Restart stopped containers |
| docker compose ps | View running services |
| docker compose logs | Show logs |
| docker compose logs -f | Follow logs |
| docker compose logs SERVICE | Logs of one service |
| docker compose config | Validate compose file |
| docker compose up --build | Rebuild and start |
| docker compose down -v | Remove containers, network, and volumes |

---

# Key Learnings

- Docker Compose manages multi-container applications using a single YAML file.
- Compose automatically creates a dedicated network.
- Service names act as hostnames for inter-container communication.
- Named volumes preserve data even after containers are removed.
- Environment variables improve security and configuration management.
- A single command can start or stop the complete application stack.

---

# Conclusion

Today I learned how Docker Compose simplifies container orchestration by defining infrastructure as code. Managing applications with Compose is faster, cleaner, and more scalable than running individual Docker commands manually.
