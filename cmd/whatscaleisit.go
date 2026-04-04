package cmd

import (
	"fmt"
	"github.com/PauloMigAlmeida/fretboard-games/game"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var whatscaleisitCmd = &cobra.Command{
	Use:   "whatscaleisit",
	Short: "Interactive scale training game to identify notes in a given scale",
	Long: `The WhatScaleIsIt game is an interactive music theory training tool that helps
musicians learn the notes that make up common scales.

HOW IT WORKS:
The game presents a root note and scale type (e.g., "C Minor") and asks you
to identify the correct set of notes from three multiple-choice options.

SCALES COVERED:
- Major
- Minor (Natural Minor)
- Melodic Minor
- Harmonic Minor

When you answer incorrectly, the game shows the correct notes and the
interval formula for that scale (e.g., 1,2,b3,4,5,b6,b7 for Minor).

Track your progress with built-in statistics shown when you quit (Ctrl+C).
`,
	Run: func(cmd *cobra.Command, args []string) {
		g := game.NewWhatScaleIsItGame(os.Stdin, os.Stdout, game.NoSeed)

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
	rootCmd.AddCommand(whatscaleisitCmd)
}
