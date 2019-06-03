package cmd

import (
	"fmt"
	"os"

	"github.com/talos-systems/gitmeta/pkg/metadata"
	"github.com/spf13/cobra"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Prints the git branch",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := metadata.NewMetadata()
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		fmt.Printf("%s", m.Git.Branch)
	},
}
