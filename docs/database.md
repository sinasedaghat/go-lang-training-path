🔔 A **database** is a system to store, manage, and retrieve structured data.

## 📚 Main Categories of Databases

### 1. Relational Databases (SQL)

- Stores data in tables (rows & columns) like Excel.
- Tables can relate to each other (foreign keys).
- Uses a language called SQL (Structured Query Language).

#### 🔶 Examples:

| Database             | Description                    |
| -------------------- | ------------------------------ |
| SQLite               | Lightweight, file-based        |
| MySQL                | Fast, open-source, widely used |
| PostgreSQL           | Advanced, feature-rich         |
| Oracle DB            | Enterprise-grade, paid         |
| Microsoft SQL Server | Used in Windows ecosystems     |

✅ Pros:

- Structured, reliable, ACID-compliant
- Mature tooling and ecosystem
- Good for complex queries and reports

❌ Cons:

- Less flexible with changing data structures
- Can be slower with huge unstructured data

---

### 2. Non-Relational Databases (NoSQL)

- Doesn’t use tables.
- Stores data in formats like:
  - JSON-like documents
  - Key-value pairs
  - Graphs
  - Columns

#### Types of NoSQL:

| Type            | Example         | How It Works                                    |
| --------------- | --------------- | ----------------------------------------------- |
| Document Store  | MongoDB         | Stores JSON-like documents                      |
| Key-Value Store | Redis, DynamoDB | Like a dictionary or map                        |
| Column Store    | Cassandra       | Stores data by columns, not rows                |
| Graph DB        | Neo4j           | Great for relationships (e.g., social networks) |

✅ Pros:

- Flexible schema (can change data structure easily)
- Great for fast reads/writes, horizontal scaling
- Ideal for unstructured or semi-structured data

❌ Cons:

- Not great for complex queries (especially joins)
- May sacrifice consistency for speed (CAP theorem)

---

### 3. In-Memory Databases

- Data stored in RAM instead of disk
- Extremely fast, often used for caching

#### 🔶 Examples:

| Database  | Description                 |
| --------- | --------------------------- |
| Redis     | Key-value store, super fast |
| Memcached | Lightweight caching         |

---

### 4. NewSQL

- Modern SQL databases that combine SQL features with NoSQL-like performance.

#### 🔶 Examples:

| Database       | Description              |
| -------------- | ------------------------ |
| CockroachDB    | Resilient, distributed   |
| Google Spanner | Globally distributed SQL |

---

## 📊 Comparison Table

| Feature            | Description                                                                                                                                                                  | Relational (SQL)        | NoSQL                        | In-Memory         |
| ------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------- | ---------------------------- | ----------------- |
| Schema             | A schema is the structure or blueprint of your data — it defines what kind of data you store, and how it’s organized.                                                        | Fixed                   | Flexible                     | Depends on DB     |
| Query Language     |                                                                                                                                                                              | SQL                     | Varies (MongoDB has its own) | Usually key-based |
| Speed (read/write) |                                                                                                                                                                              | Moderate                | High                         | Very High         |
| Relationships      | A relationship is how data in different tables or entities are connected to each other.                                                                                      | Strong support (JOINs)  | Weak                         | None              |
| Scalability        | Scalability means how your system handles more data or more users. (Vertical: Add more power to one server (CPU, RAM, SSD), Horizontal: Add more servers to share the load.) | Vertical (harder)       | Horizontal (easier)          | Vertical          |
| Transactions       | A transaction is a set of one or more database operations that should either all succeed or all fail together — never half-and-half.                                         | Strong (ACID)           | Often limited                | Limited or none   |
| Use Cases          |                                                                                                                                                                              | Banking, ERP, Analytics | Social media, real-time apps | Caching, queues   |

#### Transactions follow ACID rules:
| Letter | Meaning                              |
| ------ | ------------------------------------ |
| A      | Atomic: all or nothing               |
| C      | Consistent: valid state only         |
| I      | Isolated: safe from other operations |
| D      | Durable: survives crashes            |

---
## 📖 Terminology
### 🔷 Database OPen means:
  Create a **handle** your program uses to talk to the **DB**. It includes configuration (address/file, credentials, timeouts), but it doesn’t necessarily create a network connection yet.

### 🔷 Connection:
  A connection is one live "**session**" between your app and the DB:
  - For server DBs (Postgres/MySQL): a TCP socket + session state.
  - For SQLite: a file handle + internal session inside the process (no network).
---
  #### 🔶 Example:
  ```go
  db, err := sql.Open("sqlite3", "app.db") // returns a *sql.DB handle
  ```
  - sql.Open does not open a real connection immediately.
  - It prepares a pool manager (*sql.DB) and will open connections on demand.
---
#### 🔷🔷 connection pool:
*Creating*/*tearing down* a connection is expensive. A **pool**:
- Keeps some connections ready (already authenticated/handshaked).
- Hands them out to concurrent queries.
- Returns them as idle when a query finishes, to reuse for the next query.

**Benefits**: lower latency, higher throughput under load, controlled resource usage.