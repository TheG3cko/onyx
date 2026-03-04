package main

import (
	"fmt"
	"time"

	"github.com/TheG3cko/onyx/pkg/progressbar"
)

func main() {
	bar := progressbar.New(10)
	fmt.Printf("Init ProgressBar with %d steps\n\n", 10)
	var counter int

	for range 10 {
		bar.Update(1)
		counter++
		fmt.Printf("\033[KUpdate ProgressBar with %d steps until done\r", counter)
		time.Sleep(500 * time.Millisecond)
	}
	counter = 10
	fmt.Printf("\u001B[KFinished ProgressBar with %d steps\r", counter)
	fmt.Printf("\n\nIntermediateFor Progressbar for 3 seconds")
	go progressbar.IntermediateFor(3 * time.Second)
	time.Sleep(4 * time.Second)
	progressbar.Clear()
	fmt.Printf("\u001B[K\nIntermediate Progressbar for 2 seconds")
	progressbar.Intermediate()
	time.Sleep(2 * time.Second)
	progressbar.Clear()
}
