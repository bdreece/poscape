package escpos

import "io"

type setMotionUnits struct {
	x, y uint8
}

// WriteTo implements Command.
func (cmd setMotionUnits) WriteTo(w io.Writer) (int64, error) {
	return write(w, gs, 'P', cmd.x, cmd.y)
}

// Set horizontal and vertical motion units.
//
// Sets the horizontal and vertical motion units to 1/x inches and 1/y inches,
// respectively. When x and y are set to 0, the default settings are used (x =
// 180, y = 360).
//
// Note:
//
//   - In standard mode, the vertical direction is the feed direction.
//   - In page mode, usage of the feed direction as either the vertical or
//     horizontal direction determined by the start position of the printable
//     area.
func SetMotionUnits(x, y uint8) Command {
	return setMotionUnits{x, y}
}
