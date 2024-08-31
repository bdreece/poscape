package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
)

func Middleware() wish.Middleware {
	return bubbletea.Middleware(func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
		pty, _, _ := s.Pty()
		renderer := bubbletea.MakeRenderer(s)

		m := model{
			term:    pty.Term,
			profile: renderer.ColorProfile().Name(),
			width:   pty.Window.Width,
			height:  pty.Window.Height,
			keys:    defaultKeys,
		}

		return m, []tea.ProgramOption{
			tea.WithAltScreen(),
		}
	})
}
