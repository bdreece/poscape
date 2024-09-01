package escpos

import "io"

type (
	printPage                 struct{}
	printAndEnterStandardMode struct{}
	lineFeed                  struct{}
	carriageReturn            struct{}

	printAndFeed struct {
		n uint8
	}

	printAndFeedLines struct {
		lines uint8
	}
)

// WriteTo implements Command.
func (printPage) WriteTo(w io.Writer) (int64, error) { return write(w, esc, ff) }

// Print data in page mode.
//
// In page mode, prints all buffered data in the printable area collectively.
//
// Note:
//
//   - This command is enabled only in page mode.
//   - After printing, the printer does not clear the buffered data, the values
//     set by [SetPrintDirection] or [SetPrintingArea], nor the position for
//     buffering character data.
func PrintPage() Command { return *new(printPage) }

// WriteTo implements Command.
func (printAndEnterStandardMode) WriteTo(w io.Writer) (int64, error) { return write(w, ff) }

// Print and return to standard mode in page mode.
//
// Prints the data in the print buffer and returns to standard mode.
//
// Note:
//
//   - The buffer data is deleted after being printed.
//   - The printing area set by [SetPrintingArea] is reset to the default setting.
//   - This command sets the print position to the beginning of the line.
//   - This command is enabled only in page mode.
func PrintAndEnterStandardMode() Command { return *new(printAndEnterStandardMode) }

// WriteTo implements Command.
func (lineFeed) WriteTo(w io.Writer) (int64, error) { return write(w, '\n') }

// Print and line feed.
func LineFeed() Command { return *new(lineFeed) }

// WriteTo implements Command.
func (carriageReturn) WriteTo(w io.Writer) (int64, error) { return write(w, '\r') }

// Print and carriage return.
func CarriageReturn() Command { return *new(carriageReturn) }

// WriteTo implements Command.
func (cmd printAndFeed) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, 'J', cmd.n)
}

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
func PrintAndFeed(n uint8) Command {
	return printAndFeed{n}
}

// WriteTo implements Command.
func (cmd printAndFeedLines) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, 'd', cmd.lines)
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
func PrintAndFeedLines(lines uint8) Command {
	return printAndFeedLines{lines}
}
