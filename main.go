package main

import (
	"fmt"
	"os"

	"github.com/sulicat/drawsaface/drawsaface"
	col "github.com/sulicat/goboi/colors"
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

	daf := drawsaface.Create(os.Args[1:])
	daf.Draw()

}
