package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"rsc.io/quote"
)

/* ------------- Point --------------- */
type Point struct {
	x, y, z float64
}

func NewPoint(x float64, y float64, z float64) *Point {
	return &Point{x, y, z}
}

func (p1 Point) IsEqual(p2 Point) bool {
	return isFloatEqual(p1.x, p2.x) &&
		isFloatEqual(p1.y, p2.y) &&
		isFloatEqual(p1.z, p2.z)
}

// Add a vector to a point, gives a new point
// This is equivalent to walking from a point in the direction of the vector
func (p1 Point) AddVector(v1 Vector) *Point {
	return &Point{p1.x + v1.x, p1.y + v1.y, p1.z + v1.z}
}

// subtracting two points finds the vector between the points
func (p1 Point) Subtract(p2 Point) *Vector {
	return &Vector{p1.x - p2.x, p1.y - p2.y, p1.z - p2.z}
}

// Subtract a vector from a point, gives a new point
// This is equivalent to walking from a point in the direction of the vector
func (p1 Point) SubtractVector(v1 Vector) *Point {
	return &Point{p1.x - v1.x, p1.y - v1.y, p1.z - v1.z}
}

func (p1 Point) Negate() *Point {
	return &Point{-p1.x, -p1.y, -p1.z}
}

func (p1 Point) ScalarMultiply(scalar float64) *Point {
	return &Point{p1.x * scalar, p1.y * scalar, p1.z * scalar}
}

func (p1 Point) ScalarDivide(scalar float64) *Point {
	return &Point{p1.x / scalar, p1.y / scalar, p1.z / scalar}
}

/* ------------- Vector --------------- */
type Vector struct {
	x, y, z float64
}

func NewVector(x float64, y float64, z float64) *Vector {
	return &Vector{x, y, z}
}

func (v1 Vector) IsEqual(v2 Vector) bool {
	return isFloatEqual(v1.x, v2.x) &&
		isFloatEqual(v1.y, v2.y) &&
		isFloatEqual(v1.z, v2.z)
}

func (v1 Vector) add(v2 Vector) *Vector {
	return &Vector{v1.x + v2.x, v1.y + v2.y, v1.z + v2.z}
}

func AddVectors(vlist []Vector) *Vector {
	vectorSum := Vector{}

	for _, vec := range vlist {
		vectorSum = *vectorSum.add(vec)
	}

	return &vectorSum

}

func (v1 Vector) subtract(v2 Vector) *Vector {
	return &Vector{v1.x - v2.x, v1.y - v2.y, v1.z - v2.z}
}

// The vectors are subtracted in the order they are passed
// For eg: v3 = v1 - v2 is denoted by SubtractVector(v1, v2)
func SubtractVectors(vlist []Vector) *Vector {
	result := vlist[0]
	for i := 1; i < len(vlist); i++ {
		result = *result.subtract(vlist[i])
	}

	return &result
}

func (v1 Vector) Negate() *Vector {
	return &Vector{-v1.x, -v1.y, -v1.z}
}

func (v1 Vector) ScalarMultiply(scalar float64) *Vector {
	return &Vector{v1.x * scalar, v1.y * scalar, v1.z * scalar}
}

