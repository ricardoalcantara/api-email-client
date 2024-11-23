# Email Template Application

A web application built with Go and React for managing email templates and sending emails using custom SMTP configurations. This application allows you to create, save, and use email templates with dynamic content through JSON parsing.

## Features

- Email template management (create, edit, save, delete)
- Generate email templates UI, backed by hermes
- SMTP configuration
- Dynamic content injection using JSON objects
- React-based web interface
- Template preview functionality

## Project Structure

```
.
├── main.go
├── internal/
├── pkg/
├── frontend/          # React SPA
├── docker-compose.yml
└── README.md
```

## Installation

### Prerequisites

- Docker
- Docker Compose

### Quick Start

```bash
# Clone the repository
git clone [your-repository-url]

# Navigate to project directory
cd [project-directory]

# Start the application using Docker Compose
docker compose up -d
```

The application will be available at `http://localhost:3000`

## Configuration

Configure the application using environment variables in your `docker-compose.yml`:

```yaml
services:
  app:
    environment:
      - DB_CONNECTION=your_database_connection_string
      - JWT_SECRET=your_jwt_secret
      - SMTP_DEFAULT_HOST=smtp.example.com
```

## Usage

### Creating a Template

1. Navigate to the "Templates" section
2. Click "New Template"
3. Fill in the template details:
   - Template name
   - Subject line
   - HTML content
   - Add placeholders using `{{.VariableName}}` syntax (Go template format)

> You can also generate from UI a template using the "Generate" button

### Configuring SMTP

1. Go to "SMTP Settings"
2. Add your SMTP configuration:
   - Host
   - Port
   - Username
   - Password
   - SSL/TLS settings

### Sending Emails

```json
// Example JSON payload for sending an email
{
  "template_slug": "welcome-template",
  "smtp_slug": "main-smtp",
  "to": "user@example.com",
  "subject": "", // Optional
  "data": { // The JSON object to be used as the template data
    "UserName": "John Doe", // {{ .UserName }}
    "ActivationLink": "https://yourdomain.com/activate/123" // {{ .ActivationLink }}
  }
}
```

## API Reference

### Emails

- `POST /api/send` - Send email using template
- `POST /api/send/:id/send` - Resend email

## Development

The project consists of two main parts:

### Backend (Go)

The Go backend is in the root directory. To develop locally:

```bash
# Install Go dependencies
go mod download

# Run the backend server
go run cmd/main.go
```

### Frontend (React)

The React frontend is in the `frontend` directory:

```bash
# Navigate to frontend directory
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Support

For support, please open an issue in the GitHub repository