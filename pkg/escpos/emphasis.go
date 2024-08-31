package escpos

import "io"

// Turn emphasis mode on/off.
//
// Note:
//
//   - Emphasis mode can also be toggled by using [SetPrintMode]. However, the
//     setting of the last received command is effective.
type SetEmphasis struct {
    Enabled bool
}

// WriteTo implements Command.
func (cmd SetEmphasis) WriteTo(w io.Writer) (int64, error) {
    return write(w, esc, 'E', bit(cmd.Enabled))
}
