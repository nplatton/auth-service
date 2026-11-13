package users

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestFindByID(t *testing.T) {
	createdAt := time.Now()
	m := initialiseTest(t, createdAt)

	cases := []struct {
		name      string
		id        string
		expected  User
		expectErr bool
	}{
		{
			name: "Valid user ID",
			id:   "123456789",
			expected: User{
				ID:        "123456789",
				Username:  "test-user-1",
				Email:     "test@example.com",
				CreatedAt: createdAt,
				UpdatedAt: time.Time{},
			},
		},
		{
			name:      "Invalid user ID",
			id:        "987654321",
			expectErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := m.FindByID(context.Background(), tc.id)
			if tc.expectErr {
				if !errors.Is(err, ErrUserNotFound) {
					t.Fatalf("expected %v, got %v", ErrUserNotFound, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestFindByUsername(t *testing.T) {
	createdAt := time.Now()
	m := initialiseTest(t, createdAt)

	cases := []struct {
		name        string
		username    string
		expected    User
		expectErr   bool
		expectedErr error
	}{
		{
			name:     "Valid username",
			username: "test-user-1",
			expected: User{
				ID:        "123456789",
				Username:  "test-user-1",
				Email:     "test@example.com",
				CreatedAt: createdAt,
				UpdatedAt: time.Time{},
			},
		},
		{
			name:        "Username not found",
			username:    "invalid",
			expectErr:   true,
			expectedErr: ErrUsernameNotFound,
		},
		{
			name:        "Username exists but no user",
			username:    "test-user-2",
			expectErr:   true,
			expectedErr: ErrUserNotFound,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := m.FindByUsername(context.Background(), tc.username)
			if tc.expectErr {
				if !errors.Is(err, tc.expectedErr) {
					t.Fatalf("expected %v, got %v", tc.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	createdAt := time.Now()
	m := initialiseTest(t, createdAt)

	cases := []struct {
		name      string
		user      User
		expected  User
		expectErr bool
	}{
		{
			name: "Create new user",
			user: User{
				Username: "new-user",
				Email:    "test@example.com",
			},
			expected: User{
				Username: "new-user",
			},
		},
		{
			name: "Username already exists",
			user: User{
				Username: "test-user-1",
				Email:    "test@example.com",
			},
			expectErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := m.Create(context.Background(), tc.user)
			if tc.expectErr {
				if !errors.Is(err, ErrUsernameExists) {
					t.Fatalf("expected %v, got %v", ErrUsernameExists, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			actual := result.Username

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	createdAt := time.Now()
	m := initialiseTest(t, createdAt)

	cases := []struct {
		name        string
		user        User
		expected    User
		expectErr   bool
		expectedErr error
	}{
		{
			name: "Valid user ID",
			user: User{
				ID:       "123456789",
				Username: "new-username",
				Email:    "test@example.com",
			},
			expected: User{
				ID:       "123456789",
				Username: "new-username",
			},
		},
		{
			name: "Invalid user ID",
			user: User{
				ID:       "987654321",
				Username: "new-username",
				Email:    "test@example.com",
			},
			expectErr:   true,
			expectedErr: ErrUserNotFound,
		},
		{
			name: "Missing user ID",
			user: User{
				Username: "new-username",
				Email:    "test@example.com",
			},
			expectErr:   true,
			expectedErr: ErrInvalidUserID,
		},
		{
			name: "New username already taken",
			user: User{
				ID:       "123456789",
				Username: "test-user-3",
				Email:    "test@example.com",
			},
			expectErr:   true,
			expectedErr: ErrUsernameExists,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := m.Update(context.Background(), tc.user)
			if tc.expectErr {
				if !errors.Is(err, tc.expectedErr) {
					t.Fatalf("expected %v, got %v", tc.expectErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			actual := result.Username

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	createdAt := time.Now()
	m := initialiseTest(t, createdAt)

	cases := []struct {
		name      string
		user      User
		expectErr bool
	}{
		{
			name: "Valid user ID",
			user: User{
				ID:       "123456789",
				Username: "test-user-1",
			},
			expectErr: false,
		},
		{
			name: "Invalid user ID",
			user: User{
				ID:       "987654321",
				Username: "does-not-exist",
			},
			expectErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := m.Delete(context.Background(), tc.user.ID)
			if tc.expectErr {
				if !errors.Is(err, ErrUserNotFound) {
					t.Fatalf("expected %v, got %v", ErrUserNotFound, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			user, err := m.FindByID(context.Background(), tc.user.ID)
			if !errors.Is(err, ErrUserNotFound) {
				t.Fatalf("expected %v, got user: %v", ErrUserNotFound, user)
			}

			user, err = m.FindByUsername(context.Background(), tc.user.Username)
			if !errors.Is(err, ErrUsernameNotFound) {
				if err != nil {
					t.Fatalf("expected %v, got %v", ErrUserNotFound, err)
				}
				t.Fatalf("expected %v, got user: %v", ErrUserNotFound, user)
			}
		})
	}
}

func initialiseTest(t *testing.T, createdAt time.Time) *MemoryRepository {
	t.Helper()

	memoryRepository := NewMemoryRepository()

	memoryRepository.users = map[string]User{
		"123456789": {
			ID:        "123456789",
			Username:  "test-user-1",
			Email:     "test@example.com",
			CreatedAt: createdAt,
			UpdatedAt: time.Time{},
		},
		"234567891": {
			ID:        "234567891",
			Username:  "test-user-3",
			Email:     "test@example.com",
			CreatedAt: createdAt,
			UpdatedAt: time.Time{},
		},
	}

	memoryRepository.usernames = map[string]string{
		"test-user-1": "123456789",
		"test-user-2": "invalid",
		"test-user-3": "234567891",
	}

	return memoryRepository
}
