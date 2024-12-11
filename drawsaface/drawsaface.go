package drawsaface

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"
	"time"

	"github.com/nfnt/resize"
	"golang.org/x/term"

	col "github.com/sulicat/goboi/colors"
	"github.com/sulicat/goboi/utils"
)

type Frame image.Image

type Drawsaface struct {
	Width  int
	Height int
	Frames []Frame
	Writer *bufio.Writer
}

func Create(paths []string) *Drawsaface {
	daf := Drawsaface{Width: 100, Height: 100}
	daf.Frames = make([]Frame, 0)
	daf.Writer = bufio.NewWriter(os.Stdout)

	for _, s := range paths {
		daf.Load(s)
	}
	return &daf
}

func (daf *Drawsaface) Load(path string) {

	switch utils.FileExtension(path) {
	case "png":
		daf.LoadPng(path)
	}
}

func (daf *Drawsaface) LoadPng(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf(col.BgBrightRed+"ERROR:"+col.Reset+" %s\n", err)
		return
	}

	defer file.Close()

	image, err := png.Decode(file)
	if err != nil {
		fmt.Printf(col.BgBrightRed+"ERROR:"+col.Reset+" %s\n", err)
		return
	}

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

func DrawAsciiFrame(w *bufio.Writer, f Frame, x, y, width, height int) {
	new_image := resize.Resize(uint(width), uint(height), f, resize.NearestNeighbor)

	var sb strings.Builder
	sb.Grow(width * height * 50)

	//buff := ""
	sb.WriteString(col.MoveCursor(0, 0))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			r, g, b, a := new_image.At(x, y).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8
			a = a >> 8

			sb.WriteString(col.MoveCursor(x, y))

			if a < 200 {
				sb.WriteString(col.DrawBlank())
			} else {
				sb.WriteString(col.DrawBlock(int(r), int(g), int(b)))
				//sb.WriteString(col.DrawChar("x", int(r), int(g), int(b)))
			}

		}
	}

	buff := sb.String()

	fmt.Fprint(w, buff)
	w.Flush()
}
