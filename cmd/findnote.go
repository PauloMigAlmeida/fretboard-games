package cmd

import (
	"fmt"
	"github.com/PauloMigAlmeida/fretboard-games/game"
	"github.com/PauloMigAlmeida/fretboard-games/instrument"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var findnoteCmd = &cobra.Command{
	Use:   "findnote",
	Short: "Interactive fretboard training game to find specific notes on guitar strings",
	Long: `The FindNote game is an interactive fretboard training tool designed to help guitarists 
improve their note recognition and fretboard knowledge.

HOW IT WORKS:
The game presents you with specific notes (like A#, D, F#) and asks you to identify the 
fret numbers where those notes can be found on specific guitar strings.

GAME FLOW:
1. Configure the game by specifying:
   - How many different notes you want to find (e.g., 2-3 notes)
   - How many strings to search across (e.g., strings 1, 3, and 5)

2. The game displays a challenge like:
   "Find note(s) [ A# D ] across string(s) [ 1 3 5 ]"

3. For each string, you enter the fret numbers where the target notes appear:
   - Enter answers as comma-separated numbers (e.g., "3, 15" for frets 3 and 15)
   - The game validates your answers and provides immediate feedback

4. If incorrect, the game shows a fretboard visualization highlighting the correct positions

5. Track your progress with built-in statistics showing correct/incorrect answers
`,
	Run: func(cmd *cobra.Command, args []string) {
		fretboard := instrument.NewFretboard(24, instrument.StandardTuning())
		game := game.NewFindNoteGame(fretboard, os.Stdin, os.Stdout, game.NoSeed)

		err := game.Configure()
		if err != nil {
			fmt.Println("Error configuring the game:", err)
			os.Exit(-1)
		}

		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT)

		for {
			select {
			case _ = <-done:
				fmt.Println("SIGINT received. Existing the application...")
				game.Quit()
				return
			default:
				err = game.RunStep()
				if err != nil {
					fmt.Println("Error running game step:", err)
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(findnoteCmd)
}
