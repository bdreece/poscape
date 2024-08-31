package escpos

import "io"

// Set printing area in page mode.
//
// Each setting for the printable area is multiplied by the corresponding
// motion unit (horizontal for x, w; vertical for y, h) to provide
// measurements in inches.
//
// Note:
//
//   - If this command is input in standard mode, the printer executes only
//     the internal flag operation. This command does not affect printing in
//     standard mode.
//   - If x or y are set outside the printable area, the printer stops command
//     processing and processes the following data as normal data.
type SetPrintingArea struct {
	X, Y, W, H uint16
}

func (cmd SetPrintingArea) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, 'W',
		byte(cmd.X & 0x0F), byte(cmd.X & 0xF0),
		byte(cmd.Y & 0x0F), byte(cmd.Y & 0xF0),
		byte(cmd.W & 0x0F), byte(cmd.W & 0xF0),
		byte(cmd.H & 0x0F), byte(cmd.H & 0xF0),
	)
}
