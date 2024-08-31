package escpos

import "io"

// Set left margin.
//
// Note:
//
//   - The left margin is set to (margin x horizontal or vertical motion unit)
//     inches.
//   - Settings outside the specified printable area are ignored.
//   - If this command is input in page mode, the printer only performs the
//     internal flag operation. This command has no effect in page mode.
type SetLeftMargin struct{
    Margin uint16
}

// WriteTo implements Command.
func (cmd SetLeftMargin) WriteTo(w io.Writer) (int64, error) {
    return write(w, gs, 'L', byte(cmd.Margin << 4) | byte(cmd.Margin & 0x0F))
}
