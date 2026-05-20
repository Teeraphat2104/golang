# 🚀 User Management API

> A mock RESTful API built with **Go** and **Gin** framework for managing users

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![Gin](https://img.shields.io/badge/Gin-v1.9+-00ADD8?style=flat-square&logo=gin)](https://gin-gonic.com)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Active-brightgreen?style=flat-square)](https://github.com)

---

## 📑 Table of Contents

- [✨ Features](#-features)
- [⚡ Quick Start](#-quick-start)
- [📡 API Endpoints](#-api-endpoints)
- [🏗️ Architecture](#-architecture)
- [📝 Data Models](#-data-models)
- [💡 Usage Examples](#-usage-examples)
- [❌ Error Handling](#-error-handling)
- [📁 File Structure](#-file-structure)
- [🚀 Production Ready](#-production-ready)

---

## ✨ Features

| Feature | Description |
|---------|-------------|
| 🔄 **CRUD Operations** | Full Create, Read, Update, Delete functionality |
| 🏗️ **RESTful Design** | Standard REST API patterns and conventions |
| 📦 **JSON Format** | All requests/responses in JSON |
| 📊 **Mock Data** | Pre-loaded with 3 sample users ready to use |
| ❤️ **Health Check** | Endpoint to verify API status and availability |
| ✅ **Validation** | Email format validation and required field checking |
| 🆔 **Auto IDs** | Automatic incrementing ID assignment for new users |
| ⚙️ **Config** | Configurable port via environment variables |

---

## 🏗️ Architecture

```
┌─────────────────────────────────────────────────────┐
│         🎯 cmd/api/main.go                          │
│         Entry Point - Server Initialization         │
└──────────────────┬──────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────┐
│         🛣️  routes/routes.go                        │
│         Route Configuration & Setup                 │
└──────────────────┬──────────────────────────────────┘
                   │
         ┌─────────┴──────────┬──────────────┐
         ▼                    ▼              ▼
┌──────────────────┐  ┌──────────────┐  ┌─────────────┐
│ 🎮 handlers/     │  │ 📦 models/   │  │ ⚙️  config/ │
│ user_handler.go  │  │ user.go      │  │ database.go │
│ (Business Logic) │  │ (Structures) │  │ (Config)    │
└──────────────────┘  └──────────────┘  └─────────────┘
         │
    📊 Mock Data
         │
         ▼
   mockUsers []User
```

---

## ⚡ Quick Start

### Prerequisites
- ✅ Go 1.21 or higher
- ✅ Git
- ✅ curl or Postman (for testing)

### Installation

```bash
# Navigate to project directory
cd c:\Users\ASUS\hello

# Install dependencies
go mod download

# Start the API server
go run cmd/api/main.go
```

### Verify Installation

```bash
# Check if API is running
curl http://localhost:8080/health

# Expected response: {"status": "healthy"}
```

---

## 📡 API Endpoints

### Quick Reference

| Method | Endpoint | Action | Status |
|--------|----------|--------|--------|
| 🟢 GET | `/api/users` | Get all users | 200 OK |
| 🟢 GET | `/api/users/:id` | Get user by ID | 200/404 |
| 🔵 POST | `/api/users` | Create user | 201/400 |
| 🟡 PUT | `/api/users/:id` | Update user | 200/404 |
| 🔴 DELETE | `/api/users/:id` | Delete user | 200/404 |
| 🟢 GET | `/health` | Health check | 200 OK |

### Detailed Endpoint Documentation
```json
{
    "data": {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.com"
    }
}
```

**Error Response (404 Not Found):**
```json
{
    "error": "User not found"
}
```

---

### 3. POST /api/users
**Create New User**

- **Purpose:** Add a new user to the system
- **Request Method:** POST
- **Body:** JSON with name and email

**Request Body:**
```json
{
    "name": "Alice Wonder",
    "email": "alice@example.com"
}
```

**Success Response (201 Created):**
```json
{
    "data": {
        "id": 4,
        "name": "Alice Wonder",
        "email": "alice@example.com"
    },
    "message": "User created successfully"
}
```

**Error Response (400 Bad Request):**
```json
{
    "error": "Field validation for 'name' failed on the 'required' tag"
}
```

---

### 4. PUT /api/users/:id
**Update Existing User**

- **Purpose:** Modify user information
- **Request Method:** PUT
- **Parameter:** `id` (integer, required)
- **Body:** JSON with fields to update

**Request Body (partial update):**
```json
{
    "name": "John Updated"
}
```

**Success Response (200 OK):**
```json
{
    "data": {
        "id": 1,
        "name": "John Updated",
        "email": "john@example.com"
    },
    "message": "User updated successfully"
}
```

**Error Response (404 Not Found):**
```json
{
    "error": "User not found"
}
```

---

### 5. DELETE /api/users/:id
**Delete User**

- **Purpose:** Remove a user from the system
- **Request Method:** DELETE
- **Parameter:** `id` (integer, required)

**Success Response (200 OK):**
```json
{
    "message": "User deleted successfully"
}
```

**Error Response (404 Not Found):**
```json
{
    "error": "User not found"
}
```

---

### 6. GET /health
**Health Check**

- **Purpose:** Verify API is running
- **Request Method:** GET

**Response (200 OK):**
```json
{
    "status": "healthy"
}
```

---

## 💡 Usage Examples

### Complete Lifecycle Example

```bash
# 1. Start server
go run cmd/api/main.go

# 2. Get all users
curl http://localhost:8080/api/users

# 3. Create new user
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Sarah Connor", "email": "sarah@example.com"}'
# Returns: {"id": 4, ...}

# 4. Get the new user
curl http://localhost:8080/api/users/4

# 5. Update the user
curl -X PUT http://localhost:8080/api/users/4 \
  -H "Content-Type: application/json" \
  -d '{"email": "sarah.connor@example.com"}'

# 6. Delete the user
curl -X DELETE http://localhost:8080/api/users/4

# 7. Verify deletion (404 expected)
curl http://localhost:8080/api/users/4
```

---

## ❌ Error Handling

### HTTP Status Codes

| Code | Meaning | Scenario |
|------|---------|----------|
| 200 | ✅ OK | Successful GET, PUT, DELETE |
| 201 | ✅ Created | User successfully created |
| 400 | ❌ Bad Request | Invalid input or validation error |
| 404 | ❌ Not Found | User/resource doesn't exist |
| 500 | ❌ Server Error | Internal server error |

### Common Errors

**Missing Required Field (400):**
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John"}'
```
Response: `{"error": "Field validation for 'email' failed..."}`

**Invalid Email (400):**
```bash
curl -X POST http://localhost:8080/api/users \
  -d '{"name": "John", "email": "invalid"}'
```
Response: `{"error": "Field validation for 'email' failed..."}`

**User Not Found (404):**
```bash
curl http://localhost:8080/api/users/999
```
Response: `{"error": "User not found"}`

---

## 📁 File Structure

```
c:\Users\ASUS\hello\
├── 📄 README.md                ← Documentation
├── 📄 go.mod                   ← Dependencies
├── 📁 cmd/api/
│   └── main.go                 ← Server entry point
├── 📁 handlers/
│   └── user_handler.go         ← Business logic
├── 📁 models/
│   └── user.go                 ← Data structures
├── 📁 routes/
│   └── routes.go               ← Route definitions
├── 📁 config/
│   └── database.go             ← Config helpers
└── 📁 tmp/                      ← Temporary files
```

---

## 🚀 Production Ready

### To Deploy to Production

- 🗄️ **Real Database** - Connect PostgreSQL/MySQL
- 🔐 **Authentication** - Add JWT/OAuth2
- 📝 **Logging** - Structured logging with Zap
- 📊 **Pagination** - Add limit/offset
- ⚡ **Caching** - Redis integration
- 🛡️ **Rate Limiting** - Prevent abuse
- 🗺️ **API Docs** - Swagger/OpenAPI
- 🧪 **Testing** - Unit & integration tests
- 🌐 **CORS** - Enable cross-origin
- 🚀 **Deployment** - Docker/K8s ready

---

<div align="center">

### Made with ❤️ using Go & Gin

**[⬆ back to top](#-user-management-api)**

</div>

---

## Error Handling

### Error Response Format
All errors follow this structure:
```json
{
    "error": "Error description here"
}
```

### Common HTTP Status Codes

| Status | Meaning | Example |
|--------|---------|---------|
| 200 | OK | Successful GET, PUT, DELETE |
| 201 | Created | User successfully created (POST) |
| 400 | Bad Request | Invalid input data |
| 404 | Not Found | User ID doesn't exist |
| 500 | Server Error | Internal server error |

### Error Examples

**Missing Required Field (400):**
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John"}'
```

Response:
```json
{
    "error": "Field validation for 'email' failed on the 'required' tag"
}
```

---

## 📝 Data Models

### User Structure

```go
type User struct {
    ID    int    // Unique identifier
    Name  string // User full name (required)
    Email string // User email address (required)
}
```

### Initial Mock Data

```
User 1: John Doe (john@example.com)
User 2: Jane Smith (jane@example.com)
User 3: Bob Johnson (bob@example.com)
```
