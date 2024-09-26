# Go Authentication and Blog System

This is a RESTful authentication and blog system built with Go, using the Gin web framework and GORM for database operations. It provides endpoints for user authentication, blog management, and commenting.

## Features

- User authentication (registration, login, password reset)
- Blog post creation, retrieval, updating, and deletion
- Blog categories
- Commenting system
- JWT-based authentication for protected routes
- Environment variable configuration
- Error logging

## Prerequisites

- Go 1.16+
- MySQL
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Tope19/Go-Blog-System
```
```bash
cd go-authentication
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up your environment variables:
Copy the `.env.example` file to `.env` and fill in your specific details:
```bash
cp .env.example .env
```

4. Set up your MySQL database:
Create a new database for the project.

## Configuration

Update the `.env` file with your specific configuration:
```bash
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_HOST=localhost
DB_PORT=3306
JWT_SECRET=your_jwt_secret_key
SERVER_PORT=8080
```

## Running the Application

To start the server, run:
```bash
go run main.go
```

The server will start on the port specified in your `.env` file (default is 8000).

## API Endpoints

## Authentication

- POST `/api/v1/auth/register` - Register a new user
- POST `/api/v1/auth/login` - Login and receive a JWT token
- POST `/api/v1/auth/forgot-password` - Request a password reset
- POST `/api/v1/auth/reset-password` - Reset password with a token

### Blog System (Protected Routes)
All blog system routes require authentication. Include the JWT token in the Authorization header:
```bash
Authorization: Bearer <your_jwt_token>
```

- POST `/api/v1/blogs` - Create a new blog post
- GET `/api/v1/blogs` - List all blog posts
- GET `/api/v1/blogs/:id` - Get a specific blog post
- PUT `/api/v1/blogs/:id` - Update a blog post
- DELETE `/api/v1/blogs/:id` - Delete a blog post
- POST `/api/v1/categories` - Create a new category
- GET `/api/v1/categories` - List all categories
- GET `/api/v1/categories/:id` - Get a specific category
- PUT `/api/v1/categories/:id` - Update a category
- DELETE `/api/v1/categories/:id` - Delete a category
- POST `/api/v1/comments` - Create a new blog comment
- GET `/api/v1/comments` - List all comments
- GET `/api/v1/comments/:id` - Get a specific comment
- PUT `/api/v1/comments/:id` - Update a comment
- DELETE `/api/v1/comments/:id` - Delete a comment


## Testing

You can use Postman or any API testing tool to test the endpoints. Here are some example requests:

1. Register a new user:
POST http://localhost:8000/api/v1/auth/register
Content-Type: application/json
```bash
{
"first_name": "John",
"last_name": "Doe",
"email": "john.doe@example.com",
"password": "securepassword123"
}
```

2. Login:
POST http://localhost:8000/api/v1/auth/login
Content-Type: application/json
```bash
{
"email": "john.doe@example.com",
"password": "securepassword123"
}
```

3. Create a blog post (authenticated):
POST http://localhost:8080/api/v1/blogs/create
Content-Type: application/json
Authorization: Bearer <your_jwt_token>
```bash
{
"title": "My First Blog Post",
"content": "This is the content of my first blog post.",
"category_id": 1
}
```

## Error Logging

Errors are logged to `application.log` in the project root directory.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
