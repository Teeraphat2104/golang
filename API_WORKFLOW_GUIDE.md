# API Workflow Guide - Complete Process Flow

## Table of Contents
1. [Overview](#overview)
2. [Architecture](#architecture)
3. [Startup Process](#startup-process)
4. [Request/Response Flow](#requestresponse-flow)
5. [Data Models](#data-models)
6. [API Endpoints](#api-endpoints)
7. [Usage Examples](#usage-examples)
8. [Error Handling](#error-handling)

---

## Overview

This is a **Mock User Management API** built with Go and Gin framework. It simulates a RESTful API with in-memory data storage for managing user records.

**Key Features:**
- CRUD operations (Create, Read, Update, Delete)
- RESTful endpoint design
- JSON request/response format
- Mock data included
- Health check endpoint

---

## Architecture

```
┌─────────────────────────────────────────────────────┐
│                   cmd/api/main.go                    │
│              (Entry Point - Server Start)            │
└──────────────────┬──────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────┐
│              routes/routes.go                        │
│          (Route Configuration & Setup)              │
└──────────────────┬──────────────────────────────────┘
                   │
         ┌─────────┴──────────┬──────────────┐
         ▼                    ▼              ▼
┌──────────────────┐  ┌──────────────┐  ┌─────────────┐
│ handlers/        │  │ models/      │  │ config/     │
│ user_handler.go  │  │ user.go      │  │ database.go │
│ (Business Logic) │  │ (Data Types) │  │ (Config)    │
└──────────────────┘  └──────────────┘  └─────────────┘
         │
    (Mock Data)
         │
         ▼
   mockUsers []User
```

---

## Startup Process

### Step 1: Server Initialization
```
go run cmd/api/main.go
    ↓
1. Gin router created
2. Routes registered (SetupRoutes called)
3. Server listens on port 8080
4. Ready to receive requests
```

### Detailed Flow:

**File: `cmd/api/main.go`**

1. **Initialize Gin Router**
   ```go
   router := gin.Default()
   ```

2. **Register All Routes**
   ```go
   routes.SetupRoutes(router)
   ```
   - Calls `routes/routes.go`
   - Sets up `/api/users` group with all endpoints
   - Adds `/health` endpoint

3. **Start Server**
   ```go
   router.Run(":8080")
   ```
   - Server listens on `http://localhost:8080`
   - Waiting for incoming HTTP requests

---

## Request/Response Flow

### Typical Request Lifecycle:

```
Client Request
    │
    ▼
Hits: http://localhost:8080/api/users/1
    │
    ▼
Gin Router matches route pattern
    │
    ▼
Route calls appropriate handler function
    │
    ▼
Handler processes request:
    ├─ Parse request body (if POST/PUT)
    ├─ Validate input
    ├─ Access mock data
    ├─ Perform operation
    └─ Build response
    │
    ▼
Handler returns JSON response
    │
    ▼
HTTP Status Code + JSON Body
    │
    ▼
Client receives response
```

---

## Data Models

### User Struct
```go
type User struct {
    ID        int       // Unique identifier
    Name      string    // User full name
    Email     string    // User email address
    CreatedAt time.Time // Creation timestamp
    UpdatedAt time.Time // Last update timestamp
}
```

### Request Models

**CreateUserRequest** (for POST)
```go
{
    "name": "John Doe",      // Required
    "email": "john@example.com" // Required, must be valid email
}
```

**UpdateUserRequest** (for PUT)
```go
{
    "name": "Jane Doe",      // Optional
    "email": "jane@example.com" // Optional
}
```

### Mock Data

The API starts with 3 pre-loaded users:
```go
var mockUsers = []User{
    {ID: 1, Name: "John Doe", Email: "john@example.com"},
    {ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
    {ID: 3, Name: "Bob Johnson", Email: "bob@example.com"},
}
```

---

## API Endpoints

### 1. GET /api/users
**Get All Users**

- **Purpose:** Retrieve all users from the system
- **Request Method:** GET
- **No body required**

**Response (200 OK):**
```json
{
    "data": [
        {
            "id": 1,
            "name": "John Doe",
            "email": "john@example.com"
        },
        {
            "id": 2,
            "name": "Jane Smith",
            "email": "jane@example.com"
        }
    ],
    "total": 2
}
```

---

### 2. GET /api/users/:id
**Get Single User by ID**

- **Purpose:** Retrieve a specific user
- **Request Method:** GET
- **Parameter:** `id` (integer, required)

**Success Response (200 OK):**
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

## Usage Examples

### 1. Start the API
```bash
cd c:\Users\ASUS\hello
go run cmd/api/main.go
```

**Output:**
```
Starting API server on port 8080...
[GIN-debug] GET    /api/users            --> hello/handlers.GetAllUsers
[GIN-debug] GET    /api/users/:id        --> hello/handlers.GetUserByID
[GIN-debug] POST   /api/users            --> hello/handlers.CreateUser
[GIN-debug] PUT    /api/users/:id        --> hello/handlers.UpdateUser
[GIN-debug] DELETE /api/users/:id        --> hello/handlers.DeleteUser
[GIN-debug] GET    /health               --> main.main.func1
[GIN-debug] Listening and serving HTTP on :8080
```

---

### 2. Get All Users
```bash
curl http://localhost:8080/api/users
```

**Response:**
```json
{
    "data": [
        {"id": 1, "name": "John Doe", "email": "john@example.com"},
        {"id": 2, "name": "Jane Smith", "email": "jane@example.com"},
        {"id": 3, "name": "Bob Johnson", "email": "bob@example.com"}
    ],
    "total": 3
}
```

---

### 3. Get Specific User
```bash
curl http://localhost:8080/api/users/1
```

**Response:**
```json
{
    "data": {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.com"
    }
}
```

---

### 4. Create New User
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Sarah Connor",
    "email": "sarah@example.com"
  }'
```

**Response:**
```json
{
    "data": {
        "id": 4,
        "name": "Sarah Connor",
        "email": "sarah@example.com"
    },
    "message": "User created successfully"
}
```

**After this:**
- Total users: 4
- New user ID: 4
- Next new user will get ID: 5

---

### 5. Update User
```bash
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe Updated",
    "email": "john.updated@example.com"
  }'
```

**Response:**
```json
{
    "data": {
        "id": 1,
        "name": "John Doe Updated",
        "email": "john.updated@example.com"
    },
    "message": "User updated successfully"
}
```

---

### 6. Delete User
```bash
curl -X DELETE http://localhost:8080/api/users/3
```

**Response:**
```json
{
    "message": "User deleted successfully"
}
```

**After this:**
- User with ID 3 removed from list
- Total users: 3

---

### 7. Health Check
```bash
curl http://localhost:8080/health
```

**Response:**
```json
{
    "status": "healthy"
}
```

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

**Invalid Email Format (400):**
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John",
    "email": "not-an-email"
  }'
```

Response:
```json
{
    "error": "Field validation for 'email' failed on the 'email' tag"
}
```

**User Not Found (404):**
```bash
curl http://localhost:8080/api/users/999
```

Response:
```json
{
    "error": "User not found"
}
```

**Invalid User ID Format (400):**
```bash
curl http://localhost:8080/api/users/abc
```

Response:
```json
{
    "error": "Invalid user ID"
}
```

---

## Complete Request/Response Workflow Example

### Scenario: Create a user, then retrieve it

**Step 1: Start Server**
```bash
go run cmd/api/main.go
```

**Step 2: Create User**
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Tom Hardy", "email": "tom@example.com"}'
```

Response:
```json
{
    "data": {
        "id": 4,
        "name": "Tom Hardy",
        "email": "tom@example.com"
    },
    "message": "User created successfully"
}
```

**Step 3: Get the Created User**
```bash
curl http://localhost:8080/api/users/4
```

Response:
```json
{
    "data": {
        "id": 4,
        "name": "Tom Hardy",
        "email": "tom@example.com"
    }
}
```

**Step 4: Get All Users**
```bash
curl http://localhost:8080/api/users
```

Response:
```json
{
    "data": [
        {"id": 1, "name": "John Doe", "email": "john@example.com"},
        {"id": 2, "name": "Jane Smith", "email": "jane@example.com"},
        {"id": 3, "name": "Bob Johnson", "email": "bob@example.com"},
        {"id": 4, "name": "Tom Hardy", "email": "tom@example.com"}
    ],
    "total": 4
}
```

**Step 5: Update the User**
```bash
curl -X PUT http://localhost:8080/api/users/4 \
  -H "Content-Type: application/json" \
  -d '{"email": "tom.hardy@example.com"}'
```

Response:
```json
{
    "data": {
        "id": 4,
        "name": "Tom Hardy",
        "email": "tom.hardy@example.com"
    },
    "message": "User updated successfully"
}
```

**Step 6: Delete the User**
```bash
curl -X DELETE http://localhost:8080/api/users/4
```

Response:
```json
{
    "message": "User deleted successfully"
}
```

**Step 7: Verify Deletion**
```bash
curl http://localhost:8080/api/users/4
```

Response:
```json
{
    "error": "User not found"
}
```

---

## File Structure

```
c:\Users\ASUS\hello\
├── cmd/
│   └── api/
│       └── main.go           ← Server entry point
├── config/
│   └── database.go           ← Database config
├── handlers/
│   └── user_handler.go       ← Business logic & handlers
├── models/
│   └── user.go               ← Data structures
├── routes/
│   └── routes.go             ← Route definitions
├── main.go                   ← Original main file
├── go.mod                    ← Go module file
└── API_WORKFLOW_GUIDE.md     ← This file
```

---

## Key Points to Remember

1. **Mock Data**: All user data is stored in memory (`mockUsers` slice) - resets on server restart
2. **Auto-Incrementing IDs**: New users get the next ID automatically
3. **Validation**: Email must be valid format, name is required
4. **No Persistence**: Data is not saved to database (mock only)
5. **Port**: Default is 8080, changeable via PORT environment variable
6. **CORS**: Not configured (can be added if needed)

---

## Next Steps to Production

If converting to production:
1. Connect to real database (PostgreSQL, MySQL, etc.)
2. Add authentication (JWT tokens)
3. Add logging and monitoring
4. Add pagination for large datasets
5. Add caching layer
6. Add request rate limiting
7. Add error tracking
8. Add API documentation (Swagger/OpenAPI)
9. Add unit and integration tests
10. Configure CORS for frontend

---

**Last Updated:** 2026-05-20
