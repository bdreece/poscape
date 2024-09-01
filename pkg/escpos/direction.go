//go:generate go run golang.org/x/tools/cmd/stringer@latest -type PrintDirection -trimprefix Print
package escpos

import (
	"fmt"
	"io"
)

type (
	// PrintDirection specifies the print direction in page mode.
	PrintDirection byte

	setPrintDirection struct {
		direction PrintDirection
	}
)

const (
	PrintLeftToRight PrintDirection = iota
	PrintBottomToTop
	PrintRightToLeft
	PrintTopToBottom
)

func (d PrintDirection) GoString() string {
	return fmt.Sprintf("escpos.PrintDirection(%s)", d)
}

// WriteTo implements Command.
func (cmd setPrintDirection) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, 'T', byte(cmd.direction))
}

// Set print direction in page mode.
//
// Note:
//
//   - When the command is input in standard mode, the printer executes only
//     the internal flag operation. This command does not affect printing in
//     standard mode.
func SetPrintDirection(direction PrintDirection) Command {
	return setPrintDirection{direction}
}
