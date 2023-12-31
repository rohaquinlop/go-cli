package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go-cli",
	Long:  `All software has versions. This is go-cli's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-cli v0.1 -- HEAD")
	},
}
