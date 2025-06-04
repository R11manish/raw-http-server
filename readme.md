# Raw HTTP Server

A lightweight, custom HTTP server implementation written in Go. This project demonstrates the fundamentals of HTTP server implementation from scratch, handling raw HTTP requests and responses.

## Features

- Custom HTTP server implementation
- Raw HTTP request handling
- HTTP response generation
- Connection management
- Configurable host and port

## Project Structure

```
.
├── main.go          # Application entry point
├── server/          # Server implementation
│   ├── server.go    # Main server logic
│   ├── response.go  # HTTP response handling
│   └── connection.go # Connection management
├── model/          # Data models
└── utils/          # Utility functions
```

## Prerequisites

- Go 1.x or higher

## Installation

1. Clone the repository:

```bash
git clone https://github.com/R11manish/raw-http-server.git
cd raw-http-server
```

2. Install dependencies:

```bash
go mod download
```

## Usage

Run the server:

```bash
go run main.go
```

The server will start on `127.0.0.1:8080` by default.

## Configuration

The server can be configured by modifying the following parameters in `main.go`:

- Host: Default is "127.0.0.1"
- Port: Default is 8080

## Development

The project is structured into several key components:

- `server.go`: Contains the main HTTP server implementation
- `response.go`: Handles HTTP response generation and formatting
- `connection.go`: Manages client connections and request processing

## License

[Add your license here]

## Contributing

[Add contribution guidelines here]
