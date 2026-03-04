package helpers

import (
	"os"

	"golang.org/x/term"
)

// IsInteractive reports whether stdout is an interactive terminal.
func IsInteractive() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}
