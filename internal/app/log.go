package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"github.com/urfave/cli/v2"
)

const (
	// ErrEmptyThought is returned when no thought is provided by the user.
	ErrEmptyThought xerrors.Error = "missing thought to log"

	// ErrEmptyDirectory is returned when no directory is provided by the user.
	ErrEmptyDirectory xerrors.Error = "missing data directory"
)

// FileExtension is the file extension for note files.
const FileExtension string = ".md"

// LogAction is the action to perform when the log command is invoked.
func LogAction(ctx *cli.Context) error {
	if ctx.Args().Len() < 1 {
		return ErrEmptyThought
	}

	var (
		thought   = fmt.Sprintf("- %s\n", ctx.Args().Get(0))
		filename  = time.Now().Format("January 2, 2006") + FileExtension
		directory = ctx.String("directory")
	)

	if directory == "" {
		return ErrEmptyDirectory
	}

	path := filepath.Join(directory, filename)

	content, err := os.ReadFile(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to read notes file: %w", err)
	}

	if len(content) == 0 {
		thought = fmt.Sprintf("# %s\n\n%s", time.Now().Format("January 2, 2006"), thought)
	}

	if err := os.MkdirAll(directory, 0o755); err != nil {
		return fmt.Errorf("cannot create data directory: %w", err)
	}

	if err := os.WriteFile(path, append(content, []byte(thought)...), 0o600); err != nil {
		return fmt.Errorf("failed to write notes file: %w", err)
	}

	return nil
}
