package editor

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
)

type KeyMap struct {
	Textarea textarea.KeyMap

	Focus key.Binding
}

var DefaultKeyMap = KeyMap{
	Textarea: textarea.DefaultKeyMap,
	Focus: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "toggle focus"),
	),
}

// ShortHelp implements help.KeyMap.
func (km KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		km.Textarea.LineNext,
		km.Textarea.LinePrevious,
		km.Textarea.CharacterForward,
		km.Textarea.CharacterBackward,
	}
}

// FullHelp implements help.KeyMap.
func (km KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			km.Textarea.WordForward,
			km.Textarea.WordBackward,
			km.Textarea.InsertNewline,
			km.Textarea.Paste,
		},
		{
			km.Textarea.InputBegin,
			km.Textarea.InputEnd,
			km.Textarea.LineStart,
			km.Textarea.LineEnd,
		},
		{
			km.Textarea.LineNext,
			km.Textarea.LinePrevious,
			km.Textarea.CharacterForward,
			km.Textarea.CharacterBackward,
		},
	}
}

var _ help.KeyMap = (*KeyMap)(nil)
