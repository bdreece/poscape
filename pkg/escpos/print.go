package escpos

import "io"

type (
	// Print data in page mode.
	//
	// In page mode, prints all buffered data in the printable area collectively.
	//
	// Note:
	//
	//  - This command is enabled only in page mode.
	//  - After printing, the printer does not clear the buffered data, the values
	//    set by [SetPrintDirection] or [SetPrintingArea], nor the position for
	//    buffering character data.
	PrintPage struct{}

	// Print and return to standard mode in page mode.
	//
	// Prints the data in the print buffer and returns to standard mode.
	//
	// Note:
	//
	//  - The buffer data is deleted after being printed.
	//  - The printing area set by [SetPrintingArea] is reset to the default setting.
	//  - This command sets the print position to the beginning of the line.
	//  - This command is enabled only in page mode.
	PrintAndEnterStandardMode struct{}

	// Print and line feed.
	PrintLine struct{}

	// Print and carriage return.
	PrintCarriageReturn struct{}

	// Print and feed paper.
	//
	// Prints the data in the print buffer and feeds the paper (n x vertical or
	// horizontal motion unit).
	//
	// Note:
	//
	//   - The maximum feed amount is 1016mm (40").
	//   - After printing is completed, this command sets the printing start
	//     position to the beginning of the line.
	//   - The paper feed amount set by this comamnd does not affect the values
	//     set by [SetLineSpacing] or [ResetLineSpacing].
	//   - In standard mode, the vertical motion unit is used.
	//   - In page mode, usage of either the vertical or horizontal motion unit is
	//     determined by the start position of the printable area.
	PrintAndFeed struct {
		N uint8
	}

	// Print and feed lines.
	//
	// Prints the data in the print buffer and feeds n lines.
	//
	// Note:
	//
	//   - The maximum feed amount is 1016mm (40").
	//   - After printing is completed, this command sets the printing start
	//     position to the beginning of the line.
	//   - The paper feed amount set by this comamnd does not affect the values
	//     set by [SetLineSpacing] or [ResetLineSpacing].
	PrintAndFeedLines struct {
		Lines uint8
	}
)

// WriteTo implements Command.
func (PrintPage) WriteTo(w io.Writer) (int64, error) { return write(w, esc, ff) }

// WriteTo implements Command.
func (PrintAndEnterStandardMode) WriteTo(w io.Writer) (int64, error) { return write(w, ff) }

// WriteTo implements Command.
func (PrintLine) WriteTo(w io.Writer) (int64, error) { return write(w, '\n') }

// WriteTo implements Command.
func (PrintCarriageReturn) WriteTo(w io.Writer) (int64, error) { return write(w, '\r') }

// WriteTo implements Command.
func (cmd PrintAndFeed) WriteTo(w io.Writer) (int64, error) {
    return write(w, esc, 'J', cmd.N)
}

// WriteTo implements Command.
func (cmd PrintAndFeedLines) WriteTo(w io.Writer) (int64, error) {
    return write(w, esc, 'd', cmd.Lines)
}
