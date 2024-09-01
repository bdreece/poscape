package escpos

import "io"

type setPrintingArea struct {
	x, y, w, h uint16
}

func (cmd setPrintingArea) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, 'W',
		byte(cmd.x & 0x0F), byte(cmd.x & 0xF0),
		byte(cmd.y & 0x0F), byte(cmd.y & 0xF0),
		byte(cmd.w & 0x0F), byte(cmd.w & 0xF0),
		byte(cmd.h & 0x0F), byte(cmd.h & 0xF0),
	)
}

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
func SetPrintingArea(x, y, w, h uint16) Command {
    return setPrintingArea{x, y, w, h}
}
