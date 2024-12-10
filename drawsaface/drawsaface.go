package drawsaface

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/nfnt/resize"
	"golang.org/x/term"

	col "github.com/sulicat/goboi/colors"
	"github.com/sulicat/goboi/utils"
)

const SCALE_FACTOR_W2H = 2

type Frame image.Image

type Drawsaface struct {
	Width  int
	Height int
	Frames []Frame
}

func Create(paths []string) *Drawsaface {
	daf := Drawsaface{Width: 100, Height: 100}
	daf.Frames = make([]Frame, 0)

	for _, s := range paths {
		daf.Load(s)
	}
	return &daf
}

func (daf *Drawsaface) Load(path string) {

	switch utils.FileExtension(path) {
	case "png":
		daf.LoadPng(path)
	case "mp4":
		daf.LoadMp4()
	}
}

func (daf *Drawsaface) LoadMp4() {
	fmt.Printf(col.BgRed + "NOT IMPLEMENTED\n" + col.Reset)
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
		DrawAsciiFrame(daf.Frames[i], 0, 0, w, h)

		time.Sleep(1000000)
		i += 1
		i = i % len(daf.Frames)
	}
}

// -----------------------------------------------------------------------------------------------------------------

func DrawAsciiFrame(f Frame, x, y, width, height int) {
	//fmt.Printf("%v\n", f)

	new_image := resize.Resize(uint(width), uint(height), f, resize.NearestNeighbor)

	buff := ""
	buff += col.MoveCursor(0, 0)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			buff += col.MoveCursor(x, y)
			r, g, b, a := new_image.At(x, y).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8
			a = a >> 8

			if a < 200 {
				buff += col.DrawBlank()
			} else {
				//buff += col.DrawBlock(int(r), int(g), int(b))
				buff += col.DrawChar("o", int(r), int(g), int(b))
			}
		}
	}

	fmt.Print(buff)

}
