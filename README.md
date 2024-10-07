# Go RESTful API Server

This project is a RESTful API server built using the Go programming language. It leverages the Gin Web Framework to handle HTTP requests, `go-sqlite` for efficient database operations, and JWT for secure user authentication. The server is designed to manage user registration, event creation, and CRUD operations for events, making it an excellent foundation for a backend service.

## Technologies Used
- **Go (Golang)**: Primary language for building the API.
- **Gin Web Framework**: Lightweight framework for handling HTTP requests and routing.
- **go-sqlite**: A SQLite library for Go, providing easy database integration.
- **jwt-go**: Library for generating and verifying JSON Web Tokens (JWT) for user authentication.

## Key Features
- **User Management**:
  - **User Signup**: Endpoint for creating new user accounts with validation.
  - **User Login**: Secure login with JWT-based token generation.
  - **Credential Validation**: Middleware checks to ensure only authenticated users can access certain endpoints.
- **Event Management**:
  - **Event Creation**: Add new events with details such as name, description, and date.
  - **CRUD Operations**: Full Create, Read, Update, and Delete support for events, enabling flexible event management.
- **Authentication**:
  - **JWT Token Generation**: Generates secure tokens upon login to maintain session integrity.
  - **JWT Verification**: Ensures each request is from an authenticated user.
- **Middleware**: Implements middleware for authentication and error handling to maintain security and reliability.

## API Endpoints

- **User Endpoints**:
  - `POST /signup` - Register a new user.
  - `POST /login` - Authenticate and obtain a JWT token.
- **Event Endpoints**:
  - `GET /events` - Retrieve all events.
  - `POST /events` - Create a new event (requires authentication).
  - `PUT /events/{id}` - Update an existing event (requires authentication).
  - `DELETE /events/{id}` - Delete an event (requires authentication).

### Example Request for Creating an Event

```bash
curl -X POST http://localhost:8080/events \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <JWT_TOKEN>" \
-d '{"name": "Conference", "description": "Tech conference", "date": "2023-12-01"}'
