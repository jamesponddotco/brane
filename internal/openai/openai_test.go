package openai_test

import (
	"bytes"
	"context"
	"os"
	"testing"

	"git.sr.ht/~jamesponddotco/brane/internal/openai"
	"github.com/urfave/cli/v2"
)

func TestClient_Request(t *testing.T) {
	t.Parallel()

	key, found := os.LookupEnv("BRANE_KEY")
	if !found || key == "" {
		t.Skip("missing required environment variable: BRANE_KEY")
	}

	var (
		ctx    = context.Background()
		client = openai.NewClient(key)
	)

	tests := []struct {
		name      string
		model     string
		prompt    string
		wantError bool
	}{
		{
			name:      "Valid request with gpt-3.5-turbo",
			model:     "gpt-3.5-turbo",
			prompt:    "Hello, Brane!",
			wantError: false,
		},
		{
			name:      "Error request with non-existent model",
			model:     "non-existent-model",
			prompt:    "This should fail",
			wantError: true,
		},
		{
			name:      "Error request with empty model",
			model:     "",
			prompt:    "This should also fail",
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := client.Request(ctx, tt.model, "en", tt.prompt)
			if (err != nil) != tt.wantError {
				t.Errorf("Request() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestClient_Response(t *testing.T) {
	t.Parallel()

	key, found := os.LookupEnv("BRANE_KEY")
	if !found || key == "" {
		t.Skip("missing required environment variable: BRANE_KEY")
	}

	var (
		ctx    = context.Background()
		client = openai.NewClient(key)
	)

	stream, err := client.Request(ctx, "gpt-3.5-turbo", "en", "Hello, Brane!")
	if err != nil {
		t.Fatalf("Request() error = %v", err)
	}

	// Redirect standard output to capture the response.
	var buf bytes.Buffer
	app := &cli.App{
		Writer: &buf,
	}

	if err := client.Response(cli.NewContext(app, nil, nil), stream); err != nil {
		t.Fatalf("Response() error = %v", err)
	}

	if buf.String() == "" {
		t.Fatal("Expectd non-empty response")
	}
}
