package main

import (
	"math"

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

	floats := make([]float64, 10)
	offset := 0.0
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

		for i := range 3 {
			t.InputFloat(&floats[i])
		}

		s := float64(t.GetScroll())
		t.InputFloat(&s)
		t.Slider(&test_float, 0, 1000)

		t.InputText(&text_input, 50, 5)

		pixels := t.CreatePixels(40, 40)
		offset += 0.005
		for i := range 360 {
			ang := float64(i) * math.Pi / 180.0
			x := math.Cos(ang) * 10
			y := math.Sin(ang) * 1 * float64(int(offset)%20-10)
			x += 15
			y += 15
			// pixels[int(x)][int(y)] = term.RGB{255, int((float64((i+int(offset))%360) / 360.0) * 255), 0}
			pixels[int(x)][int(y)] = term.RGB{255, 255, 0}
		}

		t.Canvas(&pixels)

		t.Step()

	}
}
