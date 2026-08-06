# Day 32 - Docker Volumes & Networking

## Objective

Today I learned how Docker Volumes and Docker Networking solve two important problems in containerized applications:
- **Data Persistence**
- **Container Communication**

---

# Task 1 - What Happens Without Volumes?

## Run PostgreSQL Container

```bash
docker run -d \
--name postgres-test \
-e POSTGRES_PASSWORD=admin \
-p 5432:5432 \
postgres
```

### Screenshot

![PostgreSQL Container](images/postgres-container.png)

---

Connect to PostgreSQL

```bash
docker exec -it postgres-test psql -U postgres
```

Create a table and insert data.

```sql
CREATE TABLE students(
id SERIAL PRIMARY KEY,
name VARCHAR(30)
);

INSERT INTO students(name)
VALUES('Sahil');

SELECT * FROM students;
```

### Screenshot

![Database Created](images/Created-database.png)

---

Stop and remove the container.

```bash
docker stop postgres-test
docker rm postgres-test
```

Run another PostgreSQL container without using a volume.

### Screenshot

![Recreated PostgreSQL Container](images/recreated-postgres.png)

### Observation

Since no volume was attached, Docker removed the container's writable layer along with the database files. As a result, all previously created data was lost.

---

# Task 2 - Named Volumes

Create a named volume.

```bash
docker volume create postgres-data
```

Run PostgreSQL using the volume.

```bash
docker run -d \
--name postgres-volume \
-e POSTGRES_PASSWORD=admin \
-v postgres-data:/var/lib/postgresql/data \
-p 5432:5432 \
postgres
```

### Screenshot

![Named Volume](images/named-volume.png)

After recreating the container with the same volume attached, all database records were still available.

### Observation

Named volumes store data outside the container lifecycle, allowing data to persist even after containers are deleted.

---

# Task 3 - Bind Mount

Create a folder containing an `index.html` file.

Run Nginx with a bind mount.

```bash
docker run -d \
--name nginx-bind \
-p 8080:80 \
-v $(pwd):/usr/share/nginx/html \
nginx
```

Edit the HTML file on the host and refresh the browser.

### Screenshot

![Bind Mount](images/nginx-edited-bindmount-.png)

### Observation

Changes made on the host machine were reflected immediately inside the running container without rebuilding the image.

---

# Task 4 - Docker Networking Basics

List available Docker networks.

```bash
docker network ls
```

Inspect the default bridge network.

```bash
docker network inspect bridge
```

Run two Ubuntu containers.

Ping using the container name.

### Screenshot

![Ping by Name](images/pingbyname.png)

Result:

```
Name or service not known
```

Ping using the container IP address.

### Screenshot

![Ping by IP](images/pingbyIP.png)

Result:

```
64 bytes from ...
```

### Observation

The default bridge network allows communication using IP addresses but does not provide automatic DNS-based name resolution for standalone containers.

---

# Task 5 - Custom Bridge Network

Create a custom bridge network.

```bash
docker network create my-app-net
```

Run two Ubuntu containers on the custom network.

Ping one container from the other using its container name.

### Screenshot

![Custom Network Inspect](images/custom-network-inspect.png)

### Screenshot

![Ping by Name](images/pingbyname.png)

### Observation

Docker automatically provides DNS-based service discovery on user-defined bridge networks, allowing containers to communicate using their names.

---

# Task 6 - Putting Everything Together

Create a custom network and a named volume.

Run a PostgreSQL container and an application container on the same network.

Verify connectivity.

### Screenshot

![Application Connected to Database](images/ping-database.png)

### Observation

The application container successfully communicated with the PostgreSQL container using the database container name instead of its IP address. This is the recommended approach for multi-container applications.

---

# Key Learnings

- Containers are ephemeral by design.
- Named volumes provide persistent storage.
- Bind mounts are useful for development because changes are reflected immediately.
- Default bridge networks support IP communication.
- Custom bridge networks provide built-in DNS for name-based communication.
- Combining volumes and custom networks enables production-ready multi-container deployments.
