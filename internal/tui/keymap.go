package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
)

type KeyMap struct {
	textarea.KeyMap

	Blur  key.Binding
	Print key.Binding
	Help  key.Binding
	Quit  key.Binding
}

// FullHelp implements help.KeyMap.
func (k KeyMap) FullHelp() [][]key.Binding {
    return [][]key.Binding{
        {k.Blur, k.Print},
        {k.Help, k.Quit},
    }
}

// ShortHelp implements help.KeyMap.
func (k KeyMap) ShortHelp() []key.Binding {
    return []key.Binding{k.Help, k.Quit}
}

var DefaultKeyMap = KeyMap{
	KeyMap: textarea.DefaultKeyMap,
	Blur: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "toggle focus"),
	),
	Print: key.NewBinding(
		key.WithKeys("ctrl+p"),
		key.WithHelp("ctrl+p", "print"),
	),
	Help: key.NewBinding(
		key.WithKeys("ctrl+h"),
		key.WithHelp("ctrl+h", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "ctrl+q"),
		key.WithHelp("ctrl+c/ctrl+q", "quit"),
	),
}
