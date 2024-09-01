package tui

import (
	"reflect"

	"github.com/bdreece/poscape/pkg/escpos"
	tea "github.com/charmbracelet/bubbletea"
)

type StartInsertMsg struct {
	Index int
	Type  reflect.Type
}

type EndInsertMsg struct {
	Index   int
	Command escpos.Command
}

func StartInsert[T escpos.Command](index int) tea.Cmd {
	return func() tea.Msg {
		return StartInsertMsg{
			Index: index,
			Type:  reflect.TypeFor[T](),
		}
	}
}

func EndInsert(index int, cmd escpos.Command) tea.Cmd {
	return func() tea.Msg {
		return EndInsertMsg{
			Index:   index,
			Command: cmd,
		}
	}
}
