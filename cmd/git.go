package cmd

import (
	"github.com/spf13/cobra"
)

// gitCmd represents the git command
var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Retreive git metadata",
	Long:  ``,
}

func init() {
	rootCmd.AddCommand(gitCmd)
	gitCmd.AddCommand(shaCmd, branchCmd, gitTagCmd)
}
