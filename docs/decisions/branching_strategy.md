
# Project Branching Strategy

```
main
└── develop
    ├── epic/0-bootstrap
    ├── epic/1-users
    ├── epic/2-sessions
    ├── epic/3-jwt
    ├── epic/4-oidc
    ├── epic/5-oauth
    ├── epic/6-google
    └── epic/7-hardening
```

Inside each epic branch:

```
epic/1-users
    ├── auth-5-user-model
    ├── auth-6-user-repository
    ├── auth-7-register-service
    ├── auth-8-register-endpoint
```

...

...

...

# EPIC 0 — Bootstrap (foundation)

Branch: `epic/0-bootstrap`

## Commit 1 — project scaffold

### AUTH-1

Initialize Go module and base folder structure

Includes:
- `cmd/auth-server/main.go`
- `internal/api/router.go` (minimal HTTP server)
- `/health` endpoint returning 200

## Commit 2 — config system

### AUTH-2

Add configuration loader with YAML + environment support

- `internal/config/config.go`
- `loader.go`
- `validator.go`

## Commit 3 — logging system

### AUTH-3

Add structured logging with request-aware context helpers

- `internal/logging/logger.go`
- `context.go`
- middleware stub logs request start/stop

## Commit 4 — docker support

### AUTH-4

Add Dockerfile and docker-compose for local development

- container runs `/health`

---

# EPIC 1 — Users

Branch: `epic/1-users`

## Commit 5 — user model

### AUTH-5

Define user domain model

- `internal/users/model.go`

## Commit 6 — filesystem persistence

### AUTH-6

Implement filesystem-based user repository

- `internal/storage/filesystem/*`
- `internal/users/repository.go`

## Commit 7 — password hashing

### AUTH-7a

Add Argon2 password hashing utility

- `crypto/argon.go`

### AUTH-7b

Implement user registration service

- `users/service.go`

## Commit 8 — register endpoint

### AUTH-8

Expose user registration HTTP endpoint

- `api/handlers/users.go`
- `dto/users.go`

## Commit 9 — login endpoint (basic)

### AUTH-9

Implement password-based login endpoint

- `auth/login.go`
- returns success/failure only (no sessions yet)

---

# EPIC 2 — Sessions

Branch: `epic/2-sessions`

## Commit 10 — session model

### AUTH-10

Define session model and persistence layer

- `sessions/model.go`
- `/data/sessions`

## Commit 11 — session creation on login

### AUTH-11

Create session on successful login

- extend login service

## Commit 12 — auth middleware (session check)

### AUTH-12

Add session-based authentication middleware

- `api/middleware/auth.go`

## Commit 13 — logout endpoint

### AUTH-13

Implement session invalidation via logout endpoint

---

# EPIC 3 — JWT foundation

Branch: `epic/3-jwt`

## Commit 14 — signer interface

### AUTH-14

Introduce cryptographic signer abstraction

- `crypto/signer.go`

## Commit 15 — RSA implementation

### AUTH-15

Implement RSA signing and key management

- `crypto/rsa/*`

## Commit 16 — JWT encoder/decoder

### AUTH-16

Implement JWT encoding and decoding logic

- `crypto/jwt/*`

## Commit 17 — access token generation

### AUTH-17

Introduce OAuth access token generation using JWT

- `tokens/access.go`

## Commit 18 — migrate auth to JWT

### AUTH-18

Replace session authentication with JWT authentication

- middleware now validates tokens instead of sessions

---

# EPIC 4 — OIDC core

Branch: `epic/4-oidc`

## Commit 19 — discovery endpoint

### AUTH-19

Implement OpenID Connect discovery endpoint

## Commit 20 — JWKS endpoint

### AUTH-20

Expose JSON Web Key Set endpoint for public keys

## Commit 21 — ID token generation

### AUTH-21

Implement OpenID Connect ID token generation

- `tokens/id.go`

## Commit 22 — userinfo endpoint

### AUTH-22

Implement OIDC userinfo endpoint

---

# EPIC 5 — OAuth flow

Branch: `epic/5-oauth`

## Commit 23 — authorization endpoint

### AUTH-23

Implement OAuth2 authorization endpoint

## Commit 24 — auth code storage

### AUTH-24

Store and validate OAuth authorization codes

- filesystem or cache

## Commit 25 — token endpoint

### AUTH-25

Implement OAuth2 token exchange endpoint

## Commit 26 — PKCE support

### AUTH-26

Add PKCE verification for OAuth authorization flow

- `crypto/pkce.go`

---

# EPIC 6 — Google login

Branch: `epic/6-google`

## Commit 27 — provider interface

### AUTH-27

Define external OAuth provider interface

## Commit 28 — Google OAuth flow

### AUTH-28

Implement Google OAuth login integration

## Commit 29 — user linking

### AUTH-29

Link external provider identity to local user account

## Commit 30 — Google login endpoint

### AUTH-30

Expose Google login initiation endpoint

---

# EPIC 7 — Hardening

Branch: `epic/7-hardening`

## Commit 31 — Redis integration

### AUTH-31

Introduce Redis caching for sessions and OAuth state

## Commit 32 — rate limiting

### AUTH-32

Add request rate limiting middleware
## Commit 33 — key rotation

### AUTH-33

Implement RSA key rotation strategy

## Commit 34 — audit logging

### AUTH-34

Add security audit logging for authentication events

## Commit 35 — integration tests

### AUTH-35

Add end-to-end authentication flow tests

## Commit 36 — observability

### AUTH-36

Add metrics and structured observability hooks

...

...

...

# How to use this plan

Each commit should follow this rule:

> If the service breaks, the previous commit must still run.

That means:
- no giant “auth overhaul commits”
- no mixing JWT + OIDC + Google in one commit
- always build upward

## A practical way to execute it

If you want this to be extremely disciplined:

`git checkout -b epic/0-bootstrap`

For each commit:

```sh
git checkout -b auth-1
git commit -m "AUTH-1 Initialize Go module and base server"
```

Then merge into epic branch:

```sh
git checkout epic/0-bootstrap
git merge auth-1
```