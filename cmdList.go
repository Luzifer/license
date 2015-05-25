package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

func getCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists all available licenses",
		Run: func(cmd *cobra.Command, args []string) {
			l := []string{}
			maxlen := 0

			for n := range licenses {
				if len(n) > maxlen {
					maxlen = len(n)
				}
			}

			linefmt := fmt.Sprintf("%%-%ds [%%s]", maxlen+2)

			for name, lic := range licenses {
				l = append(l, fmt.Sprintf(linefmt, name, lic.Description))
			}

			sort.Strings(l)
			fmt.Println(strings.Join(l, "\n"))
		},
	}

	return cmd
}
