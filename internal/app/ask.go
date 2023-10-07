package app

import (
	"fmt"
	"os"
	"path/filepath"

	"git.sr.ht/~jamesponddotco/brane/internal/openai"
	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"git.sr.ht/~jamesponddotco/xstd-go/xunsafe"
	"github.com/urfave/cli/v2"
)

const (
	// ErrEmptyQuestion is returned when no question is provided by the user.
	ErrEmptyQuestion xerrors.Error = "missing question to ask"

	// ErrEmptyKey is returned when no API key for OpenAI is provided by the user.
	ErrEmptyKey xerrors.Error = "missing OpenAI API key"
)

// AskAction is the action to perform when the ask command is invoked.
func AskAction(ctx *cli.Context) error {
	if ctx.Args().Len() < 1 {
		return ErrEmptyQuestion
	}

	var (
		question  = ctx.Args().Get(0)
		directory = ctx.String("directory")
		key       = ctx.String("key")
		model     = ctx.String("model")
		language  = ctx.String("language")
	)

	if directory == "" {
		return ErrEmptyDirectory
	}

	if key == "" {
		return ErrEmptyKey
	}

	notes, err := NotesContent(directory)
	if err != nil {
		return fmt.Errorf("failed to read notes: %w", err)
	}

	var (
		client = openai.NewClient(key)
		prompt = question + "\n\n" + notes
	)

	req, err := client.Request(ctx.Context, model, language, prompt)
	if err != nil {
		return fmt.Errorf("failed to create request to OpenAI: %w", err)
	}

	if err := client.Response(ctx, req); err != nil {
		return fmt.Errorf("failed to get response from OpenAI: %w", err)
	}

	return nil
}

// NotesContent reads all notes files from the specified directory and return
// them as a single string to be used as context for OpenAI.
func NotesContent(directory string) (string, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	var content string

	for _, file := range files {
		if filepath.Ext(file.Name()) == FileExtension {
			data, err := os.ReadFile(filepath.Join(directory, file.Name()))
			if err != nil {
				return "", fmt.Errorf("%w", err)
			}

			content += xunsafe.BytesToString(data) + "\n\n"
		}
	}

	return content, nil
}
