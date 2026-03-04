package progressbar

import (
	"github.com/TheG3cko/onyx/pkg/helpers"
)

type Percentage struct {
	Current     int
	Steps       int
	Debug       bool
	Interactive bool
}

// New creates a new progress bar with the given total number of steps.
func New(steps int) *Percentage {
	percentage := &Percentage{Steps: steps, Current: 0}
	if helpers.IsInteractive() {
		percentage.Interactive = true
		listenForInterrupt()
	}
	return percentage
}
