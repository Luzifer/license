package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func getGitConfigParameter(key string) string {
	out, err := exec.Command("git", "config", "--get", key).Output()
	if err != nil {
		fmt.Println(err)
	}

	return strings.TrimSpace(string(out))
}
