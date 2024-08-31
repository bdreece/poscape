package escpos

import "io"

type (
	// Select default line spacing.
	//
	// Selects approximately 4.23mm (1/6") spacing.
	//
	// Note:
	//
	//   - The line spacing can be set independently in standard mode and in page
	//     mode.
	SetDefaultLineSpacing struct{}

	// Set line spacing.
	//
	// Sets the line spacing to (spacing x vertical or horizontal motion unit) inches.
	//
	// Note:
	//
	//   - The line spacing can be set independently in standard mode and in page
	//     mode.
	//   - The maximum feed amount is 1016mm (40").
	//   - In standard mode, the vertical motion unit is used.
	//   - In page mode, usage of either the vertical or horizontal motion unit is
	//     determined by the start position of the printable area.
	SetLineSpacing struct {
		Spacing uint8
	}
)

// WriteTo implements Command.
func (SetDefaultLineSpacing) WriteTo(w io.Writer) (int64, error) { return write(w, esc, '2') }

// WriteTo implements Command.
func (cmd SetLineSpacing) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, '3', cmd.Spacing)
}
