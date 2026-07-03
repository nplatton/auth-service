## internal/providers

`provider.go`

Defines the common interface implemented by external identity providers.

### `google`

`callback.go`

Handles Google's OAuth callback.

`login.go`

Initiates the Google login flow.

`provider.go`

Implements Google-specific OAuth operations.

`mapper.go`

Maps Google user information to the local user model.