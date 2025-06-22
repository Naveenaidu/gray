package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

/* ------------- Matrices --------------- */

type Matrix struct {
	rows    int
	columns int
	value   [][]float64
}

/*
Create a new matrix. Each element of "row" array defines the rows of a
matrix

To Creates a matrix like below:

	1.00   2.00   3.00   4.00
	5.50   6.50   7.50   8.50
	9.00  10.00  11.00  12.00
	13.50  14.50  15.50  16.50
*/
func NewMatrix(num_rows int, num_cols int, rows [][]float64) *Matrix {
	matrix := &Matrix{}
	matrix.rows = num_rows
	matrix.columns = num_cols
	var matrixValues [][]float64

	matrix.value = make([][]float64, matrix.rows)
	for r := range matrix.value {
		matrix.value[r] = make([]float64, matrix.columns)
	}

	// initialize all matrix values to be zero, if no values are provided
	if len(rows) == 0 {
		matrixValues = make([][]float64, matrix.rows)
		for r := range matrixValues {
			matrixValues[r] = make([]float64, matrix.columns)
		}
		rows = matrixValues
	}

	for r := 0; r < matrix.rows; r++ {
		for c := 0; c < matrix.columns; c++ {
			matrix.value[r][c] = rows[r][c]
		}
	}

	return matrix

}

func (m1 Matrix) IsEqual(m2 Matrix) bool {
	if m1.rows != m2.rows || m1.columns != m2.columns {
		return false
	}

	for r := 0; r < m1.rows; r++ {
		for c := 0; c < m2.columns; c++ {
			if !isFloatEqual(m1.value[r][c], m2.value[r][c]) {
				return false
			}
		}
	}

	return true
}

// PrintMatrix prints the matrix in a formatted way
func (m *Matrix) PrintMatrix() {
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.columns; c++ {
			fmt.Printf("%6.2f ", m.value[r][c]) // Format each value to 2 decimal places
		}
		fmt.Println() // Move to the next row
	}
}

// A 1x1 NAN matrix to depict malformed matrix operations
func NaNMatrix() *Matrix {
	return &Matrix{1, 1, [][]float64{{math.NaN()}}}
}

func (m1 Matrix) Multiply(m2 Matrix) *Matrix {
	/*
	   Two matrixs can only be multiplied, if the num of columns of first
	   matrix is equal to the number of rows of the second matrix

	   If matrix A is of size m x n and matrix B is of size n x p, then the
	   matrix can only be multipled when n = p
	*/
	if m1.columns != m2.rows {
		return NaNMatrix()
	}

	/*
		If matrix A is of size m x n and matrix B is of size n x p, then the
		resulting matrix AB will be of size m x p
	*/
	resultMatrix := NewMatrix(m1.rows, m2.columns, [][]float64{})

	for r := 0; r < resultMatrix.rows; r++ {
		for c := 0; c < resultMatrix.columns; c++ {
			var sum float64
			sum = 0
			// Multiply the row of first matrix with the column of second matrix
			for i := 0; i < m2.rows; i++ {
				sum += m1.value[r][i] * m2.value[i][c]
			}
			resultMatrix.value[r][c] = sum
		}
	}

	return resultMatrix
}

// Convert a tuple to a single column matrix
// TODO: generalize this, for now we only care about 4x1 matrix
func convertTupleToColumnMatrix(arr [4]float64) *Matrix {
	m := NewMatrix(4, 1, [][]float64{})

	for r := 0; r < len(arr); r++ {
		m.value[r][0] = arr[r]
	}

	return m
}

func (m Matrix) MultiplyTuple(tuple [4]float64) *Matrix {
	tupleMatrix := convertTupleToColumnMatrix(tuple)
	return m.Multiply(*tupleMatrix)
}

// Identity Matrix is responsible to allow us to inverse matrices
// Ref: https://www.reddit.com/r/learnmath/comments/on2s8z/taking_college_precal_what_are_the_point_of/
// for reasons why we need them
func IdentityMatrix() *Matrix {
	// Note: We only worry about 4x4 matrix right now
	return NewMatrix(4, 4, [][]float64{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})
}

func (m Matrix) Transpose() *Matrix {
	transposedMatrix := NewMatrix(m.columns, m.rows, [][]float64{})
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.columns; c++ {
			transposedMatrix.value[c][r] = m.value[r][c]
		}
	}
	return transposedMatrix
}

// Calculate deteminant of 2x2 matrix
func Determinant2(m Matrix) float64 {
	d := m.value[0][0]*m.value[1][1] - m.value[0][1]*m.value[1][0]
	return d
}

// Return back a new matrix by removing the row "row" and colum "col" from the
// matrix
func (m Matrix) SubMatrix(removeRow int, removeCol int) *Matrix {
	subMatrix := NewMatrix(m.rows-1, m.columns-1, [][]float64{})
	var subMatrixRow, subMatrixCol int

	for r := 0; r < m.rows; r++ {
		if r == removeRow {
			continue
		}
		for c := 0; c < m.columns; c++ {
			if c == removeCol {
				continue
			}
			subMatrix.value[subMatrixRow][subMatrixCol] = m.value[r][c]
			subMatrixCol += 1
		}
		subMatrixCol = 0
		subMatrixRow += 1
	}

	return subMatrix
}

// Minor for a 3x3 Matrix
// The minor of an element at row i and column j is the determinant of the sub-
// matrix at (i,j)
func Minor3(m Matrix, row int, col int) float64 {
	m_sub := m.SubMatrix(row, col)
	return Determinant2(*m_sub)
}

// Minors that have (possibly) had their sign changed
func Cofactor3(m Matrix, row int, col int) float64 {
	minor := Minor3(m, row, col)

	// If row + column is an odd number, then you negate the minor
	if (row+col)%2 != 0 {
		return minor * -1.0
	}
	return minor

}

func Determinant3(m Matrix) float64 {
	var det float64
	for col := 0; col < m.columns; col++ {
		det += m.value[0][col] * Cofactor3(m, 0, col)
	}
	return det
}

func Cofactor4(m Matrix, row int, col int) float64 {
	subM := m.SubMatrix(row, col)
	det := Determinant3(*subM)
	if (row+col)%2 != 0 {
		return det * -1.0
	}
	return det
}

// Since we focus only on 4x4 matrix, adding this as the function for the Matrix Struct
func (m Matrix) Determinant4() float64 {
	var det float64
	for col := 0; col < m.columns; col++ {
		det += m.value[0][col] * Cofactor4(m, 0, col)
	}
	return det
}

func (m Matrix) IsInvertible() bool {
	return m.Determinant4() != 0
}
