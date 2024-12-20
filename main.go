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

	for {
		// this will
		t.Step()
	}
}
