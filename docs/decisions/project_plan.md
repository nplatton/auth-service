# Auth Service Development Plan

# 1. Build philosophy

For an auth / OIDC system, the biggest mistake is building “features” in isolation.

Instead, you want a vertical progression:

> **Start with the smallest end-to-end working authentication flow, then expand capabilities around it.**

That means:

- Don’t start with OIDC full spec compliance
- Don’t start with Google login
- Don’t start with refresh tokens
- Don’t start with Redis

Instead:

> **Get: login → session → protected route working first**

Then evolve it into OAuth/OIDC.

---

# 2. Development phases overview

## Phase 0 — Infrastructure & skeleton  
Goal: service runs, responds, logs, and is testable.

## Phase 1 — Identity core (users + passwords)  
Goal: you can register/login locally.

## Phase 2 — Sessions + auth middleware  
Goal: you can protect routes with a session.

## Phase 3 — Token system (JWT foundation)  
Goal: replace session auth with token-based auth.

## Phase 4 — OIDC minimal provider  
Goal: external client can authenticate.

## Phase 5 — OAuth flows + PKCE  
Goal: real OAuth authorization code flow works.

## Phase 6 — External providers (Google login)  
Goal: “Login with Google”.

## Phase 7 — Hardening  
Goal: security, rate limiting, key rotation, observability.

---

# 3. Execution plan (Jira-style)

## 🟦 EPIC 0 — Project bootstrap

### AUTH-1: Initialise Go service skeleton
- Create `cmd/auth-server/main.go`
- Create HTTP server (net/http or chi/gin)
- Health endpoint `/health`

### AUTH-2: Add configuration system
- Implement `internal/config`
- Load YAML/env
- Validate startup config

### AUTH-3: Add structured logging
- Implement `internal/logging`
- Add request logging middleware stub

### AUTH-4: Dockerise service
- Dockerfile
- docker-compose with service only

---

## 🟩 EPIC 1 — User identity core

### AUTH-5: Implement user model
- `internal/users/model.go`

### AUTH-6: Implement filesystem user repository
- `internal/storage/filesystem`
- persist/load user JSON

### AUTH-7: User registration service
- `internal/users/service.go`
- password hashing via `crypto/argon.go`

### AUTH-8: Register endpoint
- `internal/api/handlers/users.go`
- DTOs for register request

### AUTH-9: Login endpoint
- `internal/auth/login.go`
- basic password check

---

## 🟨 EPIC 2 — Sessions

### AUTH-10: Session model + persistence
- `internal/sessions/model.go`
- stored in `/data/sessions`

### AUTH-11: Session creation on login
- extend login flow

### AUTH-12: Session validation middleware
- `internal/api/middleware/auth.go`

### AUTH-13: Logout endpoint
- invalidate session

---

## 🟧 EPIC 3 — JWT foundation

### 🟢 AUTH-14: Add crypto/signer interface
- (already exists conceptually — now implement)

### AUTH-15: Implement RSA signing
- `crypto/rsa`

### AUTH-16: Implement JWT encoder/decoder
- `crypto/jwt`

### AUTH-17: Implement Access Token generation
- `tokens/access.go`

### AUTH-18: Replace session auth with JWT auth
- middleware now validates JWT instead of session

---

## 🟪 EPIC 4 — OIDC core

### AUTH-19: Implement OIDC discovery endpoint
- `oidc/discovery.go`

### AUTH-20: Implement JWKS endpoint
- `oidc/jwks.go`

### AUTH-21: Implement ID token generation
- `tokens/id.go`

### AUTH-22: Implement `/userinfo`
- `oidc/userinfo.go`

---

## 🟥 EPIC 5 — OAuth Authorization Code flow

### AUTH-23: Implement /authorize endpoint
- `oidc/authorize.go`
- generates auth code

### AUTH-24: Implement authorization code storage
- stored in `cache/state.go` or filesystem

### AUTH-25: Implement /token endpoint
- `oidc/token.go`

### AUTH-26: Implement PKCE validation
- `crypto/pkce.go`

---

## 🟫 EPIC 6 — External providers (Google OAuth)

### 🟢 AUTH-27: Create provider interface
- `providers/provider.go`

### AUTH-28: Implement Google OAuth flow
- `providers/google/*`

### AUTH-29: Link external identity to local user
- mapping logic

### AUTH-30: Add “Login with Google” endpoint
- integrated into `/authorize`

---

## ⬛ EPIC 7 — Hardening & production readiness

### AUTH-31: Add Redis caching layer
- sessions, PKCE, rate limits

### AUTH-32: Rate limiting middleware
- `api/middleware/rate_limit.go`

### AUTH-33: Key rotation support
- RSA key rotation strategy

### AUTH-34: Audit logging
- security events

### AUTH-35: Add integration test suite
- full login → token → userinfo flow

### AUTH-36: Add observability hooks
- metrics (Prometheus-style)

---

# 4. What NOT to build early (important)

Avoid these until Phase 5–6:

### ❌ Google login
Too early, depends on full OAuth flow.

### ❌ Redis
Start with filesystem first; swap later.

### ❌ Key rotation
Adds unnecessary complexity before core works.

### ❌ Full OIDC compliance
Start minimal, then expand.

### ❌ Micro-optimised crypto abstraction
You only need one signer initially (RSA).

---

# 5. Recommended build strategy summary

If I compress everything:

> Build **vertical slices**, not layers.

Each milestone should produce:

- a running server
- a working endpoint
- a testable behaviour

Example progression:

1. Register user
2. Login user
3. Protect route
4. Issue JWT
5. Validate JWT
6. Add OIDC
7. Add OAuth
8. Add Google