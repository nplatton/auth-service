## internal/crypto

`argon.go`

Hashes and verifies user passwords using the Argon2id password hashing algorithm.

`pkce.go`

Generates and validates PKCE code verifiers and code challenges for OAuth Authorization Code flows.

`random.go`

Provides cryptographically secure random value generation for IDs, tokens, nonces and secrets.

`signer.go`

Defines the common signing interface implemented by all supported cryptographic algorithms.

### `jwt`

`encoder.go`

Encodes JWT headers and claims into the standard compact JWT format.

`decoder.go`

Parses and decodes JWTs into their header, payload and signature components.

`verifier.go`

Validates JWT structure, claims and signatures using the configured cryptographic signer.

### `rsa`

`keys.go`

Loads, generates and manages RSA public/private key pairs used for digital signatures.

`signer.go`

Implements the common signing interface using the RSA signature algorithm.