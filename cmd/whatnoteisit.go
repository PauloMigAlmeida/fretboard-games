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

var whatnoteisitCmd = &cobra.Command{
	Use:   "whatnoteisit",
	Short: "Interactive fretboard training game to identify notes at given fret positions",
	Long: `The WhatNoteIsIt game is an interactive fretboard training tool that helps guitarists
improve their note recognition by position.

HOW IT WORKS:
The game presents a specific fret and string (e.g., "fret 5 of string 2") and asks you
to name the note found at that position.

ANSWERING:
- Enter the note name, optionally followed by an accidental (e.g., "E", "F#", "Bb")
- Enharmonic equivalents are accepted (e.g., "D#" and "Eb" are both correct)

Track your progress with built-in statistics shown when you quit (Ctrl+C).
`,
	Run: func(cmd *cobra.Command, args []string) {
		fretboard := instrument.NewFretboard(24, instrument.StandardTuning())
		g := game.NewWhatNoteIsItGame(fretboard, os.Stdin, os.Stdout, game.NoSeed)

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
	rootCmd.AddCommand(whatnoteisitCmd)
}
