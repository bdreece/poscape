package escpos

import "io"

type setUpsideDown struct {
	enabled bool
}

func (cmd setUpsideDown) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, '{', bit(cmd.enabled))
}

// Turn upside-down printing mode on/off.
//
// Note:
//
//   - This command is enabled only when processed at the beginning of a line
//     in standard mode.
//   - If this command is input in page mode, the printer only performs the
//     internal flag operation. This command has no effect in page mode.
//   - In upside-down printing mode, both the text rotation and the line order
//     are flipped, such that content order is preserved.
func SetUpsideDown(enabled bool) Command {
	return setUpsideDown{enabled}
}
