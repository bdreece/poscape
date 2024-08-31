package escpos

import "io"

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
type SetMotionUnits struct{
    X, Y uint8
}

func (cmd SetMotionUnits) WriteTo(w io.Writer) (int64, error) {
    return write(w, gs, 'P', cmd.X, cmd.Y)
}
