# Go Authentication

A secure authentication system built with Go, featuring user registration, login, and password hashing using bcrypt.

## Features

- 🔐 **Secure Password Hashing**: Uses bcrypt for password encryption
- 👤 **User Registration**: Create new user accounts
- 🔑 **User Authentication**: Login with username/password
- 🗃️ **PostgreSQL Integration**: Robust database storage
- 🌍 **Environment Configuration**: Secure configuration management
- 📦 **Modular Architecture**: Clean separation of concerns

## Tech Stack

- **Language**: Go 1.21+
- **Database**: PostgreSQL
- **Password Hashing**: bcrypt
- **Environment Management**: godotenv
- **Database Driver**: lib/pq

## Project Structure

```
go_authentication/
├── main.go              # Application entry point
├── db/
│   └── db.go           # Database connection and initialization
├── models/
│   └── user.go         # User model and authentication logic
├── .env                # Environment variables (not in git)
├── go.mod              # Go module dependencies
└── README.md           # Project documentation
```

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 12+
- Git

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/CrispyChillies/go_authentication.git
   cd go_authentication
   ```

2. **Install dependencies**:

   ```bash
   go mod tidy
   ```

3. **Set up PostgreSQL database**:

   ```sql
   CREATE DATABASE go_authentication;
   CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       username VARCHAR(50) UNIQUE NOT NULL,
       password VARCHAR(255) NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
   ```

4. **Create environment file**:

   ```bash
   cp .env.example .env
   ```

5. **Configure environment variables** in `.env`:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=go_authentication
   DB_SSLMODE=disable
   ```

## Usage

### Development

Run the application in development mode:

```bash
go run main.go
```

### Production

Build and run the application:

```bash
# Build
go build -o go_authentication main.go

# Run
./go_authentication
```

### Using Make (Optional)

If you have a Makefile:

```bash
make run    # Run in development
make build  # Build for production
```

## API Examples

### User Registration

The current implementation demonstrates user registration:

```go
// Register a new user
err := models.RegisterUser(db.DB, "username", "password")
if err != nil {
    log.Fatal("Registration failed:", err)
}
```

## Environment Variables

| Variable      | Description       | Default             |
| ------------- | ----------------- | ------------------- |
| `DB_HOST`     | PostgreSQL host   | `localhost`         |
| `DB_PORT`     | PostgreSQL port   | `5432`              |
| `DB_USER`     | Database username | `postgres`          |
| `DB_PASSWORD` | Database password | _required_          |
| `DB_NAME`     | Database name     | `go_authentication` |
| `DB_SSLMODE`  | SSL mode          | `disable`           |

## Security Features

- **Password Hashing**: All passwords are hashed using bcrypt with default cost
- **SQL Injection Protection**: Uses parameterized queries
- **Environment Variables**: Sensitive data stored in environment variables
- **Password Exclusion**: Passwords excluded from JSON responses

## Development

### Adding New Features

1. **Database Models**: Add new models in `models/` directory
2. **Database Migrations**: Update database schema as needed
3. **Environment Config**: Add new variables to `.env` and `db.InitDB()`

### Code Structure

- **`main.go`**: Application bootstrap and dependency injection
- **`db/db.go`**: Database connection and configuration
- **`models/user.go`**: User-related database operations and business logic

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Troubleshooting

### Common Issues

1. **Database Connection Failed**:

   - Ensure PostgreSQL is running
   - Verify credentials in `.env` file
   - Check if database exists

2. **Missing Dependencies**:

   ```bash
   go mod tidy
   ```

3. **Environment Variables Not Loading**:
   - Ensure `.env` file is in project root
   - Check file permissions
   - Verify no trailing spaces in `.env`

### Getting Help

- Open an issue on GitHub
- Check the Go documentation: https://golang.org/doc/
- PostgreSQL documentation: https://www.postgresql.org/docs/

---

**Note**: This is a learning project demonstrating Go authentication concepts. For production use, consider additional security measures like JWT tokens, rate limiting, and comprehensive input validation.
