# 🔐 Authentication
> The process of verifying a user's identity before granting access to a system or resource. It's essentially the mechanism that confirms "you are who you say you are" before allowing you to log in, access files, or use certain features within an application.

##### 🔶 Authentication vs Authorization
| Concept        | Meaning                     | Example                              |
| -------------- | --------------------------- | ------------------------------------ |
| Authentication | Who are you?                | Logging in with email + password     |
| Authorization  | What are you allowed to do? | Can this user access `/admin` route? |

➡️ You always authenticate first, then authorize.

- 💡 You must **authenticate** a user before you can **authorize** their actions.
- 🧠 Think: AuthN (identity) before AuthZ (permissions)
***
## 🔑 Token
> A token is a small piece of data (usually a string) used to identify or validate a user, application, or session between client and server.

**✅ Benefit**
- Decouples frontend and backend (no session on server)
- Stateless architecture (no need to store auth info per user in memory)
- Scalable and secure

### ⚙️ Types of Tokens
- #### 🔶 Session Token
  - Created on server; stored in a session database or in memory
  - Sent as a cookie to the browser
  - Used in classic web apps (e.g., PHP, Rails)
  - Simple but doesn’t scale well (needs server memory)
- #### 🔶 WT (JSON Web Token)
  - Self-contained token with three parts; `<header>.<payload>.<signature>`.
      - Header: info about signing algo (e.g., HS256)
      - Payload: your data (userId, roles, expiry...)
      - Signature: server signs (and later verifies) with secret key

      *🔹 example:*
      ```
      eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
      eyJ1c2VySWQiOiIxMjM0NSIsImV4cCI6MTY5ODk5OTk5OX0.
      SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
      ```
- #### 🔶 OAuth2 Access Token
  - Used in third-party logins (Google, GitHub)
  - You request an `access_token` from provider
  - Often a JWT, sometimes opaque
  - Backend validates token with provider’s public key or introspection endpoint
- #### 🔶 API Key
  - Static token for project or user
  - Sent with header (e.g., x-api-key)
  - Used for public APIs or internal tools
- #### 🔶 CSRF Token
  - Prevents attacks from other sites pretending to be you
  - Sent with every form submission
  - Mostly needed when using cookies for auth (not tokens)

### 🗝️ JWT (JSON Web Token)
> A JWT (JSON Web Token) is a compact, self-contained, URL-safe string used to securely transmit data between parties as a JSON object. It's digitally signed and optionally encrypted.

