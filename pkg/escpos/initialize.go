package escpos

import "io"

type (
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
	Initialize struct{}

	// Enables the printer after being disabled by [DisablePrinter].
	EnablePrinter struct{}

	// Disables the printer.
	//
	// When the printer is disabled, it ignores all data except for error-recovery
	// commands until it is re-enabled.
	DisablePrinter struct{}

	// Enables the feed button.
	EnableButton struct{}

	// Disables the feed button.
	DisableButton struct{}

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
	EnterPageMode struct{}

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
	EnterStandardMode struct{}
)

// WriteTo implements Command.
func (Initialize) WriteTo(w io.Writer) (int64, error) { return write(w, esc, '@') }

// WriteTo implements Command.
func (EnablePrinter) WriteTo(w io.Writer) (int64, error) { return write(w, esc, '=', 1) }

// WriteTo implements Command.
func (DisablePrinter) WriteTo(w io.Writer) (int64, error) { return write(w, esc, '=', 0) }

// WriteTo implements Command.
func (EnableButton) WriteTo(w io.Writer) (int64, error) { return write(w, esc, 'c', '5', 1) }

// WriteTo implements Command.
func (DisableButton) WriteTo(w io.Writer) (int64, error) { return write(w, esc, 'c', '5', 0) }

// WriteTo implements Command.
func (EnterPageMode) WriteTo(w io.Writer) (int64, error) { return write(w, esc, 'L') }

// WriteTo implements Command.
func (EnterStandardMode) WriteTo(w io.Writer) (int64, error) { return write(w, esc, 'S') }
