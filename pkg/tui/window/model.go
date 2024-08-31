package window

import (
	"github.com/bdreece/poscape/pkg/tui"
	"github.com/bdreece/poscape/pkg/tui/editor"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	Model struct {
		term    string
		profile string
		width   int
		height  int
		keys    KeyMap

		help help.Model
		view tui.ViewModel
	}

	Params struct {
		Term    string
		Profile string
		Width   int
		Height  int
		Keys    KeyMap
	}
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return ""
}

func New(p Params) Model {
	return Model{
		term:    p.Term,
		profile: p.Profile,
		width:   p.Width,
		height:  p.Height,
		keys:    p.Keys,
		help:    help.New(),
		view:    editor.New(),
	}
}
