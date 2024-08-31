package editor

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	keys     KeyMap
	style    Style
	textarea textarea.Model
}

func (m Model) Keys() help.KeyMap        { return m.keys }
func (m Model) Textarea() textarea.Model { return m.textarea }

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds = make([]tea.Cmd, 0)
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch true {
		case key.Matches(msg, m.keys.Focus):
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		default:
			if !m.textarea.Focused() {
				cmds = append(cmds, m.textarea.Focus())
			}
		}
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	m.textarea.FocusedStyle = m.style.Focused
	m.textarea.BlurredStyle = m.style.Blurred

	return m.textarea.View()
}

func New(keys KeyMap, style Style) Model {
	return Model{
		keys:     keys,
		style:    style,
		textarea: textarea.New(),
	}
}

var _ (tea.Model) = (*Model)(nil)
