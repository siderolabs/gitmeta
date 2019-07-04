package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/talos-systems/gitmeta/pkg/git"
	"github.com/talos-systems/gitmeta/pkg/metadata"
)

// gitTagCmd represents the tag command
var gitTagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Prints the git tag",
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
		fmt.Printf("%s", m.Git.Tag)
	},
}

// containerTagCmd represents the tag command
var containerTagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Prints the container image tag",
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
		fmt.Printf("%s", m.Container.Image.Tag)
	},
}
