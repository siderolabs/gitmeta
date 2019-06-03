package cmd

import (
	"fmt"
	"os"

	"github.com/talos-systems/gitmeta/pkg/git"
	"github.com/talos-systems/gitmeta/pkg/metadata"
	"github.com/spf13/cobra"
)

// shaCmd represents the sha command
var shaCmd = &cobra.Command{
	Use:   "sha",
	Short: "Prints the short git SHA",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		git, err := git.NewGit()
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		m, err := metadata.NewMetadata(git)
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		fmt.Printf("%s", m.Git.SHA)
	},
}
