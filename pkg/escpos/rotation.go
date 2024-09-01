package escpos

import "io"

type setRotation struct {
	enabled bool
}

func (cmd setRotation) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, 'V', bit(cmd.enabled))
}

// Turn 90° clockwise rotation on/off.
//
// Note:
//
//   - When underline mode is turned on, the printer does not underline 90°
//     rotated characters.
//   - Double-width and double-height commands in 90° rotation mode enlarge
//     characters in the opposite directions from their respective counterparts
//     in normal rotation mode.
//   - This command affects printing in standard mode. However, the setting is
//     always effective.
func SetRotation(enabled bool) Command {
	return setRotation{enabled}
}
