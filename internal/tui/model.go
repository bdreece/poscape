package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/bdreece/poscape/pkg/escpos"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	keymap KeyMap
	styles Styles

	textarea textarea.Model
	help     help.Model

	driver escpos.Driver
	err    error
	status string
}

func (m Model) Init() tea.Cmd {
	return tea.Sequence(m.textarea.Focus(), textarea.Blink)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
		m.textarea.SetWidth(msg.Width)
		m.textarea.SetHeight(msg.Height - 5)
	case tea.KeyMsg:
		cmd = m.handleKey(msg)
		cmds = append(cmds, cmd)
    case errorMsg:
        m.err = msg.Err
        cmd = dispatchClear[errorMsg](10 * time.Second)
        cmds = append(cmds, cmd)
	case clearMsg[errorMsg]:
		m.err = nil
	case statusMsg:
		m.status = msg.Status
		cmd = dispatchClear[statusMsg](10 * time.Second)
		cmds = append(cmds, cmd)
	case clearMsg[statusMsg]:
		m.status = "\n"
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)

	m.help, cmd = m.help.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	const tpl = "%s\n%s%s\n"

	return m.styles.Main.Render(fmt.Sprintf(tpl,
		m.textarea.View(),
        m.statusView(),
		m.helpView()))
}

func (m Model) Run() error {
    program := tea.NewProgram(m,
        tea.WithAltScreen(),
        tea.WithMouseCellMotion())

    if _, err := program.Run(); err != nil {
        return err
    }

    return nil
}

func (m *Model) SetKeyMap(keymap KeyMap) {
	m.keymap = keymap
	m.textarea.KeyMap = keymap.KeyMap
}

func (m *Model) SetStyles(styles Styles) {
	m.styles = styles
	m.textarea.FocusedStyle = styles.Textarea.Focused
	m.textarea.BlurredStyle = styles.Textarea.Blurred
	m.help.Styles = styles.Help
}

func (m *Model) handleKey(msg tea.KeyMsg) tea.Cmd {
	switch {
	case key.Matches(msg, m.keymap.Quit):
		return tea.Quit
	case key.Matches(msg, m.keymap.Help):
		m.help.ShowAll = !m.help.ShowAll
	case key.Matches(msg, m.keymap.Print):
		return dispatchPrint(m.driver, m.textarea.Value())
	case key.Matches(msg, m.keymap.Blur):
		if m.textarea.Focused() {
			m.textarea.Blur()
		}
	default:
		if !m.textarea.Focused() {
			return m.textarea.Focus()
		}
	}

	return nil
}

func (m Model) statusView() string {
    if m.err != nil {
        return fmt.Sprintf("%v\n", m.err)
    }

    return m.status
}

func (m Model) helpView() string {
    view := m.help.View(m.keymap)
    padding := strings.Repeat("\n", 2 - strings.Count(view, "\n"))
    return padding + view
}

func New(driver escpos.Driver) Model {
    m := Model{
		textarea: textarea.New(),
		help:     help.New(),
		driver:   driver,
		status:   "\n",
	}

    m.SetKeyMap(DefaultKeyMap)
    m.SetStyles(DefaultStyles)
    return m
}
