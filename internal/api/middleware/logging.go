package middleware

// Uses logging/logger.go to log HTTP requests:
// Its job is:
//
// HTTP Request
// ↓
// Request started
// ↓
// Request finished
// ↓
// Status code
// ↓
// Duration
//
// Example log:
// POST /login
// Status: 200
// Duration: 18ms
