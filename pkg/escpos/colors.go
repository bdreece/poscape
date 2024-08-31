package escpos

import "io"

// Turn white/black reverse printing mode on/off.
//
// Note:
//
//   - This command is available for built-in characters and user-defined
//     characters.
//   - When white/black reverse printing mode is on, it is also applied to
//     character spacing set by [SetCharacterKerning].
//   - This command does not affect the line spacing set by [SetLineSpacing]
//     (or [ResetLineSpacing]).
//   - White/black reverse printing mode has a higher priority than underline
//     mode. Even if underline mode is on, it is disabled (but not cancelled)
//     while white/black reverse mode is on.
type SetInverseColors struct {
    Enabled bool
}

// WriteTo implements Command.
func (cmd SetInverseColors) WriteTo(w io.Writer) (int64, error) {
    return write(w, gs, 'B', bit(cmd.Enabled))
}