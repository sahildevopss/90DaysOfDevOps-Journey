# Day 37 – Docker Revision & Cheat Sheet

## Goal

Take a one-day pause to consolidate everything from Days 29–36 so Docker concepts and commands become easier to recall and apply.

---

# Self-Assessment Checklist

Mark yourself honestly:

- **Can do** — I can perform and explain it without looking at notes.
- **Shaky** — I understand it but need more practice or reference.
- **Haven't done** — I have not practiced it yet.

## 1. Run a container from Docker Hub — Interactive + Detached

**Required answer:**

Interactive:

```bash
docker run -it ubuntu bash
```

Runs an Ubuntu container interactively and opens a Bash shell.

Detached:

```bash
docker run -d nginx
```

Runs an Nginx container in the background.

**Self-assessment:** Can do

---

## 2. List, Stop, Remove Containers and Images

**Required answer:**

```bash
docker ps
docker ps -a
docker stop <container>
docker rm <container>

docker images
docker rmi <image>
```

- `docker ps` — lists running containers.
- `docker ps -a` — lists all containers.
- `docker stop` — stops a running container.
- `docker rm` — removes a container.
- `docker images` — lists local images.
- `docker rmi` — removes an image.

**Self-assessment:** Can do

---

## 3. Explain Image Layers and How Caching Works

**Required answer:**

A Docker image is made up of multiple read-only layers.

Docker can reuse unchanged layers from previous builds. This is called **build cache**.

Example:

```dockerfile
FROM node:22

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

CMD ["npm", "start"]
```

If only application source code changes, Docker can reuse the `COPY package*.json` and `RUN npm install` layers as long as the dependency files have not changed.

View image layers with:

```bash
docker history <image>
```

**Self-assessment:** Can do

---

## 4. Write a Dockerfile from Scratch with FROM, RUN, COPY, WORKDIR, CMD

**Required answer:**

```dockerfile
FROM nginx:alpine

WORKDIR /usr/share/nginx/html

COPY index.html .

RUN echo "Docker image built successfully"

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
```

- `FROM` — specifies the base image.
- `RUN` — executes a command while building the image.
- `COPY` — copies files from the build context into the image.
- `WORKDIR` — sets the working directory.
- `EXPOSE` — documents the container port.
- `CMD` — specifies the default command.

**Self-assessment:** Can do

---

## 5. Explain CMD vs ENTRYPOINT

**Required answer:**

`CMD` provides a default command or default arguments and can be overridden at runtime.

`ENTRYPOINT` defines the main executable of the container.

Example:

```dockerfile
FROM alpine

ENTRYPOINT ["echo"]

CMD ["Hello"]
```

Run normally:

```bash
docker run --rm <image>
```

Output:

```text
Hello
```

Run with a different argument:

```bash
docker run --rm <image> "Docker World"
```

Output:

```text
Docker World
```

The `ENTRYPOINT` remains `echo`, while the runtime argument replaces the default `CMD`.

**Self-assessment:** Can do

---

## 6. Build and Tag a Custom Image

**Required answer:**

Create a Dockerfile:

```dockerfile
FROM nginx:alpine

COPY index.html /usr/share/nginx/html/

EXPOSE 80
```

Build the image:

```bash
docker build -t my-nginx:v1 .
```

Here:

- `docker build` — builds an image.
- `-t` — assigns a name and tag.
- `my-nginx` — image name.
- `v1` — image tag/version.
- `.` — current directory is the build context.

Create another tag:

```bash
docker tag my-nginx:v1 my-nginx:latest
```

Check the images:

```bash
docker images
```

**Self-assessment:** Can do

---

## 7. Create and Use Named Volumes

**Required answer:**

Create a named volume:

```bash
docker volume create mydata
```

List volumes:

```bash
docker volume ls
```

Use the volume:

```bash
docker run -d --name nginx-data -v mydata:/data nginx
```

Inspect it:

```bash
docker volume inspect mydata
```

Named volumes persist independently of the container.

**Self-assessment:** Can do

---

## 8. Use Bind Mounts

**Required answer:**

```bash
docker run -v $(pwd):/app nginx
```

