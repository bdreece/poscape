package viewer

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

type (
	Model struct {
		keys   KeyMap
		styles Style

        renderer glamour.TermRenderer
		viewport viewport.Model
	}

	Params struct {
		Width  int
		Height int
		Keys   KeyMap
		Styles Style
	}
)

func (Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    m.viewport, cmd = m.viewport.Update(msg)

    return m, cmd
}

func (m Model) View() string {
    return m.viewport.View()
}

func New(p Params) Model {
	return Model{
		keys:     p.Keys,
		styles:   p.Styles,
		viewport: viewport.New(p.Width, p.Height),
	}
}
