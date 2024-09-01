//go:generate go run golang.org/x/tools/cmd/stringer@latest -type Font -trimprefix Font
//go:generate go run golang.org/x/tools/cmd/stringer@latest -type Charset -linecomment
//go:generate go run golang.org/x/tools/cmd/stringer@latest -type CodePage -linecomment
package escpos

import (
	"fmt"
	"io"
)

type (
	// Font specifies a built-in font-family.
	Font byte

	// Charset specifies an international character set.
	//
	// The international character set defines glyphs used
	// for the following symbols:
	//
	//	# $ @ [ \ ] ^ ` { | } ~
	Charset byte

	// CodePage specifies the ASCII code page used to decode the text.
	CodePage byte

	setCharacterSize struct {
		W, H uint8
	}

	setKerning struct {
		kerning uint8
	}

	setCharset struct {
		charset Charset
	}

	setFont struct {
		font Font
	}

	setCodePage struct {
		codePage CodePage
	}
)

const (
	// Character font A (12x24px)
	FontA Font = iota
	// Character font B (9x24px)
	FontB

	CharsetUnitedStates  Charset = iota // U.S.A.
	CharsetFrance                       // France
	CharsetGermany                      // Germany
	CharsetUnitedKingdom                // U.K.
	CharsetDenmark1                     // Denmark I
	CharsetSweden                       // Sweden
	CharsetItaly                        // Italy
	CharsetSpain1                       // Spain I
	CharsetJapan                        // Japan
	CharsetNorway                       // Norway
	CharsetDenmark2                     // Denmark II
	CharsetSpain2                       // Spain II
	CharsetLatinAmerica                 // Latin America

	CodePagePC437    CodePage = iota // PC437 [U.S.A., Standard Europe]
	CodePageKatakana                 // Katakana
	CodePagePC850                    // PC850 [Multilingual]
	CodePagePC860                    // PC860 [Portuguese]
	CodePagePC863                    // PC863 [Canadian-French]
	CodePagePC865                    // PC865 [Nordic]
	CodePagePC866    CodePage = 17   // PC866 [Cyrillic #2]
	CodePageSpace    CodePage = 255  // Space page
)

// GoString implements fmt.GoStringer
func (f Font) GoString() string {
	return fmt.Sprintf("escpos.Font(%s)", f)
}

// GoString implements fmt.GoStringer
func (c Charset) GoString() string {
	return fmt.Sprintf("escpos.Charset(%s)", c)
}

// GoString implements fmt.GoStringer
func (c CodePage) GoString() string {
	return fmt.Sprintf("escpos.CodePage(%s)", c)
}

// Select character size.
//
// Note:
//
//   - If the setting is outside of the defined range, this command is ignored.
//   - In standard mode, the vertical direction is the feed direction.
//   - In page mode, usage of the feed direction as either the vertical or
//     horizontal direction determined by the start position of the printable
//     area.
func SetCharacterSize(w, h uint8) Command {
	return setCharacterSize{w, h}
}

// WriteTo implements Command.
func (cmd setCharacterSize) WriteTo(w io.Writer) (int64, error) {
	if cmd.W > 7 || cmd.H > 7 {
		return 0, fmt.Errorf("invalid width or height: (%d, %d)", cmd.W, cmd.H)
	}

	return write(w, gs, '!', byte(cmd.W<<4)|byte(cmd.H&0x0F))
}

// Set right-side character spacing.
//
// Sets the character spacing for the right side of the character to
// (kerning * horizontal or vertical motion units).
//
// Note:
//
//   - The right-side character spacing for double-width mode is twice the normal
//     value.
//   - In standard mode, the horizontal motion unit is used.
//   - In page mode, usage of either the horizontal or vertical motion unit is
//     determined by the start position of the printable area.
func SetKerning(kerning uint8) Command {
	return setKerning{kerning}
}

// WriteTo implements Command.
func (cmd setKerning) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, ' ', byte(cmd.kerning))
}

// Select an international character set.
func SetCharset(charset Charset) Command {
	return setCharset{charset}
}

// WriteTo implements Command.
func (cmd setCharset) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, 'R', byte(cmd.charset))
}

// Select character font.
//
// Note:
//
//   - Character font can also be set by using [SetPrintMode]. However, the
//     setting of the last received command is effective.
func SetFont(font Font) Command {
	return setFont{font}
}

// WriteTo implements Command.
func (cmd setFont) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, 'M', byte(cmd.font))
}

// Selects character code table.
func SetCodePage(page CodePage) Command {
	return setCodePage{page}
}

// WriteTo implements Command.
func (cmd setCodePage) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, 't', byte(cmd.codePage))
}