This mounts the current directory on the host to `/app` inside the container.

A bind mount directly maps a host filesystem path to a container path.

**Self-assessment:** Can do

---

## 9. Create Custom Networks and Connect Containers

**Required answer:**

Create a custom network:

```bash
docker network create app-network
```

Run containers on the network:

```bash
docker run -d --name app1 --network app-network nginx
docker run -d --name app2 --network app-network nginx
```

Containers on the same custom network can communicate using their container names.

For example:

```text
app1 → app2
```

can communicate using:

```text
http://app2
```

instead of using the container IP address.

**Self-assessment:** Can do

---

## 10. Write a docker-compose.yml for a Multi-Container App

**Required answer:**

```yaml
services:
  frontend:
    image: nginx:alpine
    ports:
      - "8080:80"

  backend:
    image: my-backend:latest
```

Start the application:

```bash
docker compose up -d
```

Check services:

```bash
docker compose ps
```

View logs:

```bash
docker compose logs
```

Stop and remove services:

```bash
docker compose down
```

**Self-assessment:** Can do

---

## 11. Use Environment Variables and .env Files in Compose

**Required answer:**

`.env`:

```text
APP_PORT=8080
```

`docker-compose.yml`:

```yaml
services:
  frontend:
    image: nginx:alpine
    ports:
      - "${APP_PORT}:80"
```

Compose substitutes `${APP_PORT}` with the value from `.env`.

Verify the resolved configuration:

```bash
docker compose config
```

Do not commit `.env` if it contains secrets.

**Self-assessment:** Shaky

---

## 12. Write a Multi-Stage Dockerfile

**Required answer:**

```dockerfile
FROM node:22 AS build

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .
RUN npm run build

FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]
```

The first stage contains the build environment.

The second stage contains only the files required to run the application.

Multi-stage builds help create smaller and cleaner production images.

**Self-assessment:** Can do

---

## 13. Push an Image to Docker Hub

**Required answer:**

Tag the image:

```bash
docker tag my-nginx:v1 <username>/my-nginx:v1
```

Push it:

```bash
docker push <username>/my-nginx:v1
```

Pull it:

```bash
docker pull <username>/my-nginx:v1
```

**Self-assessment:** Can do

---

## 14. Use Healthchecks and depends_on

**Required answer:**

A healthcheck allows Docker to determine whether a container is healthy.

Example:

```yaml
healthcheck:
  test: ["CMD", "curl", "-f", "http://localhost:8080/health"]
  interval: 30s
  timeout: 10s
  retries: 3
```

Check the container health status with:

```bash
docker ps
```

or:

```bash
docker inspect <container>
```

`depends_on` defines service dependencies in Compose.

Example:

```yaml
services:
  backend:
    image: my-backend

  frontend:
    image: my-frontend
    depends_on:
      - backend
```

With health-based conditions, a dependent service can wait for the dependency to become healthy.

**Self-assessment:** Shaky

---

# Quick-Fire Questions

## 1. What is the difference between an image and a container?

An **image** is a read-only template containing the application, dependencies, libraries, and configuration required to create a container.

A **container** is a running or stopped instance created from an image.

```text
Image → Container
```

---

## 2. What happens to data inside a container when you remove it?

Data stored only in the container's writable layer is removed when the container is removed.

Data stored in a named volume or bind mount persists separately from the container.

---

## 3. How do two containers on the same custom network communicate?

They communicate through the Docker network and can use container names for DNS-based service discovery.

Example:

```text
backend:8080
```

The containers do not need to know each other's IP addresses.

---

## 4. What does `docker compose down -v` do differently from `docker compose down`?

```bash
docker compose down
```

Stops and removes Compose containers and networks.

```bash
docker compose down -v
```

Also removes the volumes associated with the Compose project.

---

## 5. Why are multi-stage builds useful?

Multi-stage builds separate the build environment from the runtime environment.

The final image can contain only the application artifacts and runtime dependencies, reducing unnecessary image size and build dependencies.

---

## 6. What is the difference between COPY and ADD?

`COPY` copies files and directories into the image.

