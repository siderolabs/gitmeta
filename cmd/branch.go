package cmd

import (
	"fmt"
	"os"

	"github.com/talos-systems/gitmeta/pkg/git"
	"github.com/talos-systems/gitmeta/pkg/metadata"
	"github.com/spf13/cobra"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Prints the git branch",
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
		fmt.Printf("%s", m.Git.Branch)
	},
}
