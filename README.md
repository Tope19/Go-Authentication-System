# Go Authentication System

This is a RESTful authentication system built with Go, using the Gin web framework and GORM for database operations. It provides endpoints for user registration, login, password reset, and more.

## Features

- User registration
- User login with JWT token generation
- Forgot password functionality
- Reset password functionality
- Environment variable configuration
- Error logging

## Prerequisites

- Go 1.16+
- MySQL
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Tope19/Go-Authentication-System
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

- POST `/api/v1/auth/register` - Register a new user
- POST `/api/v1/auth/login` - Login and receive a JWT token
- POST `/api/v1/auth/forgot-password` - Request a password reset
- POST `/api/v1/auth/reset-password` - Reset password with a token

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

## Error Logging

Errors are logged to `application.log` in the project root directory.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)