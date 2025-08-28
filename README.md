# Go Authentication

A robust and secure authentication API built with Go, featuring user registration, login, password hashing with bcrypt, and PostgreSQL integration.

## Features

- 🔐 **Password Hashing:** Secure user passwords using bcrypt
- 👤 **User Registration & Login:** Simple endpoints for account creation and authentication
- 🗃️ **PostgreSQL Database:** Persistent user storage
- 🌱 **Environment Variables:** Easy configuration with `.env` and `godotenv`
- 🧩 **Modular Design:** Clean separation of database, models, and handlers

## Tech Stack

- **Go** (1.21+)
- **PostgreSQL**
- **bcrypt** (`golang.org/x/crypto/bcrypt`)
- **godotenv** for environment management
- **lib/pq** PostgreSQL driver

## Project Structure

```
go_authentication/
├── main.go              # Application entry point
├── db/
│   └── db.go            # Database connection logic
├── models/
│   └── user.go          # User model and DB operations
├── handlers/            # HTTP handlers (register, login, etc.)
├── utils/               # Utility functions (e.g., JWT, password hashing)
├── .env                 # Environment variables (not committed)
├── go.mod
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 12+
- Git

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/CrispyChillies/go_authentication.git
   cd go_authentication
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Set up PostgreSQL:**

   ```sql
   CREATE DATABASE go_authentication;
   CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       username VARCHAR(50) UNIQUE NOT NULL,
       password VARCHAR(255) NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
   ```

4. **Configure environment variables:**
   - Copy the example file and edit it:
     ```bash
     cp .env.example .env
     ```
   - Fill in your `.env` file:
     ```env
     DB_HOST=localhost
     DB_PORT=5432
     DB_USER=postgres
     DB_PASSWORD=your_password
     DB_NAME=go_authentication
     DB_SSLMODE=disable
     JWT_SECRET=your_jwt_secret_key_here
     ```

### Running the Application

**Development:**

```bash
go run main.go
```

**Production:**

```bash
go build -o go_authentication main.go
./go_authentication
```

**With Makefile (optional):**

```bash
make run    # Run in development
make build  # Build for production
```

## API Endpoints

- `POST /register` — Register a new user
- `POST /login` — Authenticate and receive a JWT
- `GET /api/profile` — Get user profile (JWT required)

## Environment Variables

| Variable    | Description        | Example/Default     |
| ----------- | ------------------ | ------------------- |
| DB_HOST     | PostgreSQL host    | localhost           |
| DB_PORT     | PostgreSQL port    | 5432                |
| DB_USER     | Database username  | postgres            |
| DB_PASSWORD | Database password  | your_password       |
| DB_NAME     | Database name      | go_authentication   |
| DB_SSLMODE  | SSL mode           | disable             |
| JWT_SECRET  | JWT signing secret | (generate securely) |

## Security

- Passwords are hashed with bcrypt before storage
- SQL queries use parameterization to prevent injection
- JWT secret is loaded from environment variables
- Passwords are never returned in API responses

## Troubleshooting

- **Database connection errors:** Check your `.env` values and ensure PostgreSQL is running.
- **Dependencies:** Run `go mod tidy` if you see missing package errors.
- **Environment not loading:** Ensure `.env` is in the project root and properly formatted.

## Contributing

1. Fork the repo
2. Create a feature branch
3. Commit and push your changes
4. Open a Pull Request

## License

MIT License

---

**Note:** This project is for educational purposes. For production, implement additional security such as HTTPS, JWT expiration, refresh tokens, and input validation.
