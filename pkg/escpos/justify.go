//go:generate go run golang.org/x/tools/cmd/stringer@latest -type Justification -trimprefix Justify
package escpos

import (
	"fmt"
	"io"
)

type (
	// JustifyMode specifies how to justify text perpendicular to the
	// feed direction.
	Justification byte

	setJustification struct {
		justify Justification
	}
)

const (
	JustifyLeft Justification = iota
	JustifyCenter
	JustifyRight
)

// GoString implements fmt.GoStringer
func (j Justification) GoString() string {
	return fmt.Sprintf("escpos.Justification(%s)", j)
}

// WriteTo implements Command.
func (cmd setJustification) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, 'a', byte(cmd.justify))
}

// Select justification.
//
// Note:
//
//   - The command is enabled only when processed at the beginning of the line
//     in standard mode.
//   - If this command is input in page mode, the printer only performs the
//     internal flag operation. This command has no effect in page mode.
//   - This command justifies based on the specified printing area.
func SetJustification(justify Justification) Command {
	return setJustification{justify}
}
