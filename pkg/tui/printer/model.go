package printer

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
    KeyMap
}

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
