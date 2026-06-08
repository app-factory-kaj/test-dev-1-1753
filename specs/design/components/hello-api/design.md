---
type: service
language: Go
buildpack: docker
appPath: hello-api
entrypoint: deployment/service
---

# hello-api

Implement a public Go HTTP API on port 9090 with two endpoints:
- GET /hello — returns {"message": "Hello, World!"} with 200 status
- GET /health — returns "OK" with 200 status (no auth required, for platform readiness probe)

Use net/http from the standard library. No authentication, no CORS middleware (this is a public API). Module path: example.com/hello-api, Go version 1.25.

Dockerfile: Use FROM golang:1.25-alpine AS builder (HARD requirement — the build pod runs with GOTOOLCHAIN=local and will NOT auto-download a newer Go toolchain). Multi-stage build to alpine:3.20 runtime, expose port 9090.

No database, no persistence. This is a stateless service returning a static greeting.
