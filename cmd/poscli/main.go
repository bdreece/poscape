package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/bdreece/poscape/pkg/escpos"
)

var (
    marginTop = flag.Int("mt", 0, "top margin (lines)")
    marginBottom = flag.Int("mb", 0, "bottom margin (lines)")
)

func main() {
    defer quit()
    flag.Parse()
    driver := escpos.NewDriver(os.Stdout)
    if err := driver.Initialize(); err != nil {
        panic(err)
    }

    if *marginTop != 0 {
        if err := driver.PrintAndFeed(uint8(*marginTop)); err != nil {
            panic(err)
        }
    }

    if _, err := io.Copy(driver, os.Stdin); err != nil {
        panic(err)
    }

    if *marginBottom != 0 {
        if err := driver.PrintAndFeed(uint8(*marginBottom)); err != nil {
            panic(err)
        }
    }
}

func quit() {
    if r := recover(); r != nil {
        fmt.Fprintf(os.Stderr, "unexpected panic occurred: %v", r)
        os.Exit(1)
    }
}
