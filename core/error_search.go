package core

import (
	"fmt"
	"os/exec"
	"strings"
)

// SearchErrorOnGoogle opens a browser to search for the error on Google.
func SearchErrorOnGoogle(errorMsg string) error {
	query := strings.ReplaceAll(errorMsg, " ", "+")
	url := fmt.Sprintf("https://www.google.com/search?q=%s", query)
	cmd := exec.Command("xdg-open", url) // Use xdg-open for Linux; modify for macOS/Windows.
	return cmd.Run()
}