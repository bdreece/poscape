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

	// Select character size.
	//
	// Note:
	//
	//   - If the setting is outside of the defined range, this command is ignored.
	//   - In standard mode, the vertical direction is the feed direction.
	//   - In page mode, usage of the feed direction as either the vertical or
	//     horizontal direction determined by the start position of the printable
	//     area.
	SetCharacterSize struct {
		W, H uint8
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
	SetKerning struct {
		Kerning uint8
	}

	// Select an international character set.
	SetCharset struct {
		Charset Charset
	}

	// Select character font.
	//
	// Note:
	//
	//   - Character font can also be set by using [SetPrintMode]. However, the
	//     setting of the last received command is effective.
	SetFont struct {
		Font Font
	}

	// Selects character code table.
	SetCodePage struct {
		Page CodePage
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

// WriteTo implements Command.
func (cmd SetCharacterSize) WriteTo(w io.Writer) (int64, error) {
	if cmd.W > 7 || cmd.H > 7 {
		return 0, fmt.Errorf("invalid width or height: (%d, %d)", cmd.W, cmd.H)
	}

	return write(w, gs, '!', byte(cmd.W<<4)|byte(cmd.H&0x0F))
}

// WriteTo implements Command.
func (cmd SetKerning) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, ' ', byte(cmd.Kerning))
}

// WriteTo implements Command.
func (cmd SetCharset) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, 'R', byte(cmd.Charset))
}

// WriteTo implements Command.
func (cmd SetFont) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, 'M', byte(cmd.Font))
}

// WriteTo implements Command.
func (cmd SetCodePage) WriteTo(w io.Writer) (int64, error) {
	return write(w, esc, 't', byte(cmd.Page))
}
