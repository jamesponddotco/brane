package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"github.com/urfave/cli/v2"
)

const (
	// ErrEmptyInput is returned when the about command is invoked without an input.
	ErrEmptyInput xerrors.Error = "missing information to log"
)

// Filename is the filename for the about me file.
const Filename string = "About" + FileExtension

// AboutAction is the action to perform when the about command is invoked.
func AboutAction(ctx *cli.Context) error {
	if ctx.Args().Len() < 1 {
		return ErrEmptyInput
	}

	var (
		input     = fmt.Sprintf("- %s\n", ctx.Args().Get(0))
		directory = ctx.String("directory")
	)

	if directory == "" {
		return ErrEmptyDirectory
	}

	path := filepath.Join(directory, Filename)

	content, err := os.ReadFile(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to read about file: %w", err)
	}

	if len(content) == 0 {
		input = fmt.Sprintf("# About me\n\n%s", input)
	}

	if err := os.MkdirAll(directory, 0o755); err != nil {
		return fmt.Errorf("cannot create data directory: %w", err)
	}

	if err := os.WriteFile(path, append(content, []byte(input)...), 0o600); err != nil {
		return fmt.Errorf("failed to write about file: %w", err)
	}

	return nil
}
