//go:generate go run golang.org/x/tools/cmd/stringer@latest -type Underline -trimprefix Underline
package escpos

import "io"

type (
	// Underline specifies the kind of underline.
	Underline byte

	// Turn underline mode on/off.
	//
	// Note:
	//
	//   - The printer can underline all characters (including the right-side spacing),
	//     but cannot underline the space set by HT.
	//   - The printer cannot underline 90° clockwise rotated characters and white/black
	//     inverted characters.
	//   - Changing the character size does not affect the current underline thickness.
	//   - Underline mode can also be toggled by using [SetPrintMode]. However, the
	//     setting of the last received command is effective.
	SetUnderline struct {
		Underline Underline
	}
)

const (
	// 0px underline.
	UnderlineNone Underline = iota
	// 1px underline.
	UnderlineSingle
	// 2px underline.
	UnderlineDouble
)

func (cmd SetUnderline) WriteTo(w io.Writer) (int64, error) {
    return write(w, esc, '-', byte(cmd.Underline))
}
