package escpos

import "io"

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
type SetRotation struct{
    Enabled bool
}

func (cmd SetRotation) WriteTo(w io.Writer) (int64, error) {
    return write(w, esc, 'V', bit(cmd.Enabled))
}
