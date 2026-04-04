package cmd

import (
	"fmt"
	"github.com/PauloMigAlmeida/fretboard-games/game"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var namethescaleCmd = &cobra.Command{
	Use:   "namethescale",
	Short: "Interactive scale training game to identify a scale by its notes",
	Long: `The NameTheScale game is an interactive music theory training tool that helps
musicians identify scale names from a given set of notes.

HOW IT WORKS:
The game presents the notes of a randomly chosen scale (e.g., "A - C - D - E - G") and asks
you to identify the correct scale name from three multiple-choice options.

SCALES COVERED:
- Major
- Minor (Natural Minor)
- Melodic Minor
- Harmonic Minor
- Major Pentatonic
- Minor Pentatonic
- Blues
- Spanish
- Persian

When you answer incorrectly, the game shows the correct scale name and its interval formula.

Track your progress with built-in statistics shown when you quit (Ctrl+C).
`,
	Run: func(cmd *cobra.Command, args []string) {
		g := game.NewNameTheScaleGame(os.Stdin, os.Stdout, game.NoSeed)

		err := g.Configure()
		if err != nil {
			fmt.Println("Error configuring the game:", err)
			os.Exit(-1)
		}

		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT)

		for {
			select {
			case <-done:
				fmt.Println("SIGINT received. Exiting the application...")
				g.Quit()
				return
			default:
				err = g.RunStep()
				if err != nil {
					fmt.Println("Error running game step:", err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(namethescaleCmd)
}
