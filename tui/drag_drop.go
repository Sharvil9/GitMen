package tui

import (
	"github.com/charmbracelet/bubbletea"
)

type DragDropModel struct {
	Files []string // List of files dragged into the terminal.
}

func (m DragDropModel) Init() bubbletea.Cmd {
	return nil
}

func (m DragDropModel) Update(msg bubbletea.Msg) (bubbletea.Model, bubbletea.Cmd) {
	switch msg := msg.(type) {
	case bubbletea.MouseMsg:
		// Handle drag-and-drop logic (mocked for now).
		m.Files = append(m.Files, "dragged-file.txt")
	}
	return m, nil
}

func (m DragDropModel) View() string {
	return "Drag-and-Drop Mode: Drop files here."
}