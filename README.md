# blaze

![logo](https://github.com/paologaleotti/blaze/assets/45665769/a0c691df-b76b-4a4a-ac44-c622dd458352)

Simple and minimal Go template for building fast, simple and mantainable web services and backend applications.

You can find the **full documentation** with examples [here](https://github.com/paologaleotti/blaze/wiki).

A full REST API example using SQLite, sqlx and Prometheus metrics can be found [here](https://github.com/paologaleotti/blaze-api-example).

## Features

- Very minimal and zero overhead
- Production ready
- Simple and conventional structure
- Separation of business logic and API logic
- Fully compatible and based on standard [net/http](https://pkg.go.dev/net/http)
- Strict linting with [golangci-lint](https://golangci-lint.run/)
- Custom HTTP error handling and error mapping (separate service errors from API errors)
- Request payload validation
- Structured logging with [zerolog](https://github.com/rs/zerolog)
- Full AWS Lambda support (see [serverless](https://github.com/paologaleotti/blaze/tree/feature/serverless) branch)

All utilities are implemented in the `httpx` and `util` package.

## Stack

- **chi**: HTTP router (std net/http compatible)
- **chi/middleware**: default middlewares and utils
- **validator/v10**: request body struct validation
- **zerolog**: Structured logging

## Get started

You can start by reading the small [wiki](https://github.com/paologaleotti/blaze/wiki) with examples.

To scaffold a new blaze project, simply run this command:

```bash
go run github.com/paologaleotti/blaze-cli@master
```

This will use the [blaze-cli](https://github.com/paologaleotti/blaze-cli) to scaffold the project. You can also use the GitHub template or simply clone the repo.
