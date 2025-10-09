# ESDC Backend (GO)

## API Documentation

This project includes Swagger/OpenAPI documentation for all API endpoints.

### Accessing the Documentation

1. Start the server: `go run main.go`
2. Open your browser and navigate to: `http://localhost:9090/swagger/index.html`

### Regenerating Documentation

After making changes to API endpoints or models, regenerate the documentation:

```bash
./generate-docs.sh
```

Or manually:

```bash
swag init
```

### API Features

- **Authentication**: JWT-based authentication
- **User Management**: Registration, login, profile management
- **Project Management**: CRUD operations for projects
- **File Upload**: Support for file uploads
- **Admin Panel**: Administrative functions
