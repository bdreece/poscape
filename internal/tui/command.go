package tui

import (
	"fmt"
	"time"

	"github.com/bdreece/poscape/pkg/escpos"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	clearMsg[T tea.Msg] struct{}
	statusMsg           struct {
		Status string
	}
	errorMsg struct {
		Err error
	}
)

func dispatchClear[T tea.Msg](d time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(d)
		return clearMsg[T]{}
	}
}

func dispatchPrint(driver escpos.Driver, content string) tea.Cmd {
	return func() tea.Msg {
		if err := driver.PrintAndFeed(4); err != nil {
			return errorMsg{Err: err}
		}

		n, err := driver.Write([]byte(content))
		if err != nil {
			return errorMsg{Err: err}
		}

		if err := driver.PrintAndFeed(4); err != nil {
			return errorMsg{Err: err}
		}

		return statusMsg{Status: fmt.Sprintf("printed %d bytes!\n", n)}
	}
}
