package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dirsCmd)
}

var dirsCmd = &cobra.Command{
	Use:   "dirs [path of directory to print]",
	Short: "Print the directories of go-cli",
	Long:  `Print the directories of go-cli`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := errors.New("path of directory to print is required")
			fmt.Println(err)
			os.Exit(1)
		}

		for _, arg := range args {
			var path string
			if arg == "." {
				local_path, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				path = local_path
			} else {
				path = arg
			}

			files, err := ioutil.ReadDir(path)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			for _, file := range files {
				fmt.Println(file.Name())
			}
		}
	},
}
