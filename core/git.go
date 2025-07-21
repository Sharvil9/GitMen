package core

import (
    "os/exec"
    "strings"
)

func GitCommand(args ...string) (string, error) {
    cmd := exec.Command("git", args...)
    output, err := cmd.CombinedOutput()
    return strings.TrimSpace(string(output)), err
}