# Task Management API Documentation

## Overview
The Task Management API allows users to manage tasks with functionalities for creating, reading, updating, and deleting tasks. The system also supports user authentication and authorization using JWT, with role-based access control for admins and regular users. The project is implemented using the Clean Architecture pattern to ensure scalability, maintainability, and separation of concerns.

### Base URL
[http://localhost:8080](http://localhost:8080)

## Prerequisites
### Setting up MongoDB
1. **Install MongoDB**  
   Follow the instructions on the [MongoDB installation guide](https://www.mongodb.com/docs/manual/installation/) to install MongoDB.

2. **Start MongoDB Server**  
   Run the `mongod` command to start the MongoDB server.

3. **Install Project Dependencies**  
   Navigate to your project directory and run:
   ```sh
   go get .
   ```

### Setting up Environment Variables
Create a `.env` file in the root of your project with the following content:
```sh
JWT_SECRET=your_jwt_secret_key
MONGO_URI=mongodb://localhost:27017
```

## Project Structure

```sh
task-manager/
├── Delivery/
│   ├── main.go
│   ├── controllers/
│   │   └── controller.go
│   └── routers/
│       └── router.go
├── Domain/
│   └── domain.go
├── Infrastructure/
│   ├── auth_middleWare.go
│   ├── jwt_service.go
│   └── password_service.go
├── Repositories/
│   ├── task_repository.go
│   └── user_repository.go
└── Usecases/
    ├── task_usecases.go
    └── user_usecases.go
```

### Folder Structure Explanation

- **Delivery/:** Handles the incoming HTTP requests and responses.
  - **main.go:** Initializes the server and dependencies, and sets up the routing configuration.
  - **controllers/controller.go:** Manages HTTP requests and delegates them to the appropriate use cases.
  - **routers/router.go:** Configures the routes and initializes the Gin router.

- **Domain/:** Represents the core business logic and entities.
  - **domain.go:** Defines core entities like `Task` and `User` structs, encapsulating business rules.

- **Infrastructure/:** Implements external services and dependencies.
  - **auth_middleWare.go:** Middleware for JWT-based authentication and authorization.
  - **jwt_service.go:** Handles JWT token generation and validation.
  - **password_service.go:** Manages password hashing and comparison for secure credential storage.

- **Repositories/:** Contains interfaces and implementations for data access.
  - **task_repository.go:** Handles data access operations related to tasks.
  - **user_repository.go:** Handles data access operations related to users.

- **Usecases/:** Contains the application-specific business logic.
  - **task_usecases.go:** Implements the core functionalities for task management, such as creating, updating, and deleting tasks.
  - **user_usecases.go:** Implements user-related functionalities, such as registration, login, and role management.

## Authentication & Authorization

### JWT Authentication
The API uses JSON Web Tokens (JWT) for authenticating users. A valid JWT must be included in the `Authorization` header of each request to access protected routes.

### User Roles
- **Admin:** Full access to all endpoints, including managing tasks and promoting users.
- **User:** Limited access, only able to retrieve tasks.

### Middleware
- **AuthMiddleware:** Validates the presence and correctness of JWT tokens.
- **RoleMiddleware:** Ensures that users have the appropriate roles to access certain endpoints.

### Creating the First Admin User
If no users exist in the database, the first registered user will automatically be assigned the `admin` role.

## API Endpoints

### 1. User Registration
**URL:** `/register`  
**Method:** `POST`  
**Description:** Registers a new user. The first user is automatically assigned the admin role.

**Request Body:**
```json
{
    "email": "user@example.com",
    "password": "password123"
}
```

**Response:**
- **Status:** 201 Created
- **Body:**
```json
{
    "id": "user_id",
    "email": "user@example.com",
    "role": "admin"
}
```

### 2. User Login
**URL:** `/login`  
**Method:** `POST`  
**Description:** Authenticates a user and returns a JWT token.

**Request Body:**
```json
{
    "email": "user@example.com",
    "password": "password123"
}
```

**Response:**
- **Status:** 200 OK
- **Body:**
```json
{
    "message": "User logged in successfully!",
    "token": "jwt_token"
}
```

### 3. Get All Tasks
**URL:** `/tasks`  
**Method:** `GET`  
**Description:** Retrieves a list of all tasks.  
**Authorization:** Bearer `<JWT token>`

**Response:**
- **Status:** 200 OK
- **Body:**
```json
[
    {
        "id": "task_id",
        "title": "Task 1",
        "description": "Description of Task 1",
        "due_date": "2024-08-08T07:23:10Z",
        "status": "pending"
    }
]
```

### 4. Get a Task by ID
**URL:** `/tasks/:id`  
**Method:** `GET`  
**Description:** Retrieves a task by its ID.  
**Authorization:** Bearer `<JWT token>`

**Response:**
- **Status:** 200 OK
- **Body:**
```json
{
    "id": "task_id",
    "title": "Task 1",
    "description": "Description of Task 1",
    "due_date": "2024-08-08T07:23:10Z",
    "status": "pending"
}
```

### 5. Create a New Task
**URL:** `/tasks/`  
**Method:** `POST`  
**Description:** Creates a new task.  
**Authorization:** Bearer `<JWT token>` (Admin role required)

**Request Body:**
```json
{
    "title": "Task 1",
    "description": "Description of Task 1",
    "due_date": "2024-08-08T07:23:10Z",
    "status": "pending"
}
```

**Response:**
- **Status:** 201 Created
- **Body:**
```json
{
    "id": "task_id",
    "title": "Task 1",
    "description": "Description of Task 1",
    "due_date": "2024-08-08T07:23:10Z",
    "status": "pending"
}
```

### 6. Update a Task
**URL:** `/tasks/:id`  
**Method:** `PUT`  
**Description:** Updates an existing task.  
**Authorization:** Bearer `<JWT token>` (Admin role required)

**Request Body:**
```json
{
    "title": "Updated Task",
    "description": "Updated description",
    "due_date": "2024-08-12T12:00:00Z",
    "status": "completed"
}
```

**Response:**
- **Status:** 200 OK
- **Body:**
```json
{
    "id": "task_id",
    "title": "Updated Task",
    "description": "Updated description",
    "due_date": "2024-08-12T12:00:00Z",
    "status": "completed"
}
```

### 7. Delete a Task
**URL:** `/tasks/:id`  
**Method:** `DELETE`  
**Description:** Deletes a task by its ID.  
**Authorization:** Bearer `<JWT token>` (Admin role required)

**Response:**
- **Status:** 204 No Content

### 8. Promote User to Admin
**URL:** `/users/promote/:id`  
**Method:** `PUT`  
**Description:** Promotes a user to the admin role.  
**Authorization:** Bearer `<JWT token>` (Admin role required)

**Response:**
- **Status:** 200 OK
- **Body:**
```json
{
    "message": "User promoted successfully!"
}
```

## Design Decisions

### Clean Architecture
- **Separation of Concerns:** The architecture is designed to ensure that the business logic, data access, and delivery mechanisms are independent of each other.
- **Testability:** By decoupling the different layers, each part of the system can be tested in isolation.
- **Scalability:** The architecture allows for easy scaling of the application by introducing new features or modifying existing ones without affecting other parts of the system.

### Future Development Guidelines
- **Consistency:** Maintain the current folder structure and separation of concerns when adding new features.
- **Interfaces:** Ensure that all new functionalities are abstracted using interfaces, keeping implementation details hidden.
- **Documentation:** Update the documentation regularly to reflect any architectural changes or new features.
