# Golang Rate Limiter
A lightweight and efficient rate limiter implemented in Go, designed to control the rate of incoming HTTP requests. This project provides middleware and utility functions for request throttling, making it suitable for APIs, web servers, and Go applications that need to prevent abuse and ensure fair resource usage.

## Features

- Implements the fixed window algorithm for rate limiting
- Simple integration with HTTP handlers as middleware
- In-memory storage backend for fast performance

## Getting Started

Install the package:

```bash
go get github.com/luk3skyw4lker/rate-limiter
```

Import and use in your project:

```go
import "github.com/luk3skyw4lker/rate-limiter"
```

## Usage

```go
limiter := ratelimiter.New(5, time.Minute) // 5 requests per minute
http.Handle("/", limiter.Middleware(yourHandler))
```

## License

MIT
