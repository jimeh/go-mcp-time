package server

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTimeServer(t *testing.T) {
	tests := []struct {
		name            string
		localTimezone   string
		expectError     bool
		expectedVersion string
	}{
		{
			name:            "valid timezone UTC",
			localTimezone:   "UTC",
			expectError:     false,
			expectedVersion: "1.0.0",
		},
		{
			name:            "valid timezone America/New_York",
			localTimezone:   "America/New_York",
			expectError:     false,
			expectedVersion: "1.0.0",
		},
		{
			name:            "empty timezone",
			localTimezone:   "",
			expectError:     false,
			expectedVersion: "1.0.0",
		},
		{
			name:            "valid timezone Europe/London",
			localTimezone:   "Europe/London",
			expectError:     false,
			expectedVersion: "1.0.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, err := NewTimeServer(tt.localTimezone)

			if tt.expectError {
				require.Error(t, err)
				assert.Nil(t, server)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, server)
			assert.NotNil(t, server.server)
		})
	}
}

func TestTimeServerRegisterTools(t *testing.T) {
	t.Run("register tools successfully", func(t *testing.T) {
		server, err := NewTimeServer("UTC")
		require.NoError(t, err)

		// Test that tools were registered by checking the server exists
		assert.NotNil(t, server.server)
	})

	t.Run("register tools with different timezone", func(t *testing.T) {
		server, err := NewTimeServer("America/New_York")
		require.NoError(t, err)

		// Test that tools were registered by checking the server exists
		assert.NotNil(t, server.server)
	})

	t.Run("register tools with empty timezone", func(t *testing.T) {
		server, err := NewTimeServer("")
		require.NoError(t, err)

		// Test that tools were registered by checking the server exists
		assert.NotNil(t, server.server)
	})
}

func TestTimeServerServe(t *testing.T) {
	t.Run("serve returns when context is canceled", func(t *testing.T) {
		server, err := NewTimeServer("UTC")
		require.NoError(t, err)

		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		// Note: In a real implementation, Serve might return an error
		// when context is canceled
		// For now, we're just testing that it doesn't panic and returns
		_ = server.Serve(ctx)
		// The actual behavior depends on the underlying server implementation
		// We're just verifying it doesn't panic
	})
}
