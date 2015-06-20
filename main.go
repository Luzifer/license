package main

import (
	"bufio"
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/spf13/cobra"
)

const (
	version = "1.0.0"
)

var (
	licenses = map[string]license{}
)

type license struct {
	Description string
	Template    string
	URL         string
	Name        string
}

func addLicense(name string, l license) {
	licenses[name] = l
}

func preFillLicense(key string) ([]byte, error) {
	gitUserName := getGitConfigParameter("user.name")
	gitUserEmail := getGitConfigParameter("user.email")

	var buf bytes.Buffer
	f := bufio.NewWriter(&buf)

	if license, ok := licenses[key]; ok {
		tpl := template.New("license")
		tpl.Parse(license.Template)
		tpl.Execute(f, map[string]interface{}{
			"Year":   time.Now().Format("2006"),
			"Author": fmt.Sprintf("%s <%s>", gitUserName, gitUserEmail),
		})
		f.WriteString("\n")
		f.Flush()
		return buf.Bytes(), nil
	}
	return nil, fmt.Errorf("License not found: %s", key)

}

func main() {
	app := &cobra.Command{
		Use:   "license",
		Short: "license is a tool to quickly put your code under a license",
	}

	app.AddCommand(
		getCmdList(),
		getCmdShow(),
		getCmdWrite(),
		getCmdVersion(),
		getCmdBadge(),
	)

	app.Execute()
}