func (v1 Vector) ScalarDivide(scalar float64) *Vector {
	return &Vector{v1.x / scalar, v1.y / scalar, v1.z / scalar}
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

// Converts an arbitrary vector into a unit vector.
// This keep calculations anchored relative to a common scale (the unit vector)
func (v1 Vector) Normalize() *Vector {
	v1_magnitude := v1.Magnitude()
	normalized_x := v1.x / v1_magnitude
	normalized_y := v1.y / v1_magnitude
	normalized_z := v1.z / v1_magnitude
	return &Vector{normalized_x, normalized_y, normalized_z}
}

// Calculates the dot product of vector
// one use case, dot products of unit vectors help find the angle between vectors
func (v1 Vector) DotProduct(v2 Vector) float64 {
	return (v1.x*v2.x + v1.y*v2.y + v1.z*v2.z)
}

// Cross product
// cross product of X and Y gets Z, Y and Z get X, i.e results are always perpendicular
func (v1 Vector) CrossProduct(v2 Vector) *Vector {
	crossProduct_x := v1.y*v2.z - v1.z*v2.y
	crossProduct_y := v1.z*v2.x - v1.x*v2.z
	crossProduct_z := v1.x*v2.y - v1.y*v2.x

	return &Vector{crossProduct_x, crossProduct_y, crossProduct_z}
}

func main() {
	fmt.Println(quote.Glass())

	// Projectile
	p := Projectile{position: Point{0, 1, 0}, velocity: Vector{1, 1, 0}}
	e := Environment{gravity: Vector{0, -0.1, 0}, wind: Vector{-0.01, 0, 0}}
	finalProjection := ThrowProjectile(e, p)
	fmt.Printf("\nfinal projectile position: %v+\n", finalProjection.position)
}

/* ------------- Colors --------------- */

type Color struct {
	r, g, b float64
}

var Red = NewColor(1.0, 0.0, 0.0)
var Green = NewColor(0.0, 1.0, 0.0)
var Blue = NewColor(0.0, 0.0, 1.0)
var Black = NewColor(0.0, 0.0, 0.0)

func NewColor(r float64, g float64, b float64) *Color {
	return &Color{r, g, b}
}

func (c1 Color) IsEqual(c2 Color) bool {
	return isFloatEqual(c1.r, c2.r) &&
		isFloatEqual(c1.g, c2.g) &&
		isFloatEqual(c1.b, c2.b)
}

func (c *Color) clamp() *Color {
	c.r = Clamp(c.r)
	c.g = Clamp(c.g)
	c.b = Clamp(c.b)

	return c
}

func (c1 Color) add(c2 Color) *Color {
	return &Color{c1.r + c2.r, c1.g + c2.g, c1.b + c2.b}
}

func AddColors(vlist []Color) *Color {
	ColorSum := Color{}

	for _, vec := range vlist {
		ColorSum = *ColorSum.add(vec)
	}

	return &ColorSum

}

func (c1 Color) subtract(c2 Color) *Color {
	return &Color{c1.r - c2.r, c1.g - c2.g, c1.b - c2.b}
}

// The Colors are subtracted in the order they are passed
// For eg: v3 = v1 - v2 is denoted by SubtractColor(v1, v2)
func SubtractColors(vlist []Color) *Color {
	result := vlist[0]
	for i := 1; i < len(vlist); i++ {
		result = *result.subtract(vlist[i])
	}

	return &result
}

func (c1 Color) multiply(c2 Color) *Color {
	return &Color{c1.r * c2.r, c1.g * c2.g, c1.b * c2.b}
}

// The Colors are multiplied in the order they are passed
// Also called as "hadamard_product"
// For eg: v3 = v1 * v2 is denoted by MultiplyColors(v1, v2)
func MultiplyColors(vlist []Color) *Color {
	result := vlist[0]
	for i := 1; i < len(vlist); i++ {
		result = *result.multiply(vlist[i])
	}

	return &result
}

/* ------------- Canvas --------------- */

type Canvas struct {
	width  int
	height int
	color  [][]Color // represents the colors of each pixel
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
	canvas.height = height
	canvas.width = width

	// Create the color slice
	// The X-axis corresponds to the outer slice of canvas.color, representing the
	// width of the canvas. Each element in this outer slice is an inner slice
	// (subarray) that represents the Y-axis, or the height of the canvas.
	canvas.color = make([][]Color, canvas.width)
	for x := range canvas.color {
		canvas.color[x] = make([]Color, canvas.height)
	}

	// the default value of float32 is 0.0
	// hence no need to initialze it with 0's
	if color.IsEqual(*Black) {
		return canvas
	}

	for x := 0; x < canvas.width; x++ {
		for y := 0; y < canvas.height; y++ {
			canvas.color[x][y] = color
		}
	}

	return canvas
}

func (c *Canvas) WritePixel(x int, y int, color Color) {
	c.color[x][y] = color
}

func (c *Canvas) PixelAt(x int, y int) Color {
	return c.color[x][y]
}

func (c *Canvas) WriteToPPM(fileName string) error {

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	//Write the PPM header
	_, err = w.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", c.width, c.height))
	if err != nil {
		return err
	}

	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			color := c.color[x][y].clamp()
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
