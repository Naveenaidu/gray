package math

import (
	"fmt"
	"math"
)

type Matrix struct {
	Rows    int
	Columns int
	Value   [][]float64
}

/*
Create a new matrix. Each element of "row" array defines the Rows of a
matrix

To Creates a matrix like below:

	1.00   2.00   3.00   4.00
	5.50   6.50   7.50   8.50
	9.00  10.00  11.00  12.00
	13.50  14.50  15.50  16.50
*/
func NewMatrix(num_rows int, num_cols int, Rows [][]float64) *Matrix {
	matrix := &Matrix{}
	matrix.Rows = num_rows
	matrix.Columns = num_cols
	var matrixValues [][]float64

	matrix.Value = make([][]float64, matrix.Rows)
	for r := range matrix.Value {
		matrix.Value[r] = make([]float64, matrix.Columns)
	}

	// initialize all matrix values to be zero, if no values are provided
	if len(Rows) == 0 {
		matrixValues = make([][]float64, matrix.Rows)
		for r := range matrixValues {
			matrixValues[r] = make([]float64, matrix.Columns)
		}
		Rows = matrixValues
	}

	for r := 0; r < matrix.Rows; r++ {
		for c := 0; c < matrix.Columns; c++ {
			matrix.Value[r][c] = Rows[r][c]
		}
	}

	return matrix

}

func (m1 Matrix) IsEqual(m2 Matrix) bool {
	if m1.Rows != m2.Rows || m1.Columns != m2.Columns {
		return false
	}

	for r := 0; r < m1.Rows; r++ {
		for c := 0; c < m2.Columns; c++ {
			if !IsFloatEqual(m1.Value[r][c], m2.Value[r][c]) {
				return false
			}
		}
	}

	return true
}

// PrintMatrix prints the matrix in a formatted way
func (m *Matrix) PrintMatrix() {
	for r := 0; r < m.Rows; r++ {
		for c := 0; c < m.Columns; c++ {
			fmt.Printf("%6.2f ", m.Value[r][c]) // Format each value to 2 decimal places
		}
		fmt.Println() // Move to the next row
	}
}

// A 1x1 NAN matrix to depict malformed matrix operations
func NaNMatrix() *Matrix {
	return &Matrix{1, 1, [][]float64{{math.NaN()}}}
}

// Note: During transformations points are converted to 4x1 matrix
func (m1 Matrix) ToPoint() *Point {
	return NewPoint(m1.Value[0][0], m1.Value[1][0], m1.Value[2][0])
}

// Note: During transformations vector are converted to 4x1 matrix
func (m1 Matrix) ToVector() *Vector {
	return NewVector(m1.Value[0][0], m1.Value[1][0], m1.Value[2][0])
}

func (m1 Matrix) Multiply(m2 Matrix) *Matrix {
	/*
	   Two matrixs can only be multiplied, if the num of Columns of first
	   matrix is equal to the number of Rows of the second matrix

	   If matrix A is of size m x n and matrix B is of size n x p, then the
	   matrix can only be multipled when n = p
	*/
	if m1.Columns != m2.Rows {
		return NaNMatrix()
	}

	/*
		If matrix A is of size m x n and matrix B is of size n x p, then the
		resulting matrix AB will be of size m x p
	*/
	resultMatrix := NewMatrix(m1.Rows, m2.Columns, [][]float64{})

	for r := 0; r < resultMatrix.Rows; r++ {
		for c := 0; c < resultMatrix.Columns; c++ {
			var sum float64
			sum = 0
			// Multiply the row of first matrix with the column of second matrix
			for i := 0; i < m2.Rows; i++ {
				sum += m1.Value[r][i] * m2.Value[i][c]
			}
			resultMatrix.Value[r][c] = sum
		}
	}

	return resultMatrix
}

// Convert a tuple to a single column matrix
// TODO: This can be removed since we now have ToPoint() and ToVector()
func convertTupleToColumnMatrix(arr [4]float64) *Matrix {
	m := NewMatrix(4, 1, [][]float64{})

	for r := 0; r < len(arr); r++ {
		m.Value[r][0] = arr[r]
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

/*
This video explains the concept of transpose better: https://www.youtube.com/watch?v=g4ecBFmvAYU

Some notes from the above:
  - Transpose means transforming some measurement device in other spaces back to
    that in original
  - The job of transpose is to find out which covector could we use to directly
    measure the original vector, such that we get the same measurement at the
    end
  - Transpose "looks like" inverse, but treat gridlines as covectors (my
    thought: we use transpose when we want to retain the length of the covector.
    That;s why we use this for computing normals)
*/
func (m Matrix) Transpose() *Matrix {
	transposedMatrix := NewMatrix(m.Columns, m.Rows, [][]float64{})
	for r := 0; r < m.Rows; r++ {
		for c := 0; c < m.Columns; c++ {
			transposedMatrix.Value[c][r] = m.Value[r][c]
		}
	}
	return transposedMatrix
}

// Calculate deteminant of 2x2 matrix
func Determinant2(m Matrix) float64 {
	d := m.Value[0][0]*m.Value[1][1] - m.Value[0][1]*m.Value[1][0]
	return d
}

// Return back a new matrix by removing the row "row" and colum "col" from the
// matrix
func (m Matrix) SubMatrix(removeRow int, removeCol int) *Matrix {
	subMatrix := NewMatrix(m.Rows-1, m.Columns-1, [][]float64{})
	var subMatrixRow, subMatrixCol int

	for r := 0; r < m.Rows; r++ {
		if r == removeRow {
			continue
		}
		for c := 0; c < m.Columns; c++ {
			if c == removeCol {
				continue
			}
			subMatrix.Value[subMatrixRow][subMatrixCol] = m.Value[r][c]
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
	for col := 0; col < m.Columns; col++ {
		det += m.Value[0][col] * Cofactor3(m, 0, col)
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
	for col := 0; col < m.Columns; col++ {
		det += m.Value[0][col] * Cofactor4(m, 0, col)
	}
	return det
}

func (m Matrix) IsInvertible() bool {
	return m.Determinant4() != 0
}

func (m Matrix) Inverse() *Matrix {
	if !m.IsInvertible() {
		return NaNMatrix()
	}

	det := m.Determinant4()
	invertedMatrix := NewMatrix(m.Rows, m.Columns, [][]float64{})

	for r := 0; r < m.Rows; r++ {
		for c := 0; c < m.Columns; c++ {
			cofactor := Cofactor4(m, r, c)
			// note that "col, row" here, instead of "row, col",
			// accomplishes the transpose operation!
			invertedMatrix.Value[c][r] = cofactor / det
		}
	}

	return invertedMatrix

}
