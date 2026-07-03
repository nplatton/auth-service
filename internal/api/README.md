## internal/api

`responses.go`

Contains helper functions for constructing consistent HTTP responses.

`router.go`

Creates and configures the HTTP router.

`routes.go`

Registers every API endpoint with the router.

### `dto`

`auth.go`

Request and response models used by authentication endpoints.

`oidc.go`

DTOs used by OIDC endpoints.

`users.go`

DTOs relating to user operations.

### `handlers`

`auth.go`

Implements login, registration, logout and refresh HTTP handlers.

`oidc.go`

Implements the OAuth/OIDC protocol endpoints.

`users.go`

Handles user-related API requests.

`health.go`

Implements health and readiness endpoints.

### `middleware`

`auth.go`

Verifies JWTs and enforces authentication for protected endpoints.

`cors.go`

Adds Cross-Origin Resource Sharing headers.

`logging.go`

Logs incoming HTTP requests and responses.

`rate_limit.go`

Applies request rate limiting.

`recovery.go`

Recovers from panics and returns standard error responses.