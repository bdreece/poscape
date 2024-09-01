package escpos

import "io"

type setLeftMargin struct {
	margin uint16
}

// WriteTo implements Command.
func (cmd setLeftMargin) WriteTo(w io.Writer) (int64, error) {
	return write(w, gs, 'L', byte(cmd.margin<<4)|byte(cmd.margin&0x0F))
}

// Set left margin.
//
// Note:
//
//   - The left margin is set to (margin x horizontal or vertical motion unit)
//     inches.
//   - Settings outside the specified printable area are ignored.
//   - If this command is input in page mode, the printer only performs the
//     internal flag operation. This command has no effect in page mode.
func SetLeftMargin(margin uint16) Command {
	return setLeftMargin{margin}
}
