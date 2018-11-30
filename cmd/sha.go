package cmd

import (
	"fmt"
	"os"

	"github.com/autonomy/gitmeta/internal/metadata"
	"github.com/spf13/cobra"
)

// shaCmd represents the sha command
var shaCmd = &cobra.Command{
	Use:   "sha",
	Short: "Prints the short git SHA",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := metadata.NewMetadata()
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		fmt.Printf("%s", m.Git.SHA)
	},
}
