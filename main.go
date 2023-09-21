package main

import (
	"os"

	"git.sr.ht/~jamesponddotco/brane/internal/app"
)

func main() {
	os.Exit(app.Run(os.Args))
}
