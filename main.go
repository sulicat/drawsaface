package main

import (
	term "github.com/sulicat/goboi/term"
)

// func usage() {
// 	fmt.Printf(col.Yellow + "Drawsaface" + col.Reset + " usage:\n")
// 	fmt.Printf(col.Red)
// 	fmt.Printf("\t drawsaface <input_file>\n")
// 	fmt.Printf("\t    <input_file> could be: .png\n")
// 	fmt.Printf(col.Reset)
// }

func main() {

	t := term.Create(50, 50)
	// t.SetFullscreen(false)
	// test_float := 123.0

	is_clicked := false

	for {
		// this will

		t.SetColor(term.RGB{255, 0, 255})
		t.Label("asd")

		t.SameLine()

		t.SetColor(term.RGB{0, 255, 0})
		t.Label("gello")

		t.SameLine()

		t.SetColor(term.RGB{255, 255, 255})
		t.Label("Hello wolrd")

		t.Label("gello2")

		if t.Button("ClickMe") {
			is_clicked = !is_clicked
		}

		if is_clicked {
			t.SetColor(term.RGB{255, 0, 0})
			t.Label("Button Clicked")
		}

		// t.InputFloat(&test_float)

		t.Step()
	}
}
