package cmd

import (
	"fmt"
	"os"

	"github.com/autonomy/gitmeta/internal/metadata"
	"github.com/spf13/cobra"
)

// gitTagCmd represents the tag command
var gitTagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Prints the git tag",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := metadata.NewMetadata()
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
		m, err := metadata.NewMetadata()
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		fmt.Printf("%s", m.Container.Image.Tag)
	},
}
