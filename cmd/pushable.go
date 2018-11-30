package cmd

import (
	"fmt"
	"os"

	"github.com/autonomy/gitmeta/internal/metadata"
	"github.com/spf13/cobra"
)

// pushableCmd represents the pushable command
var pushableCmd = &cobra.Command{
	Use:   "pushable",
	Short: "Prints a boolean value indicating if the image should be pushed",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := metadata.NewMetadata()
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		if m.Git.IsClean && m.Git.Branch == "master" {
			fmt.Printf("true")
		} else if m.Git.IsClean && m.Git.IsTag {
			fmt.Printf("true")
		} else {
			fmt.Printf("false")
		}
	},
}

func init() {
	rootCmd.AddCommand(pushableCmd)
}
