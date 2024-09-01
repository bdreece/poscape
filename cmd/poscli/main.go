package main

import (
	"fmt"
	"os"

	"github.com/bdreece/poscape/pkg/escpos"
)

func main() {
	cmds := []escpos.Command{
		escpos.Initialize(),
		escpos.PrintAndFeedLines(4),
		escpos.SetUnderline(escpos.UnderlineSingle),
		escpos.Text("Hello, world!"),
		escpos.SetUnderline(escpos.UnderlineNone),
		escpos.LineFeed(),
		escpos.SetLeftMargin(2),
		escpos.Text("- Item 1"),
		escpos.LineFeed(),
		escpos.Text("- Item 1"),
		escpos.LineFeed(),
		escpos.Text("- Item 1"),
		escpos.LineFeed(),
		escpos.PrintAndFeedLines(4),
	}

	for _, cmd := range cmds {
		fmt.Printf("%#v\n", cmd)
	}

	if _, err := escpos.Write(os.Stdout, cmds...); err != nil {
		panic(err)
	}
}
