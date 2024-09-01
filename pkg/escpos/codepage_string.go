// Code generated by "stringer -type CodePage -linecomment"; DO NOT EDIT.

package escpos

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CodePagePC437-15]
	_ = x[CodePageKatakana-16]
	_ = x[CodePagePC850-17]
	_ = x[CodePagePC860-18]
	_ = x[CodePagePC863-19]
	_ = x[CodePagePC865-20]
	_ = x[CodePagePC866-17]
	_ = x[CodePageSpace-255]
}

const (
	_CodePage_name_0 = "PC437 [U.S.A., Standard Europe]KatakanaPC850 [Multilingual]PC860 [Portuguese]PC863 [Canadian-French]PC865 [Nordic]"
	_CodePage_name_1 = "Space page"
)

var (
	_CodePage_index_0 = [...]uint8{0, 31, 39, 59, 77, 100, 114}
)

func (i CodePage) String() string {
	switch {
	case 15 <= i && i <= 20:
		i -= 15
		return _CodePage_name_0[_CodePage_index_0[i]:_CodePage_index_0[i+1]]
	case i == 255:
		return _CodePage_name_1
	default:
		return "CodePage(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}