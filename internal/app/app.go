// Package app is the main package for the application.
package app

import (
	"fmt"
	"os"

	"git.sr.ht/~jamesponddotco/brane/internal/meta"
	"git.sr.ht/~jamesponddotco/brane/internal/osutil"
	"github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
)

// Run is the entry point for the application.
func Run(args []string) int {
	app := cli.NewApp()
	app.Name = meta.Name
	app.Version = meta.Version
	app.Usage = meta.Description
	app.HideHelpCommand = true

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "directory",
			Aliases: []string{"d"},
			Usage:   "the directory where note files are stored",
			Value:   osutil.UserDataDir(),
			EnvVars: []string{"BRANE_DIRECTORY"},
		},
		&cli.StringFlag{
			Name:    "model",
			Aliases: []string{"m"},
			Usage:   "the OpenAI model to use",
			Value:   openai.GPT3Dot5Turbo16K,
			EnvVars: []string{"BRANE_MODEL"},
		},
		&cli.StringFlag{
			Name:    "key",
			Aliases: []string{"k"},
			Usage:   "the OpenAI API key to use",
			EnvVars: []string{"BRANE_KEY"},
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "log",
			Aliases: []string{"l"},
			Usage:   "log a new thought",
			Action:  LogAction,
		},
		{
			Name:    "ask",
			Aliases: []string{"a"},
			Usage:   "ask the AI questions about your thoughts",
			Action:  AskAction,
		},
	}

	if err := app.Run(args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)

		return 1
	}

	return 0
}
