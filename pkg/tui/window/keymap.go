package window

import (
	"github.com/bdreece/poscape/pkg/tui"
	"github.com/bdreece/poscape/pkg/tui/editor"
)

type KeyMap struct {
	Global tui.KeyMap
    Editor editor.KeyMap
}

