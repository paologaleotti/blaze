# blaze🔥

Simple and minimal Go template for building fast and type-safe backend applications

> [!IMPORTANT]
> While blaze is already used and working flawlessly in production, it's still evolving and may have breaking changes in the future.

## Features
- Minimal and low overhead
- Production ready
- Simple and conventional structure
- API-first friendly
- Mainly based on `std` library
- HTTP utilities and error handling
- Fast and simple build tool (single `Makefile`)

## Stack
- **chi**: HTTP router (std net/http compatible)
- **chi/middleware**: middleware and hooks
- **chi/render**: request/reply payload utility
- **validator/v10**: request body struct validation
- **zap**: Structured logging
