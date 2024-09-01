package window

import (
	"github.com/bdreece/poscape/pkg/tui/editor"
	"github.com/bdreece/poscape/pkg/tui/printer"
	"github.com/bdreece/poscape/pkg/tui/viewer"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	Model struct {
		width  int
		height int
		keys   KeyMap

		help help.Model
		page pageModel
	}

	pageModel interface {
		help.KeyMap

		Init() tea.Cmd
		View() string
	}
)

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.editor.Init(), m.printer.Init(), m.viewer.Init())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch true {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = true
		case key.Matches(msg, m.keys.Edit):
			m.page = editor.New()
			cmds = append(cmds, m.page.Init())
		case key.Matches(msg, m.keys.Print):
			m.page = printer.New()
			cmds = append(cmds, m.page.Init())
		case key.Matches(msg, m.keys.View):
			m.page = viewer.New()
			cmds = append(cmds, m.page.Init())
		}
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Center,
		m.page.View(),
		m.help.View(m.page),
	)
}

func New(width, height int) Model {
	return Model{
		width:  width,
		height: height,
		keys:   DefaultKeyMap(),
		help:   help.New(),
		page:   editor.New(),
	}
}
