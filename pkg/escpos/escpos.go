package escpos

import (
	"fmt"
	"io"
)

type (
    // Command provides the contract for data sent to an ESC/POS-compliant printer.
	Command interface {
		io.WriterTo
	}

    // Text represents plain, writable ASCII text sent as a Command.
	Text string

    Raw []byte
)

const (
	// Horizontal-tab character
	ht byte = 0x09
	// Form-feed character
	ff byte = 0x0C
	// Escape character
	esc byte = 0x1B
	// File-separator character
	fs byte = 0x1C
	// Group-separator character
	gs byte = 0x1D
)

// GoString implements fmt.GoStringer.
func (txt Text) GoString() string { return fmt.Sprintf("Text: %q", string(txt)) }

// WriteTo implements Command.
func (txt Text) WriteTo(w io.Writer) (int64, error) { return write(w, []byte(txt)...) }

// GoString implements fmt.GoStringer.
func (raw Raw) GoString() string { return fmt.Sprintf("Raw: %X", string(raw)) }

// WriteTo implements Command.
func (raw Raw) WriteTo(w io.Writer) (int64, error) { return write(w, []byte(raw)...) }

func Write(w io.Writer, cmds ...Command) (total int64, err error) {
    var n int64
    for _, cmd := range cmds {
        n, err = cmd.WriteTo(w)
        total += n
        if err != nil {
            return
        }
    }

    return
}

func write(w io.Writer, data ...byte) (int64, error) {
	n, err := w.Write(data)
	return int64(n), err
}

func bit(val bool) byte {
    if val {
        return 1
    } else {
        return 0
    }
}
