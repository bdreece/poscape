package editor

import (
	"github.com/bdreece/poscape/pkg/escpos"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	Model struct {
		KeyMap

		width  int
		height int
		cmds   []escpos.Command
	}

	Options struct {
		Width    int
		Height   int
		Commands []escpos.Command
	}
)

func (Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return ""
}

func New() Model {
	return Model{}
}
