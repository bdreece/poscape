package escpos

import "io"

// Turn double-strike mode on/off.
//
// Note:
//
//   - Printer output is the same in double-strike mode and in emphasis mode.
//   - Double-strike mode can also be toggled by using [SetPrintMode]. However, the
//     setting of the last received command is effective.
type SetDoubleStrike struct {
    Enabled bool
}

func (cmd SetDoubleStrike) WriteTo(w io.Writer) (int64, error) {
    return write(w, esc, 'G', bit(cmd.Enabled))
}

