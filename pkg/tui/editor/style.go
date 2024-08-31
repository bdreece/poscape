package editor

import "github.com/charmbracelet/bubbles/textarea"

type Style struct {
    Focused textarea.Style
    Blurred textarea.Style
}

var DefaultStyles Style

func init() {
    DefaultStyles.Focused, DefaultStyles.Blurred = textarea.DefaultStyles()
}
