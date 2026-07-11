# Auth Service <!-- omit from toc -->

A Go-based authentication service.

---

**Contents**
- [1. Getting Started](#1-getting-started)
  - [1.1. Local](#11-local)
  - [1.2. Docker](#12-docker)
- [2. Development](#2-development)
- [3. Configuration](#3-configuration)
- [4. Health Check](#4-health-check)
- [5. Logging](#5-logging)
- [6. Shutdown](#6-shutdown)

## 1. Getting Started

### 1.1. Local

Run the service directly:

```bash
go run ./cmd/auth-server/
```

By default the service loads:

```
configs/local.yaml
```

An alternative config file can be provided using `CONFIG_PATH`:

```bash
CONFIG_PATH=configs/custom.yaml go run ./cmd/auth-server/
```

The service starts on the configured host and port.

### 1.2. Docker

Build and start the service using Docker Compose:

```bash
docker compose up
```

To also recompile the image run:

```bash
docker compose up --build
```

The build can be stopped with:

```bash
docker compose down -v
```

The Docker deployment:
1. Builds the service using a multi-stage Docker build
2. Runs the compiled binary in a minimal runtime image
3. Runs as a non-root user
4. Mounts configuration externally
5. Writes logs to stdout/stderr

## 2. Development

Rebuild the Docker image after changing application code:

```bash
docker compose up --build
```

If only configuration changes are made, a restart is sufficient:

```bash
docker compose up
```

## 3. Configuration

Configuration is loaded from a YAML file.

Example:
```yaml
server:
  host: 0.0.0.0
  port: 8080

logging:
  level: debug
```

When running in Docker, the configuration file is mounted into the container and the path is provided through:

```bash
CONFIG_PATH=/configs/config.yaml
```

This allows configuration changes without rebuilding the image.

## 4. Health Check

The service exposes:

```
GET /health
```

Example:

```bash
curl http://localhost:8080/health
```

Example response:

```json
{
  "status": "healthy"
}
```

## 5. Logging

The service uses structured JSON logging.

Logs are written to stdout/stderr so they integrate with container logging systems.

View Docker logs:

```bash
docker compose logs
```

or:

```bash
docker logs <container>
```

## 6. Shutdown

The service supports graceful shutdown.
When receiving:

- `SIGINT` (Ctrl+C)
- `SIGTERM` (`docker stop`)

the service will:

1. stop accepting new requests
2. wait for in-flight requests to complete
3. shut down cleanly