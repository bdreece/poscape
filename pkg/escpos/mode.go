//go:generate go run golang.org/x/tools/cmd/stringer@latest -type PrintMode -trimprefix With
package escpos

import (
	"fmt"
	"io"
)

type (
	// PrintMode specifies bit-flag print settings, which may be combined using
	// bitwise-OR.
	PrintMode byte

	// Select print mode(s).
	//
	// Note:
	//
	//   - The thickness of the underline is selected by SetEmphasisMode, regardless of
	//     the character size.
	//   - When some characters in a line are double or more height, all the characters
	//     on the line are aligned at the baseline.
	//   - [SetEmphasisMode] and [SetUnderlineMode] can also toggle emphasis and underline,
	//     respectively. However, the setting of the last received command is effective.
	setPrintMode struct {
		mode PrintMode
	}
)

const (
	// Print with alternate font.
	WithAlternateFont PrintMode = 1 << iota

	_
	_

	// Print with emphasis mode on.
	WithEmphasis

	// Prints with double height.
	//
	// If WithDoubleWidth is also selected, characters will be quadruple-sized.
	WithDoubleHeight

	// Prints with double width.
	//
	// If WithDoubleHeight is also selected, characters will be quadruple-sized.
	WithDoubleWidth

	_

	// Prints with underline mode on.
	WithUnderline
)

func (mode PrintMode) GoString() string {
	return fmt.Sprintf("escpos.PrintMode(%s)", mode)
}

// WriteTo implements Command.
func (cmd setPrintMode) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, '!', byte(cmd.mode))
}

func SetPrintMode(mode PrintMode) Command {
	return setPrintMode{mode}
}
