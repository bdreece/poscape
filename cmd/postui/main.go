package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/bdreece/poscape/internal/tui"
	"github.com/bdreece/poscape/pkg/escpos"
)

var devicePath = flag.String("p", "", "device path")

func main() {
    defer quit()
    flag.Parse()

    driver, err := escpos.OpenDriver(*devicePath)
    if err != nil {
        panic(err)
    }

    defer driver.(io.Closer).Close()
    if err = driver.Initialize(); err != nil {
        panic(err)
    }

    if err = tui.New(driver).Run(); err != nil {
        panic(err)
    }
}

func quit() {
    if r := recover(); r != nil {
        fmt.Fprintf(os.Stderr, "unexpected panic occurred: %v", r)
        os.Exit(1)
    }
}
