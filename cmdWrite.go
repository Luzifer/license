package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var (
	force bool
)

func getCmdWrite() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "write [license]",
		Short: "Write the pre-filled license template to ./LICENSE",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cmd.Usage()
				fmt.Println("a license name is required")
				return
			}

			licenseFile := "./LICENSE"

			if _, err := os.Stat(licenseFile); err == nil && !force {
				fmt.Printf("LICENSE file already exists. Exiting.\n")
				os.Exit(1)
			}

			lic, err := preFillLicense(args[0])
			if err != nil {
				fmt.Printf("License '%s' is currently not included.\n", args[0])
				os.Exit(1)
			}

			ioutil.WriteFile(licenseFile, lic, 0644)
			fmt.Println("License written.")
		},
	}

	cmd.Flags().BoolVarP(&force, "force", "f", false, "Force overwrite the license file")

	return cmd
}
