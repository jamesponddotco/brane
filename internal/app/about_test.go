package app_test

import (
	"flag"
	"os"
	"testing"

	"git.sr.ht/~jamesponddotco/brane/internal/app"
	"github.com/urfave/cli/v2"
)

func TestAboutAction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		args     []string
		setup    func(dir string)
		wantErr  bool
		wantFile string
	}{
		{
			name:    "ErrEmptyInput",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "ErrEmptyDirectory",
			args:    []string{"info"},
			wantErr: true,
		},
		{
			name:     "SuccessfulWrite",
			args:     []string{"info"},
			wantFile: "# About me\n\n- info\n",
		},
		{
			name: "AppendInfo",
			args: []string{"more info"},
			setup: func(dir string) {
				initialContent := "# About me\n\n- info\n"

				if err := os.WriteFile(dir+"/"+app.Filename, []byte(initialContent), 0o600); err != nil {
					t.Fatal(err)
				}
			},
			wantFile: "# About me\n\n- info\n- more info\n",
		},
		{
			name: "ErrorReadingFile",
			args: []string{"info"},
			setup: func(dir string) {
				os.Mkdir(dir+"/"+app.Filename, 0o755) // Create directory with the name "About.md"
			},
			wantErr: true,
		},
		{
			name: "ErrorWritingFile",
			args: []string{"info"},
			setup: func(dir string) {
				os.WriteFile(dir+"/"+app.Filename, []byte{}, 0o400) // Create a read-only file
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			directory := t.TempDir()

			if tt.setup != nil {
				tt.setup(directory)
			}

			set := flag.NewFlagSet("test", 0)
			if tt.name != "ErrEmptyDirectory" {
				set.String("directory", directory, "doc")
			}

			set.Parse(tt.args)

			var (
				ctx = cli.NewContext(nil, set, nil)
				err = app.AboutAction(ctx)
			)

			if tt.wantErr && err == nil {
				t.Fatal("expected an error, got none")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("did not expect an error, but got: %v", err)
			}

			if tt.wantFile != "" {
				content, _ := os.ReadFile(directory + "/" + app.Filename)
				if string(content) != tt.wantFile {
					t.Fatalf("expected content %v, got %v", tt.wantFile, string(content))
				}
			}
		})
	}
}