- It’s often used for **stateless authentication** *(a method of verifying user identity where the server doesn't store session information for each user.)*.
- Backend generates and signs a token after login.
- Frontend stores the token and sends it with each request.

#### Parts of JWT
JWT is URL-safe string: `<header>.<payload>.<signature>`
- #### 🔸 header: 
  The **header** tells us how the JWT is **signed** and what **type** of token it is.

    *🔹 example:*
    ```json
    {
      "alg": "HS256",
      "typ": "JWT"
    }
    ```
    | Field | Meaning                                            |
    | ----- | -------------------------------------------------- |
    | `alg` | The signing algorithm (e.g. `HS256` = HMAC-SHA256) |
    | `typ` | Type of token (usually `"JWT"`)                    |

- #### 🔸 payload:
  The **payload** contains the **claims** — information about the user or token.

    *🔹 example:*
    ```json
    {
      "userId": 123,
      "role": "admin",
      "exp": 1700000000
    }
    ```
    | Claim Type | Example                    | Description                           |
    | ---------- | -------------------------- | ------------------------------------- |
    | Registered | `exp`, `iat`, `sub`, `iss` | Standard fields for time and identity |
    | Public     | `name`, `email`, `role`    | General use (avoid conflicts)         |
    | Private    | `userId`, `teamId`         | App-specific data                     |

    **✅ Common Standard Claims:**
    - `exp`: Expiration timestamp (Unix)
    - `iat`: Issued at
    - `nbf`: Not before
    - `sub`: Subject (e.g., user ID)
    - `iss`: Issuer
    - `aud`: Audience

    **🧠 Note:**
    - Payload is not encrypted — it’s readable by anyone with the token.
    - But it should not contain sensitive info like passwords.
- #### 🔸 signature: 
  The **signature** ensures that the token hasn’t been tampered with.
  - Combines encoded header + payload
  - Signs with a secret or private key:

    **✅ How it's created:**
    ```text
    HMAC-SHA256(
      base64url(header) + "." + base64url(payload),
      secret
    )
    ```
    - This part is **critical**: it verifies that the header and payload were created by someone with the secret.
    - Only the backend should know the secret key.
    - When verifying the token, the backend **recomputes** the signature and checks if it matches.
***
#### Summary Table:
| Part      | What It Is                       | Used For                   |
| --------- | -------------------------------- | -------------------------- |
| Header    | Metadata (type + signing method) | How to verify the token    |
| Payload   | Claims (user info, expiry)       | Who the user is, validity  |
| Signature | HMAC or RSA signed hash          | Ensures data isn’t changed |

#### 💡 Advantages of JWT
| Feature          | Description                                    |
| ---------------- | ---------------------------------------------- |
| **Stateless**    | No server-side session storage needed          |
| **Scalable**     | Easy for microservices and distributed systems |
| **Portable**     | Works across languages, platforms, services    |
| **Flexible**     | Custom claims (e.g., role, userId, teamId)     |
| **Compact**      | Can be stored in localStorage, cookies, etc.   |
| **Tamper-proof** | Signed (and optionally encrypted)              |

#### ⚠️ JWT Limitations

| Limitations | Description |
| - | - |
| **No easy revocation** | can’t kill a token unless you track it manually          |
| **Size** | Larger than session IDs |
| **Leaked tokens** | equal to full access |
| **Not automatic invalidation** | Doesn’t support automatic invalidation on logout unless you store tokens in a blacklist (e.g., DB, Redis)|
---
### 🗝️ Session Token
> A Session Token is a random identifier (usually UUID) stored in server memory or a database, tied to a user session.

- User logs in → backend creates a session in memory or DB
- Session ID (token) is stored in a cookie and sent to client

#### 💡 Advantages of Session Token
| Feature               | Description                           |
| --------------------- | ------------------------------------- |
| **Easy to revoke**    | Delete session from DB and it's gone  |
| **Short and secure**  | Random UUID stored in HttpOnly cookie |
| **Logout support**    | Simple session deletion               |
| **Secure by default** | No sensitive data in token itself     |

#### ⚠️ Limitations of Session Token
- Requires server-side storage
- Harder to scale horizontally (requires shared session store like Redis)
- Not ideal for SPAs or mobile APIs that need stateless design
---
## 🧭 Access & Refresh Token Mechanism (Modern Auth Flow)

### 📄 Access Token (Short-lived JWT)

- Sent in `Authorization` header
- Contains user identity, role, permissions, etc.
- Expires in 5–15 minutes
- ⚠️ Once expired, can't use anymore

### 🔁 Refresh Token (Long-lived UUID or JWT)

- Sent once during login
- Stored in secure cookie or secure storage
- Used to get a new access token when it expires
- Should be stored in **SQLite/Redis** and invalidated on logout

### 🧭 Full Flow: Login, Access, Refresh
1. **Login:** 
    - User provides credentials
    - Backend: 
      - Creates access token (JWT)
      - Creates refresh token (UUID)
      - Stores refresh token in DB (with user ID and expiry)
    - Frontend: 
      - `access_token` (store in memory/localStorage)
      - `refresh_token` (store in HttpOnly cookie or secure store)
2. **Accessing Protected Route:**
    - Frontend in requests:
        ```http
          GET /api/profile
          Authorization: Bearer <access_token>
        ```
3. **When access token expires:**
    - Frontend send:
        ```http
          POST /refresh
          Cookie: refresh_token=<token>
        ```
    - Backend:
      - Validates refresh token from DB
      - Returns new access token
4. **Logout:**
    - Frontend:
      - Remove `access_token` and `refresh_token`
    - Backend:
      - Deletes refresh token from DB
      - Optionally, blacklists the JWT (if needed)
---
## 🔐 Summary: When to Use What?
| Use Case                   | Recommendation             |
| -------------------------- | -------------------------- |
| Traditional web apps       | Session tokens (cookies)   |
| SPAs, mobile apps          | JWT (access + refresh)     |
| Microservices architecture | JWT                        |
| Needs token revocation     | Session or refresh token   |
| OAuth2 login               | JWT or opaque access token |
***
## 🔑 RBAC (Role-Based Access Control)
In most real applications, users are not equal. Example:

- Simple User → Can only access their own data (profile, orders, etc.).
- Supervisor → Can view and manage other users’ data but maybe not system-wide settings.
- Admin → Full control: manage users, change system configuration, etc.

👉 Without role separation:
- Every user would have the same permissions (insecure).
- You couldn’t restrict sensitive actions (like deleting other users).
- Hard to scale your system (more roles will be added over time).

This is why we introduce Role-Based Access Control (RBAC).

### 🧑‍🧒‍🧒 All Common RBAC Strategies:
#### 🔶 Flat Roles (Single Role per User):
Simplest model; each user has exactly one role.\
&nbsp; &nbsp; Example: a person has one of role `user` or `supervisor` or `admin`.

**🛠️ Simple implementation: (DB design)**
```sql
  users(id, username, role, ...) --add column for role in user table
```

✅ Pros: Easy, fast, minimal code.\
❌ Cons: Inflexible if you later need multiple roles per user.


#### 🔶 Multi-Role per User (User ↔ Roles Many-to-Many):
A user can have multiple roles.\
&nbsp; &nbsp; Example: a person could be both `editor` and `moderator`.

**🛠️ Simple implementation: (DB design)**
```sql
  users(id, username, ...)
  roles(id, name)                -- e.g. "admin", "supervisor", "user"
  user_roles(user_id, role_id)   -- mapping table
```

✅ Pros: Flexible, scalable.\
❌ Cons: More complex queries, requires joins.


#### 🔶 Role Hierarchies (Inheritance):
Roles can inherit permissions from other roles.\
&nbsp; &nbsp; Example: `admin` > `supervisor` > `user`. If you’re an `admin`, you automatically get `supervisor` and user `permissions`.

✅ Pros: Cleaner than repeating permissions.\
❌ Cons: Hierarchy rules can get tricky if system grows.


#### 🔶 Permissions / Actions (Fine-Grained RBAC):
Instead of only roles, you define permissions (like "create_user", "delete_order").

&nbsp; &nbsp; Roles group permissions.\
&nbsp; &nbsp; Users get permissions through roles.

**🛠️ Simple implementation: (DB design)**
```sql
  users(id, username, ...)
  roles(id, name)
  permissions(id, action)                -- "view_reports", "delete_user"
  role_permissions(role_id, permission_id)
  user_roles(user_id, role_id)
```

✅ Pros: Very flexible, industry standard.\
❌ Cons: More complexity, need a permission engine.

> A permission engine is a central system (library, service, or module) that decides:\
👉 “Is this user allowed to do this action on this resource?”\
👉 sA permission engine is like a decision-maker for access control.


#### 🔶 Attribute-Based Access Control (ABAC):
Goes beyond RBAC — rules are based on attributes (user, resource, context).\
&nbsp; &nbsp; Example:\
&nbsp; &nbsp; &nbsp; &nbsp; User: `{ role: "user", department: "sales" }`\
&nbsp; &nbsp; &nbsp; &nbsp; Rule: `"user" can only view records where department = 'sales'`

✅ Pros: Powerful, fine-tuned.\
❌ Cons: Complex, harder to maintain.


#### 🔶 Contextual RBAC (Dynamic / Policy-Based):
Mix RBAC with context conditions:
- Time (only during working hours).
- Location (only from office IP).
- Ownership (user can edit only their own profile).

✅ Pros: Enterprise-level, very flexible.
❌ Cons: Setup and policy management overhead.


#### 🔶 Hybrid RBAC + ABAC (Modern Approach):
Many modern systems mix RBAC for coarse-grained control and ABAC for fine-grained checks.

&nbsp; &nbsp; Example:
&nbsp; &nbsp; &nbsp; &nbsp; RBAC ensures `admin` can access `/admin`.
&nbsp; &nbsp; &nbsp; &nbsp; ABAC ensures `admin` can only modify users in their organization.


***

### 🎯 Summary Table:

| Strategy           | Example                    | Best For                        |
| ------------------ | -------------------------- | ------------------------------- |
| Flat role per user | user.role = "admin"        | MVP, simple apps                |
| Multi-role         | user has many roles        | Medium apps needing flexibility |
| Hierarchical RBAC  | admin > supervisor > user  | Org structures                  |
| Permissions RBAC   | role → permissions → users | Enterprise apps                 |
| ABAC               | based on attributes        | Fine-grained access rules       |
| Contextual RBAC    | policies (time, location)  | Security-critical apps          |
| Hybrid RBAC + ABAC | both together              | Modern scalable systems         |

***

### ⚙️ Implementation Approaches:

1. Hardcoded roles in JWT claims (fastest, good for MVP).
    - Token contains "role": "admin".
    - Middleware checks claim.

2. DB lookup on each request (stronger security).
    - Token only contains user_id.
    - Middleware loads roles/permissions from DB.

3. Use RBAC/ABAC libraries:
    - Casbin (supports RBAC, ABAC, and policies.)
    - Oso policy engine for Go.
    - Auth0 / Keycloak (external providers).