package progressbar

import (
	"fmt"
	"os"
)

// Update advances the progress bar by the given number of steps.
func (p *Percentage) Update(steps int) {
	if p.Interactive {
		p.Current += steps

		if p.Current >= p.Steps {
			p.Current = p.Steps
		}

		fmt.Printf("\x1b]9;4;1;%.0f\a", float64(p.Current)/float64(p.Steps)*100)
		if p.Debug {
			fmt.Fprintf(os.Stderr, "%+v", p)
		}
	}

}
