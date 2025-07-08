# blaze

![logo](https://github.com/paologaleotti/blaze/assets/45665769/a0c691df-b76b-4a4a-ac44-c622dd458352)

Minimal Go template for building fast, simple and mantainable web services and backend applications.

## Features

- Very minimal and zero overhead
- Production ready
- Simple and conventional structure, no black magic
- Separation of business logic and API logic
- Fully compatible and based on standard [net/http](https://pkg.go.dev/net/http)
- Strict linting with [golangci-lint](https://golangci-lint.run/)
- Custom HTTP error handling and error mapping (separate service errors from API errors)
- Request payload validation
- Structured logging with [zerolog](https://github.com/rs/zerolog)
- Full AWS Lambda support (see [serverless](https://github.com/paologaleotti/blaze/tree/feature/serverless) branch)

All utilities are implemented in the `httpx` and `util` package.

This repo contains a more real world example of a CRUD API with:

- sqlite3 storage
- Embedded sql migrations
- Apply migrations on startup with goose

## Base dependencies

- **chi**: HTTP router (std net/http compatible)
- **chi/middleware**: default middlewares and utils
- **validator/v10**: request body struct validation
- **zerolog**: structured logging

## Get started

You can start by reading the small [wiki](https://github.com/paologaleotti/blaze/wiki) with examples.

To get started, simply run:

```bash
make
```

This will build the project, you just need Go installed and nothing else.

To run the server:

```bash
./bin/api/main
```

This will create a sqlite3 file in the root of the project, run migrations and start the server on port 3000.
