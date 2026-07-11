FROM golang:1.26 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /auth-service  ./cmd/auth-server

# FROM build-stage AS run-test-stage
# RUN go test -v ./...

FROM alpine AS runtime-stage

WORKDIR /

COPY --from=build-stage /auth-service /auth-service

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
    CMD wget http://0.0.0.0:8080/health || exit 1

RUN adduser -D appuser
USER appuser

CMD ["/auth-service"]