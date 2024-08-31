package tui

import "github.com/charmbracelet/bubbles/key"

type (
	KeyMap struct {
		Help key.Binding
		Quit key.Binding

		NormalMode key.Binding
	}

	NormalModeKeyMap struct {
		EditMode  key.Binding
		ViewMode  key.Binding
		PrintMode key.Binding
	}
)

var (
	DefaultKeys = KeyMap{
		Help: key.NewBinding(
			key.WithKeys("ctrl+h"),
			key.WithHelp("ctrl+h", "toggle help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "quit"),
		),
		NormalMode: key.NewBinding(
			key.WithKeys("esc"),
            key.WithHelp("esc", "normal mode"),
		),
	}

	DefaultNormalModeKeys = NormalModeKeyMap{
		EditMode: key.NewBinding(
            key.WithKeys("e", "ctrl+e"),
            key.WithHelp("e", "edit mode"),
        ),
        ViewMode: key.NewBinding(
            key.WithKeys("v", "ctrl+v"),
            key.WithHelp("v", "view mode"),
        ),
        PrintMode: key.NewBinding(
            key.WithKeys("p", "ctrl+p"),
            key.WithHelp("p", "print mode"),
        ),
	}
)

func (km KeyMap) ShortHelp() []key.Binding {
    return []key.Binding{km.NormalMode, km.Help, km.Quit}
}

func (km KeyMap) FullHelp() [][]key.Binding {
    return [][]key.Binding{km.ShortHelp()}
}

func (km NormalModeKeyMap) ShortHelp() []key.Binding {
    return []key.Binding{km.EditMode, km.ViewMode, km.PrintMode}
}

func (km NormalModeKeyMap) FullHelp() [][]key.Binding {
    return [][]key.Binding{km.ShortHelp()}
}
