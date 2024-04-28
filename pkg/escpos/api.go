package escpos

import (
	"fmt"
	"io"
	"io/fs"
	"os"
)

const esc byte = 0x1B

type PrintMode uint8

const (
	AlternateFontMode PrintMode = 1
	EmphasisMode      PrintMode = 1 << 3
	DoubleHeightMode  PrintMode = 1 << 4
	DoubleWidthMode   PrintMode = 1 << 5
	UnderlineMode     PrintMode = 1 << 7
)

type Justification uint8

const (
	JustifyLeft   Justification = 0
	JustifyCenter Justification = 1
	JustifyRight  Justification = 2
)

type Driver interface {
	io.WriteCloser

	Initialize() error
	SelectPrintMode(mode PrintMode) error
	SelectLineSpacing(spacing uint8) error
	SelectJustification(justify Justification) error
	PrintAndFeed(lines uint8) error
}

type driver struct {
	io.WriteCloser
}

func (d driver) Initialize() error {
	cmd := []byte{esc, '@'}
	if _, err := d.Write(cmd); err != nil {
		return fmt.Errorf("failed to initialize driver: %v", err)
	}

	return nil
}

func (d driver) SelectPrintMode(mode PrintMode) error {
	cmd := []byte{esc, '!', byte(mode)}
	if _, err := d.Write(cmd); err != nil {
		return fmt.Errorf("failed to select print mode %d: %v", mode, err)
	}

	return nil
}

func (d driver) SelectLineSpacing(spacing uint8) error {
	cmd := []byte{esc, 3, spacing}
	if _, err := d.Write(cmd); err != nil {
		return fmt.Errorf("failed to select line spacing %d: %v", spacing, err)
	}

	return nil
}

func (d driver) SelectJustification(justify Justification) error {
	cmd := []byte{esc, 'a', byte(justify)}
	if _, err := d.Write(cmd); err != nil {
		return fmt.Errorf("failed to select justification %d: %v", justify, err)
	}

	return nil
}

func (d driver) PrintAndFeed(lines uint8) error {
	cmd := []byte{esc, 'd', lines}
	if _, err := d.Write(cmd); err != nil {
		return fmt.Errorf("failed to print and feed %d lines: %v", lines, err)
	}

	return nil
}

func NewDriver(w io.WriteCloser) Driver {
	return driver{w}
}

func OpenDriver(path string) (Driver, error) {
	const (
		flag int         = os.O_WRONLY | os.O_APPEND
		mode fs.FileMode = 0o0644
	)

	f, err := os.OpenFile(path, flag, mode)
	if err != nil {
		return nil, err
	}

	return driver{f}, nil
}
