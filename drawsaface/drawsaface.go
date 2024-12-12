package drawsaface

import (
	"bufio"
	"fmt"
	"image"
	"math"
	"os"
	"strings"
	"time"

	"gocv.io/x/gocv"
	"golang.org/x/term"

	col "github.com/sulicat/goboi/colors"
)

type Frame image.Image

type Drawsaface struct {
	Width  int
	Height int
	Frames []gocv.Mat
	Writer *bufio.Writer
}

const char_stable = "x"
const angle_setting = 0
const stable_setting = 200

func Create(paths []string) *Drawsaface {
	daf := Drawsaface{Width: 100, Height: 100}
	daf.Frames = make([]gocv.Mat, 0)
	daf.Writer = bufio.NewWriter(os.Stdout)

	for _, s := range paths {
		daf.Load(s)
	}
	return &daf
}

func (daf *Drawsaface) Load(path string) {
	image := gocv.IMRead(path, gocv.IMReadColor)
	daf.Frames = append(daf.Frames, image)
}

func (daf *Drawsaface) Draw() {
	fmt.Print(col.Clear())

	i := 0
	for {
		//w, h, err := getTerminalSize()
		w, h, _ := term.GetSize(0)
		start := time.Now()
		DrawAsciiFrame(daf.Writer, daf.Frames[i], 0, 0, w, h)
		dur := time.Since(start)
		fmt.Fprintf(os.Stderr, "%v\n", dur)
		// TODO: proper framerate
		time.Sleep(100000000)
		i += 1
		i = i % len(daf.Frames)
	}
}

// -----------------------------------------------------------------------------------------------------------------

// var writer := bufio.NewWriter(os.Stdout)

func choose_char(sx, sy int) string {

	sx_abs := math.Abs(float64(sx))
	sy_abs := math.Abs(float64(sy))

	if sx_abs < stable_setting && sy_abs < stable_setting {
		return char_stable
	} else {

		if sx_abs-angle_setting < sy_abs {
			return "—"
		}

		if sy_abs-angle_setting < sx_abs {
			return "|"
		}

		if sx_abs > sy_abs {
			return "/"
		} else {
			return "\\"
		}

		if sy > sx {
			return "\\"
		}

	}
	return "."
}

func DrawAsciiFrame(w *bufio.Writer, f gocv.Mat, x, y, width, height int) {

	small_f := gocv.NewMat()
	defer small_f.Close()

	small_f_gray := gocv.NewMat()
	defer small_f_gray.Close()

	gocv.Resize(f, &small_f, image.Point{width, height}, float64(0), float64(0), gocv.InterpolationArea)
	gocv.CvtColor(small_f, &small_f_gray, gocv.ColorBGRToGray)

	// Initialize matrices to store the gradients
	gradX := gocv.NewMat()
	gradY := gocv.NewMat()
	defer gradX.Close()
	defer gradY.Close()

	gocv.Sobel(small_f_gray, &gradX, gocv.MatTypeCV16S, 1, 0, 3, 1, 0, gocv.BorderReflect101)
	gocv.Sobel(small_f_gray, &gradY, gocv.MatTypeCV16S, 0, 1, 3, 1, 0, gocv.BorderReflect101)

	var sb strings.Builder
	sb.Grow(width * height * 50)
	sb.WriteString(col.MoveCursor(0, 0))

	rows := small_f.Rows()
	cols := small_f.Cols()

	for x := 0; x < cols; x++ {
		for y := 0; y < rows; y++ {

			vec := small_f.GetVecbAt(y, x)
			r := vec[2]
			g := vec[1]
			b := vec[0]

			sb.WriteString(col.MoveCursor(x, y))

			sx := gradX.GetShortAt(y, x)
			sy := gradY.GetShortAt(y, x)

			char := choose_char(int(sx), int(sy))
			sb.WriteString(col.DrawChar(char, int(r), int(g), int(b)))

		}
	}

	buff := sb.String()
	fmt.Fprint(w, buff)
	w.Flush()
}
