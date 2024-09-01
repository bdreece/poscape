package printer

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
}

// FullHelp implements help.KeyMap.
func (k KeyMap) FullHelp() [][]key.Binding {
	panic("unimplemented")
}

// ShortHelp implements help.KeyMap.
func (k KeyMap) ShortHelp() []key.Binding {
	panic("unimplemented")
}

func DefaultKeyMap() KeyMap {
    return KeyMap{}
}
