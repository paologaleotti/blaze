# blaze

![logo](https://github.com/paologaleotti/blaze/assets/45665769/a0c691df-b76b-4a4a-ac44-c622dd458352)

Simple and minimal Go template for building fast, simple and mantainable web services and backend applications.

> [!IMPORTANT]
> While blaze is already used and working flawlessly in production, it's still evolving and may have breaking changes in the future.

You can find the **full documentation** with examples [here](https://github.com/paologaleotti/blaze/wiki).

## Features

- Minimal and low overhead
- Production ready
- Simple and conventional structure
- Fully compatible and base on standard [net/http](https://pkg.go.dev/net/http)
- Strict linting with [golangci-lint](https://golangci-lint.run/)
- Custom HTTP error handling
- Request payload validation
- Structured logging with [zerolog](https://github.com/rs/zerolog)
- Full AWS Lambda support (see [serverless](https://github.com/paologaleotti/blaze/tree/feature/serverless) branch)

All utilities are implemented in the `httpcore` and `util` package.

## Stack

- **chi**: HTTP router (std net/http compatible)
- **chi/middleware**: middleware and hooks
- **validator/v10**: request body struct validation
- **zerolog**: Structured logging
