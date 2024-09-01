package window

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Help  key.Binding
	Quit  key.Binding
	Edit  key.Binding
	View  key.Binding
	Print key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
        Help: key.NewBinding(
            key.WithKeys("h", "ctrl+h"),
            key.WithHelp("h/ctrl+h", "help"),
        ),
		Edit: key.NewBinding(
			key.WithKeys("e", "ctrl+e"),
			key.WithHelp("e/ctrl+e", "edit"),
		),
		View: key.NewBinding(
			key.WithKeys("v", "ctrl+v"),
			key.WithHelp("v/ctrl+v", "view"),
		),
		Print: key.NewBinding(
			key.WithKeys("p", "ctrl+p"),
			key.WithHelp("p/ctrl+p", "print"),
		),
	}
}
