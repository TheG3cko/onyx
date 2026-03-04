package progressbar

import (
	"fmt"
	"time"

	"github.com/TheG3cko/onyx/pkg/helpers"
)

// Intermediate shows an indeterminate progress bar until Clear() is called.
func Intermediate() {
	if helpers.IsInteractive() {
		listenForInterrupt()
		fmt.Printf("\x1b]9;4;3;%d\a", 50)
	}
}

// IntermediateFor shows an indeterminate progress bar for the given duration, then clears it automatically.
// Must be called in a goroutine to be non-blocking.
func IntermediateFor(d time.Duration) {
	if helpers.IsInteractive() {
		listenForInterrupt()
		fmt.Printf("\x1b]9;4;3;%d\a", 50)
		time.Sleep(d)
		fmt.Printf("\x1b]9;4;0;%d\a", 0)
	}

}
