//Converts an array on ingeres as a graphic
//for stdout or gif representation
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

//region Members

var Fps int
var Mode int
var MaxHeight int
var test bool = false

type FrameGen func([]int, int)

//endregion

//region Visualizer interface

type Visualizer interface {
	Setup(string)
	AddFrame([]int, int)
	Complete()
}

func (fg FrameGen) Setup(name string) {
}

func (fg FrameGen) AddFrame(array []int, currentValue int) {
	fg(array, currentValue)
}

func (fg FrameGen) Complete() {
}

//endregion

//region Gif Visualizer

//Represents name of an algorithm and GIF data of the algorithm
type GifVisualizer struct {
	name    string
	gifdata *gif.GIF
}

//Setups base config of gif image
func (gv *GifVisualizer) Setup(name string) {
	gv.gifdata = &gif.GIF{
		LoopCount: 0, //infinite
	}

	gv.name = name
}

//Builds the frame and added in to a GIF data
func (gv *GifVisualizer) AddFrame(array []int, currentValue int) {
	frame := buildImage(array, currentValue)
	gv.gifdata.Image = append(gv.gifdata.Image, frame)
	gv.gifdata.Delay = append(gv.gifdata.Delay, 0)
}

//Writes GIF image to the filesystem when alorithm is done
func (gv *GifVisualizer) Complete() {
	WriteGif(gv.name, gv.gifdata)
}

//endregion

//region Public Methods

//Represents the array as chart on the screen
func WriteStdout(array []int, currentValue int) {
	var buffer bytes.Buffer

	for y := 0; y < MaxHeight; y++ {
		for x := 0; x < len(array); x++ {
			if array[x] == y {
				buffer.WriteByte(byte('^')) //highest value in column
			} else if array[x] < y && Mode == 1 { //symbols under the value
				buffer.WriteByte(byte('*'))
			} else if array[x] > y && Mode == 2 { //symbols above the value
				buffer.WriteByte(byte('*'))
			} else {
				buffer.WriteByte(byte(' ')) //free space
			}
		}
		buffer.WriteByte('\n')
	}

	if !test {
		time.Sleep(time.Second / time.Duration(Fps))
		fmt.Print("\033[2J") //clear screen
		fmt.Print(buffer.String())
	}
}

//Writes GIF image to the filesystem
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

//endregion

//region Private Methods

func buildImage(array []int, currentValue int) *image.Paletted {
	frame := image.NewPaletted(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{len(array), MaxHeight},
		},
		color.Palette{
			color.Gray{uint8(255)},                                   //0 - backgound
			color.RGBA{uint8(53), uint8(183), uint8(219), uint8(1)},  //1 - values
			color.RGBA{uint8(166), uint8(231), uint8(255), uint8(1)}, //2 - value's backgound
			color.RGBA{uint8(250), uint8(178), uint8(35), uint8(1)},  //3 - current value
			color.RGBA{uint8(252), uint8(219), uint8(98), uint8(1)},  //4 - current value's backgound

		},
	)

	for x, value := range array {

		if x == currentValue {
			frame.SetColorIndex(x, MaxHeight-value, uint8(3))
		} else {
			frame.SetColorIndex(x, MaxHeight-value, uint8(1))
		}

		if Mode == 2 {
			for y := MaxHeight - value + 1; y < MaxHeight; y++ {
				if x == currentValue {
					frame.SetColorIndex(x, y, uint8(4))
				} else {
					frame.SetColorIndex(x, y, uint8(2))
				}
			}
		}
	}

	return frame
}

//endregion
