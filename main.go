package main

import (
	"fmt"
	"os"

	col "github.com/sulicat/goboi/colors"
	term "github.com/sulicat/goboi/term"
)

func usage() {
	fmt.Printf(col.Yellow + "Drawsaface" + col.Reset + " usage:\n")
	fmt.Printf(col.Red)
	fmt.Printf("\t drawsaface <input_file>\n")
	fmt.Printf("\t    <input_file> could be: .png\n")
	fmt.Printf(col.Reset)
}

func main() {
	// expect 1 input file
	if len(os.Args) < 2 {
		usage()
		return
	}

	t := term.Create(50, 50)
	// t.SetFullscreen(false)
	// test_float := 123.0

	for {
		// this will

		t.SameLine()
		t.Label("asd")
		t.SetColor(term.RGB{0, 255, 0})
		t.Label("gello")
		t.Label("Some Random info")
		t.Label("These two are on the same line")
		t.SetColor(term.RGB{0, 0, 255})
		t.Label("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabcdefghi")
		// t.InputFloat(&test_float)

		t.Step()
	}
}
