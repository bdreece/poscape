package escpos

import "io"

type setDoubleStrike struct {
    enabled bool
}

func (cmd setDoubleStrike) WriteTo(w io.Writer) (int64, error) {
    return write(w, esc, 'G', bit(cmd.enabled))
}

// Turn double-strike mode on/off.
//
// Note:
//
//   - Printer output is the same in double-strike mode and in emphasis mode.
//   - Double-strike mode can also be toggled by using [SetPrintMode]. However, the
//     setting of the last received command is effective.
func SetDoubleStrike(enabled bool) Command {
    return setDoubleStrike{enabled}
}
