# Day 32 - Docker Volumes & Networking

## Objective

Today I learned how Docker stores persistent data using volumes and how containers communicate using Docker networks.

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

Connect to PostgreSQL

```bash
docker exec -it postgres-test psql -U postgres
```

Create Database

```sql
CREATE TABLE students(
id SERIAL PRIMARY KEY,
name VARCHAR(30)
);

INSERT INTO students(name)
VALUES('Sahil');

SELECT * FROM students;
```

Exit

```
\q
```

Remove Container

```bash
docker stop postgres-test
docker rm postgres-test
```

Run another PostgreSQL container

```bash
docker run -d \
--name postgres-test \
-e POSTGRES_PASSWORD=admin \
-p 5432:5432 \
postgres
```

Check data

Result

```
Table not found.
```

### Observation

The container was removed.

Its writable layer was also deleted.

Therefore all database data was lost.

---

# Task 2 - Named Volumes

Create Volume

```bash
docker volume create postgres-data
```

Verify

```bash
docker volume ls
```

Run PostgreSQL with Volume

```bash
docker run -d \
--name postgres-volume \
-e POSTGRES_PASSWORD=admin \
-v postgres-data:/var/lib/postgresql/data \
-p 5432:5432 \
postgres
```

Create table

```sql
CREATE TABLE employees(
id SERIAL PRIMARY KEY,
name VARCHAR(50)
);

INSERT INTO employees(name)
VALUES('Shubham');
```

Stop and Remove

```bash
docker stop postgres-volume
docker rm postgres-volume
```

Run Fresh Container

```bash
docker run -d \
--name postgres-volume-new \
-e POSTGRES_PASSWORD=admin \
-v postgres-data:/var/lib/postgresql/data \
-p 5432:5432 \
postgres
```

Check Data

```sql
SELECT * FROM employees;
```

Output

```
 id |  name
----+---------
1   | Shubham
```

### Observation

Data still exists because Docker stores it inside the named volume instead of inside the container.

Inspect Volume

```bash
docker volume ls
docker volume inspect postgres-data
```

---

# Task 3 - Bind Mount

Create folder

```bash
mkdir website
cd website
```

Create HTML

```bash
nano index.html
```

Example

```html
<!DOCTYPE html>
<html>
<head>
<title>Docker Day 32</title>
</head>
<body>

<h1>Hello from Bind Mount!</h1>

</body>
</html>
```

Run Nginx

```bash
docker run -d \
--name nginx-bind \
-p 8080:80 \
-v $(pwd):/usr/share/nginx/html \
nginx
```

Visit

```
http://localhost:8080
```

Edit

```html
<h1>Docker Bind Mount Works!</h1>
```

Refresh browser

Changes appear instantly without rebuilding the image.

---

## Named Volume vs Bind Mount

| Named Volume | Bind Mount |
|--------------|------------|
| Managed by Docker | Managed by Host OS |
| Stores persistent application data | Shares host files directly |
| Better for databases | Better for development |
| Docker controls location | User chooses location |

---
