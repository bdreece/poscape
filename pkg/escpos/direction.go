//go:generate go run golang.org/x/tools/cmd/stringer@latest -type PrintDirection -trimprefix Print
package escpos

import "io"

type (
	// PrintDirection specifies the print direction in page mode.
	PrintDirection byte

	// Set print direction in page mode.
	//
	// Note:
	//
	//   - When the command is input in standard mode, the printer executes only
	//     the internal flag operation. This command does not affect printing in
	//     standard mode.
	SetPrintDirection struct {
		Direction PrintDirection
	}
)

const (
	PrintLeftToRight PrintDirection = iota
	PrintBottomToTop
	PrintRightToLeft
	PrintTopToBottom
)

// WriteTo implements Command.
func (cmd SetPrintDirection) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, 'T', byte(cmd.Direction))
}
