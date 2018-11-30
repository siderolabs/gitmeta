package cmd

import (
	"fmt"
	"os"

	"github.com/autonomy/gitmeta/internal/metadata"
	"github.com/spf13/cobra"
)

// builtCmd represents the built command
var builtCmd = &cobra.Command{
	Use:   "built",
	Short: "Prints a timestamp",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := metadata.NewMetadata()
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		fmt.Printf("%s", m.Built)
	},
}

func init() {
	rootCmd.AddCommand(builtCmd)
}