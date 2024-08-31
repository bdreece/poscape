package escpos

import "io"

type (
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
	SetAbsoluteX struct {
		Position uint16
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
	SetAbsoluteY struct {
		Position uint16
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
	SetRelativeX struct {
		Position uint16
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
	SetRelativeY struct {
		Position uint16
	}
)

// WriteTo implements Command.
func (cmd SetAbsoluteX) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, '$', byte(cmd.Position), byte(cmd.Position>>8))
}

// WriteTo implements Command.
func (cmd SetAbsoluteY) WriteTo(w io.Writer) (int64, error) {
	return write(w, gs, '$', byte(cmd.Position), byte(cmd.Position>>8))
}

// WriteTo implements Command.
func (cmd SetRelativeX) WriteTo(w io.Writer) (int64, error) {
    return write(w, esc, '\\', byte(cmd.Position), byte(cmd.Position>>8))
}

// WriteTo implements Command.
func (cmd SetRelativeY) WriteTo(w io.Writer) (int64, error) {
    return write(w, gs, '\\', byte(cmd.Position), byte(cmd.Position>>8))
}
