package world

import (
	"bufio"
	"fmt"
	"os"
)

type Canvas struct {
	Width  int
	Height int
	Color  [][]Color // represents the colors of each pixel
}

/*
Canvas looks like this

	  (0,0) -------------> (width) (X axis)
			|
			|
			|
			|
			|
		    \/
		   (height) (Y axis)
*/
func NewCanvas(width int, height int, color Color) *Canvas {
	canvas := &Canvas{}
	canvas.Height = height
	canvas.Width = width

	// Create the color slice
	// The X-axis corresponds to the outer slice of canvas.Color, representing the
	// width of the canvas. Each element in this outer slice is an inner slice
	// (subarray) that represents the Y-axis, or the height of the canvas.
	canvas.Color = make([][]Color, canvas.Width)
	for x := range canvas.Color {
		canvas.Color[x] = make([]Color, canvas.Height)
	}

	// the default value of float32 is 0.0
	// hence no need to initialze it with 0's
	if color.IsEqual(*Black) {
		return canvas
	}

	for x := 0; x < canvas.Width; x++ {
		for y := 0; y < canvas.Height; y++ {
			canvas.Color[x][y] = color
		}
	}

	return canvas
}

func (c *Canvas) WritePixel(x int, y int, color Color) {
	c.Color[x][y] = color
}

func (c *Canvas) PixelAt(x int, y int) Color {
	return c.Color[x][y]
}

func (c *Canvas) WriteToPPM(fileName string) error {

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	//Write the PPM header
	_, err = w.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", c.Width, c.Height))
	if err != nil {
		return err
	}

	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			color := c.Color[x][y].clamp()
			r := int(color.r * 255)
			g := int(color.g * 255)
			b := int(color.b * 255)
			_, err = w.WriteString(fmt.Sprintf("%d %d %d\n", r, g, b))
			if err != nil {
				return err
			}
		}
	}

	// Flush the writer
	err = w.Flush()
	if err != nil {
		return err
	}

	return nil
}
