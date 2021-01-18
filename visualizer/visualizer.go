package visualizer

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
	"time"
)

var Fps int
var Mode int
var MaxHeight int
var test bool = false

type Visualizer interface {
	Setup(string)
	AddFrame([]int)
	Complete()
}

func (fg FrameGen) Setup(name string) {
}

func (fg FrameGen) AddFrame(array []int) {
	fg(array)
}

func (fg FrameGen) Complete() {
}

type FrameGen func([]int)

type GifVisualizer struct {
	name    string
	gifdata *gif.GIF
}

func (gv *GifVisualizer) Setup(name string) {
	gv.gifdata = &gif.GIF{
		LoopCount: 1,
	}

	gv.name = name
}

func (gv *GifVisualizer) AddFrame(array []int) {
	frame := buildImage(array)
	gv.gifdata.Image = append(gv.gifdata.Image, frame)
	gv.gifdata.Delay = append(gv.gifdata.Delay, 2)
}

func (gv *GifVisualizer) Complete() {
	WriteGif(gv.name, gv.gifdata)
}

func WriteStdout(array []int) {
	var buffer bytes.Buffer

	for y := 0; y < MaxHeight; y++ {
		for x := 0; x < len(array); x++ {
			if array[x] == y {
				buffer.WriteByte(byte('#'))
			} else if array[x] < y && Mode == 1 {
				buffer.WriteByte(byte('#'))
			} else if array[x] > y && Mode == 2 {
				buffer.WriteByte(byte('#'))
			} else {
				buffer.WriteByte(byte(' '))
			}
		}
		buffer.WriteByte('\n')
	}

	if !test {
		time.Sleep(time.Second / time.Duration(Fps))
		fmt.Print("\033[2J")
		fmt.Print(buffer.String())
	}
}

func WriteGif(name string, gifdata *gif.GIF) {
	file, err := os.Create(name + ".gif")
	if err != nil {
		fmt.Println("os.Create")
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("file.Close()")
			panic(err)
		}
	}()

	err = gif.EncodeAll(file, gifdata)
	if err != nil {
		fmt.Println("gif.EncodeAll")
		panic(err)
	}
}

func buildImage(array []int) *image.Paletted {
	frame := image.NewPaletted(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{len(array), MaxHeight},
		},
		color.Palette{
			color.Gray{uint8(255)},
			color.Gray{uint8(0)},
		},
	)

	for index, value := range array {
		frame.SetColorIndex(index, MaxHeight-value, uint8(1))
		if Mode == 2 {
			for y := MaxHeight - value + 1; y < MaxHeight; y++ {
				frame.SetColorIndex(index, y, uint8(1))
			}
		}
	}

	return frame
}
