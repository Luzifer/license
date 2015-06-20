package main

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
)

func getCmdBadge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "badge [license]",
		Short: "Gives you a markdown sniplet to show the license in your README.md",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cmd.Usage()
				fmt.Println("a license name is required")
				return
			}

			if license, ok := licenses[args[0]]; ok {
				badge := fmt.Sprintf("https://badge.luzifer.io/v1/badge?color=5d79b5&title=license&text=%s", url.QueryEscape(license.Name))
				fmt.Printf("[![License: %s](%s)](%s)\n", license.Name, badge, license.URL)
				return
			}

			fmt.Printf("License not found: %s", args[0])
		},
	}

	return cmd
}
