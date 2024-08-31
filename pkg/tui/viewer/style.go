package viewer

import (
	"github.com/charmbracelet/lipgloss"
)

type Style struct {
    Viewport lipgloss.Style
}

var DefaultStyles = Style{
    Viewport: lipgloss.NewStyle(),
}
