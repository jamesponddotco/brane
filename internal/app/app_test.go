package app_test

import (
	"os"
	"testing"

	"git.sr.ht/~jamesponddotco/brane/internal/app"
)

func TestRun(t *testing.T) { //nolint:paralleltest // t.Setenv is not thread-safe
	key, found := os.LookupEnv("BRANE_KEY")
	if !found || key == "" {
		t.Skip("missing required environment variable: BRANE_KEY")
	}

	tests := []struct {
		name string
		args []string
		env  map[string]string
		want int
	}{
		{
			name: "no args",
			args: []string{"app"},
			want: 0,
		},
		{
			name: "log without thought",
			args: []string{"app", "log"},
			want: 1,
		},
		{
			name: "log with thought",
			args: []string{"app", "log", "Test thought"},
			env:  map[string]string{"BRANE_DIRECTORY": "/tmp"},
			want: 0,
		},
		{
			name: "log without a directory",
			args: []string{"app", "log", "Test thought"},
			env:  map[string]string{"BRANE_DIRECTORY": ""},
			want: 1,
		},
		{
			name: "ask with question",
			args: []string{"app", "ask", "What is Go?"},
			env:  map[string]string{"BRANE_DIRECTORY": "/tmp", "BRANE_KEY": key},
			want: 0,
		},
		{
			name: "ask without question",
			args: []string{"app", "ask"},
			want: 1,
		},
		{
			name: "ask without a directory",
			args: []string{"app", "ask", "What is Go?"},
			env:  map[string]string{"BRANE_DIRECTORY": ""},
			want: 1,
		},
		{
			name: "ask with question but missing key",
			args: []string{"app", "ask", "What is Go?"},
			env:  map[string]string{"BRANE_DIRECTORY": "/tmp", "BRANE_KEY": ""},
			want: 1,
		},
	}

	for _, tt := range tests { //nolint:paralleltest // t.Setenv is not thread-safe
		t.Run(tt.name, func(t *testing.T) {
			for key, val := range tt.env {
				t.Setenv(key, val)
			}

			got := app.Run(tt.args)
			if got != tt.want {
				t.Fatalf("Run() = %d, want %d", got, tt.want)
			}
		})
	}
}