`ADD` provides additional functionality, such as handling local tar archives.

For normal file copying, `COPY` is generally preferred.

---

## 7. What does `-p 8080:80` mean?

```text
Host port 8080 → Container port 80
```

Example:

```bash
docker run -p 8080:80 nginx
```

The application must actually be listening on port `80` inside the container.

`EXPOSE 80` alone does not publish the port.

---

## 8. How do you check how much disk space Docker is using?

```bash
docker system df
```

It shows Docker's disk usage for images, containers, volumes, and build cache.

---

# Revisit Weak Spots

## Weak Spot 1 — CMD vs ENTRYPOINT

Create a Dockerfile:

```dockerfile
FROM alpine

ENTRYPOINT ["echo"]

CMD ["Hello"]
```

Build:

```bash
docker build -t cmd-entrypoint-demo .
```

Run normally:

```bash
docker run --rm cmd-entrypoint-demo
```

Run with a different argument:

```bash
docker run --rm cmd-entrypoint-demo "Docker World"
```

The `ENTRYPOINT` remains `echo`, while the runtime argument replaces the default `CMD`.

---

## Weak Spot 2 — Environment Variables in Compose

Create `.env`:

```text
APP_PORT=8080
```

Use it in Compose:

```yaml
services:
  frontend:
    image: nginx:alpine
    ports:
      - "${APP_PORT}:80"
```

Verify:

```bash
docker compose config
```

The value of `${APP_PORT}` should resolve to `8080`.

---

# Docker Port Mapping Reminder

A Docker port mapping connects a host port to a container port.

```bash
docker run -p 8080:80 nginx
```

means:

```text
Host:8080
    ↓
Container:80
    ↓
Nginx
```

The application inside the container must actually listen on the mapped container port.

For example, if the application listens on `4173`:

```text
Application → Container:4173
```

then:

```text
8080:80
```

will not reach it.

The correct mapping would be:

```text
8080:4173
```

if you intend to expose that application through host port `8080`.

---

# Docker Cheat Sheet

## Container Commands

```bash
docker run -it ubuntu bash
docker run -d nginx
docker ps
docker ps -a
docker stop <container>
docker rm <container>
docker exec -it <container> bash
docker logs <container>
```

## Image Commands

```bash
docker images
docker pull <image>
docker build -t <name>:<tag> .
docker tag <image> <username>/<repo>:<tag>
docker push <username>/<repo>:<tag>
docker rmi <image>
docker history <image>
```

## Volume Commands

```bash
docker volume create <volume>
docker volume ls
docker volume inspect <volume>
docker volume rm <volume>
```

## Network Commands

```bash
docker network create <network>
docker network ls
docker network inspect <network>
docker network connect <network> <container>
```

## Compose Commands

```bash
docker compose up -d
docker compose up -d --build
docker compose ps
docker compose logs -f
docker compose build
docker compose down
docker compose down -v
```

## Cleanup Commands

```bash
docker system df
docker system prune
docker system prune -a
```

## Dockerfile Instructions

```dockerfile
FROM
RUN
COPY
WORKDIR
EXPOSE
CMD
ENTRYPOINT
```

### Quick Reference

| Instruction | Purpose |
|---|---|
| `FROM` | Selects the base image |
| `RUN` | Executes a command during image build |
| `COPY` | Copies files into the image |
| `WORKDIR` | Sets the working directory |
| `EXPOSE` | Documents a container port |
| `CMD` | Provides the default command/arguments |
| `ENTRYPOINT` | Defines the main executable |

---

# Day 37 Summary

Day 37 revised:

- Running interactive and detached containers
- Container lifecycle
- Docker images and image layers
- Docker build cache
- Dockerfiles
- `CMD` vs `ENTRYPOINT`
- Building and tagging custom images
- Named volumes
- Bind mounts
- Custom networks
- Docker Compose
- Environment variables and `.env`
- Multi-stage builds
- Docker Hub
- Healthchecks
- `depends_on`
- Docker cleanup
- Port mapping

The goal of this revision was to understand not only **which Docker command to use**, but also **what the command actually does and why it is used**.
