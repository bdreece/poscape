package viewer

import "github.com/charmbracelet/bubbles/viewport"

type KeyMap struct {
    Viewport viewport.KeyMap
}

var DefaultKeys = KeyMap{
    Viewport: viewport.DefaultKeyMap(),
}
