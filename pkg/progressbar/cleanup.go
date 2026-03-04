package progressbar

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/TheG3cko/onyx/pkg/helpers"
)

var interruptListening bool

// Clear removes the progress bar from the terminal.
func Clear() {
	if helpers.IsInteractive() {
		fmt.Printf("\x1b]9;4;0;%d\a", 0)
	}
}

func listenForInterrupt() {
	if helpers.IsInteractive() && !interruptListening {
		interruptListening = true
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		go func() {
			<-c
			Clear()
			os.Exit(0)
		}()
	}
}
