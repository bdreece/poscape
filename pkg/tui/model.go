package tui

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type ViewModel interface {
    tea.Model

    Keys() help.KeyMap
}
