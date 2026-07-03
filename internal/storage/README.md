internal/storage

### `filesystem`

`reader.go`

Reads persisted objects from disk.

`writer.go`

Writes objects to disk.

`atomic.go`

Performs atomic file writes to prevent corruption.

`locking.go`

Provides file locking to avoid concurrent write conflicts.

### `indexes`

`email.go`

Maintains the email-to-user lookup index.

`session.go`

Maintains the session lookup index.

`client.go`

Maintains the OAuth client lookup index.