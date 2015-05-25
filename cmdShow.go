package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func getCmdShow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show [license]",
		Short: "Prints the pre-filled license template to your terminal",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cmd.Usage()
				fmt.Println("a license name is required")
				return
			}

			lic, err := preFillLicense(args[0])
			if err != nil {
				fmt.Printf("License '%s' is currently not included.\n", args[0])
				os.Exit(1)
			}

			fmt.Println(string(lic))
		},
	}

	return cmd
}
