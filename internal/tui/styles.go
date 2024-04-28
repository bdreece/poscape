package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/lipgloss"
)

type TextareaStyles struct {
    Focused textarea.Style
    Blurred textarea.Style
}

type Styles struct {
	Main     lipgloss.Style
	Textarea TextareaStyles
	Help     help.Styles
}

var DefaultStyles = Styles{
    Main: lipgloss.NewStyle(),
    Help: help.New().Styles,
}

func init() {
    DefaultStyles.Textarea.Focused, DefaultStyles.Textarea.Blurred = textarea.DefaultStyles()
}
