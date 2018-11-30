package cmd

import (
	"github.com/spf13/cobra"
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Retreive image metadata",
	Long:  ``,
}

func init() {
	rootCmd.AddCommand(imageCmd)
	imageCmd.AddCommand(containerTagCmd)
}
