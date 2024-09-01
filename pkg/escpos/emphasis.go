package escpos

import "io"

type setEmphasis struct {
	enabled bool
}

// WriteTo implements Command.
func (cmd setEmphasis) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, 'E', bit(cmd.enabled))
}

// Turn emphasis mode on/off.
//
// Note:
//
//   - Emphasis mode can also be toggled by using [SetPrintMode]. However, the
//     setting of the last received command is effective.
func SetEmphasis(enabled bool) Command {
	return setEmphasis{enabled}
}
