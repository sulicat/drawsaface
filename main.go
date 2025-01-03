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
	is_clicked_2 := false

	is_checked := false

	test_float := 500.123
	test_float2 := 0.000

	text_input := "hello world"

	for {
		t.ResetColorscheme()

		t.Label("asd")

		t.Label("gello")
		t.SameLine()
		t.Label("Hello wolrd")

		t.Label("gello2")

		if t.Button("ClickMe") {
			is_clicked = !is_clicked
		}

		if is_clicked {
			t.SameLine()
			if t.Button("Button Clicked") {
				is_clicked_2 = !is_clicked_2
			}

			t.Label("Button Clicked")
		}

		if is_clicked_2 {
			t.Label("SHOLD SEE INSIDE")
			t.InputFloat(&test_float2)
		}

		// checkbox
		t.CheckBox("Check me", &is_checked)
		t.CheckBox("Check me2", &is_checked)
		t.CheckBox("Check me3", &is_checked)
		t.CheckBox("Check me4", &is_checked)

		t.InputFloat(&test_float)
		t.Slider(&test_float, 0, 1000)

		t.InputText(&text_input, 50, 5)

		t.Step()
	}
}
