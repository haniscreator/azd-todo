## 📝 TODO Demo – Vue + Go + ClickHouse + Goose + DBT + Vitest

This is a small full-stack demo application built to practice the technologies, including:

- Go + Gin for the API
- Goose for database migrations
- ClickHouse as the analytical database
- Vue 3 + Vite for the frontend
- Vitest for frontend testing
- DBT + ClickHouse adapter for analytics modeling

The project is intentionally simple:
A TODO list app where the user only needs to enter a username (no authentication).
All todos are stored in ClickHouse and can be analyzed with DBT.


## 🚀 Features
### ✅ Frontend (Vue + Axios + Vite)

- Choose username (stored in localStorage)
- Add todos
- List todos for current user
- Mark a todo as complete
- Simple UI with reactive updates
- Basic unit tests using Vitest

### ✅ Backend (Go + Gin)
- POST /api/todos — create todo
- GET /api/todos?username=x — list todos
- PATCH /api/todos/:id/complete — complete todo
- Health endpoints: /health, /db-health

### ✅ Database (ClickHouse)
- Stores todos using MergeTree engine
- Fast analytical queries
- Uses background mutations for updates

### ✅ Goose Migrations
```bash
CREATE TABLE IF NOT EXISTS todos (
  id UInt64,
  username String,
  title String,
  is_completed UInt8,
  created_at DateTime,
  completed_at Nullable(DateTime)
)
ENGINE = MergeTree()
ORDER BY (username, id);
```

### ✅ DBT (Analytics)
DBT model: todo_stats.sql
```bash
SELECT
    username,
    toDate(created_at)        AS day,
    count()                   AS total_todos,
    sum(is_completed)         AS completed_todos,
    countIf(is_completed = 0) AS pending_todos
FROM default.todos
GROUP BY username, day
ORDER BY username, day
```

### 📦 Project Structure
```bash
/backend
    main.go
    db.go
    todo.go
    migrations/
        0001_create_todos_table.sql

/frontend
    src/
        App.vue
        api.ts
        todoUtils.ts
        todoUtils.test.ts
    vite.config.js
    package.json

/todo_analytics
    dbt_project.yml
    models/
        todo_stats.sql
```

### 🏗️ How to Run
## 1️⃣ Start ClickHouse
```bash
docker run -d --name clickhouse \
  -p 8123:8123 -p 9000:9000 \
  -e CLICKHOUSE_DB=default \
  -e CLICKHOUSE_USER=default \
  -e CLICKHOUSE_PASSWORD=secret123 \
  clickhouse/clickhouse-server
```
## 2️⃣ Apply Goose migrations
```bash
cd backend
GOOSE_DRIVER=clickhouse \
GOOSE_DBSTRING="clickhouse://default:secret123@localhost:9000/default" \
goose -dir migrations up
```

## 3️⃣ Run backend (Go)
```bash
cd backend
go run .
```
Backend runs at:
👉 http://localhost:8080

## 4️⃣ Run frontend (Vue)
cd frontend
npm install
npm run dev

Frontend runs at:
👉 http://localhost:5173

## 5️⃣ Run DBT (analytics)
Test connection:
```bash
dbt debug
```
Run models:
```bash
dbt run
```
Query results in ClickHouse UI:
```bash
SELECT * FROM todo_stats;
```

### 🧪 Testing (Vitest)
```bash
cd frontend
npm test
```
Simple example:
```bash
expect(summary.completed).toBe(1)
```

### 🧩 System Design Decisions
Why Go + Gin?
- Fast, lightweight, perfect for JSON APIs.

Why ClickHouse?
- Excellent for analytical queries (columnar storage)
- Very fast aggregation

Why Goose?
- Simple, versioned migrations for Go projects

Why DBT?
- Brings software engineering practices to SQL analytics
- Reproducible models, documentation, and tests

Why Vitest?
- Works seamlessly with Vite + Vue
- Faster than Jest
- Simple, modern test runner