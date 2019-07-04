package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/talos-systems/gitmeta/pkg/git"
	"github.com/talos-systems/gitmeta/pkg/metadata"
)

var (
	negate bool
)

// pushableCmd represents the pushable command
var pushableCmd = &cobra.Command{
	Use:   "pushable",
	Short: "Prints a boolean value indicating if the image should be pushed",
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
		if m.Git.IsClean && m.Git.Branch == "master" {
			if negate {
				fmt.Printf("false")
			} else {
				fmt.Printf("true")
			}
		} else if m.Git.IsClean && m.Git.IsTag {
			if negate {
				fmt.Printf("false")
			} else {
				fmt.Printf("true")
			}
		} else {
			if negate {
				fmt.Printf("true")
			} else {
				fmt.Printf("false")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pushableCmd)
	pushableCmd.Flags().BoolVarP(&negate, "negate", "n", false, "negate the printed boolean")
}
