package tests

import (
    "testing"
    "gitmen/core"
)

func TestGitCommand(t *testing.T) {
    _, err := core.GitCommand("status")
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
}