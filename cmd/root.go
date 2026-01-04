package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fretboard-games",
	Short: "Interactive guitar fretboard training games",
	Long: `Fretboard Games is a collection of interactive training tools designed to help 
guitarists improve their fretboard knowledge and note recognition skills.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
