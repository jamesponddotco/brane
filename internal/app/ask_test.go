package app_test

import (
	"os"
	"path/filepath"
	"testing"

	"git.sr.ht/~jamesponddotco/brane/internal/app"
)

func TestNotesContent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		setup   func(dir string)
		want    string
		wantErr bool
	}{
		{
			name:    "no notes",
			want:    "",
			wantErr: false,
		},
		{
			name: "single note",
			setup: func(dir string) {
				if err := os.WriteFile(filepath.Join(dir, "test.md"), []byte("Test note."), 0o600); err != nil {
					t.Fatal(err)
				}
			},
			want:    "Test note.\n\n",
			wantErr: false,
		},
		{
			name: "multiple notes",
			setup: func(dir string) {
				if err := os.WriteFile(filepath.Join(dir, "test1.md"), []byte("First test note."), 0o600); err != nil {
					t.Fatal(err)
				}

				if err := os.WriteFile(filepath.Join(dir, "test2.md"), []byte("Second test note."), 0o600); err != nil {
					t.Fatal(err)
				}
			},
			want:    "First test note.\n\nSecond test note.\n\n",
			wantErr: false,
		},
		{
			name: "restricted directory permissions",
			setup: func(dir string) {
				if err := os.Chmod(dir, 0o000); err != nil {
					t.Fatal(err)
				}
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "unreadable note file",
			setup: func(dir string) {
				noteFile := filepath.Join(dir, "unreadable.md")

				if err := os.WriteFile(noteFile, []byte("This file is unreadable."), 0o600); err != nil {
					t.Fatal(err)
				}

				if err := os.Chmod(noteFile, 0o000); err != nil {
					t.Fatal(err)
				}
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tmpDir := t.TempDir()
			if tt.setup != nil {
				tt.setup(tmpDir)
			}

			got, err := app.NotesContent(tmpDir)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NotesContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Fatalf("NotesContent() = %s, want %s", got, tt.want)
			}
		})
	}
}
