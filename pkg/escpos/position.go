package escpos

import "io"

type (
	setAbsoluteX struct {
		position uint16
	}

	setAbsoluteY struct {
		position uint16
	}

	setRelativeX struct {
		position uint16
	}

	setRelativeY struct {
		position uint16
	}
)

// WriteTo implements Command.
func (cmd setAbsoluteX) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, '$', byte(cmd.position), byte(cmd.position>>8))
}

// Set absolute print position.
//
// Sets the distance from the beginning of the line to the position at which
// subsequent characters are to be printed.
//
// Note:
//
//   - The distance from the beginning of the line to the print position is
//     (pos * vertical or horizontal motion unit) inches.
//   - Settings outside the specified printable area are ignored.
//   - In standard mode, the horizontal motion unit is used.
//   - In page mode, usage of either the horizontal or vertical motion unit is
//     determined by the start position of the printable area.
func SetAbsoluteX(pos uint16) Command {
	return setAbsoluteX{pos}
}

// WriteTo implements Command.
func (cmd setAbsoluteY) WriteTo(w io.Writer) (int64, error) {
	return write(w, gs, '$', byte(cmd.position), byte(cmd.position>>8))
}

// Set absolute vertical print position in page mode.
//
// Note:
//
//   - This commands sets the absolute vertical print position to
//     (pos * vertical or horizontal motion unit) inches.
//   - Settings outside the specified printable area are ignored.
//   - In standard mode, the vertical motion unit is used.
//   - In page mode, usage of either the horizontal or vertical motion unit is
//     determined by the start position of the printable area.
func SetAbsoluteY(pos uint16) Command {
	return setAbsoluteY{pos}
}

// WriteTo implements Command.
func (cmd setRelativeX) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, '\\', byte(cmd.position), byte(cmd.position>>8))
}

// Set relative print position.
//
// Sets the print starting position based on the current position.
//
// Note:
//
//   - The distance from the current position to the print position is
//     (pos * vertical or horizontal motion unit).
//   - Settings outside the specified printable area are ignored.
//   - In standard mode, the horizontal motion unit is used.
//   - In page mode, usage of either the horizontal or vertical motion unit is
//     determined by the start position of the printable area.
func SetRelativeX(pos uint16) Command {
	return setRelativeX{pos}
}

// WriteTo implements Command.
func (cmd setRelativeY) WriteTo(w io.Writer) (int64, error) {
	return write(w, gs, '\\', byte(cmd.position), byte(cmd.position>>8))
}

// Set relative vertical print position in page mode.
//
// Sets the vertical print starting position relative to the current
// position in page mode.
//
// Note:
//
//   - This commands sets the vertical print position relative to the current
//     position to (pos * vertical or horizontal motion unit) inches.
//   - Settings outside the specified printable area are ignored.
//   - In standard mode, the vertical motion unit is used.
//   - In page mode, usage of either the horizontal or vertical motion unit is
//     determined by the start position of the printable area.
func SetRelativeY(pos uint16) Command {
	return setRelativeY{pos}
}
