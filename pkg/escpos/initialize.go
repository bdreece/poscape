package escpos

import "io"

type (
	initialize        struct{}
	enablePrinter     struct{}
	disablePrinter    struct{}
	enableButton      struct{}
	disableButton     struct{}
	enterPageMode     struct{}
	enterStandardMode struct{}
)

// WriteTo implements Command.
func (initialize) WriteTo(w io.Writer) (int64, error) { return write(w, esc, '@') }

// Initialize printer.
//
// Clears the data in the print buffer and resets the print mode to the mode
// that was in effect when the printer was powered on.
//
// Note:
//
//   - The DIP switch settings are not checked again.
//   - The data in the receive buffer is not cleared.
//   - The macro definition is not cleared.
//   - The NV bit image data is not cleared.
//   - The data of the NV user memory is not cleared.
func Initialize() Command { return *new(initialize) }

// WriteTo implements Command.
func (enablePrinter) WriteTo(w io.Writer) (int64, error) { return write(w, esc, '=', 1) }

// Enables the printer after being disabled by [DisablePrinter].
func EnablePrinter() Command { return *new(enablePrinter) }

// WriteTo implements Command.
func (disablePrinter) WriteTo(w io.Writer) (int64, error) { return write(w, esc, '=', 0) }

// Disables the printer.
//
// When the printer is disabled, it ignores all data except for error-recovery
// commands until it is re-enabled.
func DisablePrinter() Command { return *new(disablePrinter) }

// WriteTo implements Command.
func (enableButton) WriteTo(w io.Writer) (int64, error) { return write(w, esc, 'c', '5', 1) }

// Enables the feed button.
func EnableButton() Command { return *new(enablePrinter) }

// WriteTo implements Command.
func (disableButton) WriteTo(w io.Writer) (int64, error) { return write(w, esc, 'c', '5', 0) }

// Disables the feed button.
func DisableButton() Command { return *new(disablePrinter) }

// WriteTo implements Command.
func (enterPageMode) WriteTo(w io.Writer) (int64, error) { return write(w, esc, 'L') }

// Select page mode.
//
// Switches from standard mode to page mode.
//
// Note:
//
//   - This command has no effect in page mode.
//   - This command is enabled only when input at the beginning of a line in
//     stadard mode.
//   - The values set by [SetCharacterKerning] and [SetLineSpacing] (or
//     [ResetLineSpacing]) are scoped to the current mode.
func EnterPageMode() Command { return *new(enterPageMode) }

// WriteTo implements Command.
func (enterStandardMode) WriteTo(w io.Writer) (int64, error) { return write(w, esc, 'S') }

// Select standard mode.
//
// Switches from page mode to standard mode.
//
// Note:
//
//   - This command is only effective in page mode.
//   - Data buffered in page mode and the printable area developed in page mode
//     are cleared.
//   - The values set by [SetCharacterKerning] and [SetLineSpacing] (or
//     [ResetLineSpacing]) are scoped to the current mode.
func EnterStandardMode() Command { return *new(enterStandardMode) }
