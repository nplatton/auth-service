## internal/cache

`redis.go`

Creates and manages Redis connections.

`state.go`

Stores OAuth state parameters.

`pkce.go`

Stores PKCE challenges during OAuth flows.

`rate_limit.go`

Stores counters used for request rate limiting.

`sessions.go`

Caches active session information.