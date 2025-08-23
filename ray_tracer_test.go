package main

import (
	"fmt"
	"math"
	"testing"

	color "github.com/Naveenaidu/gray/src/core/color"
	core "github.com/Naveenaidu/gray/src/core/math"
	"github.com/Naveenaidu/gray/src/lighting"
	"github.com/Naveenaidu/gray/src/material"
	"github.com/Naveenaidu/gray/src/rayt"
	"github.com/Naveenaidu/gray/src/rendering"
	"github.com/Naveenaidu/gray/src/shape"
)

func TestAddTuples(t *testing.T) {
	// Case 1: Adding vector to point gives a new point
	a1 := core.NewPoint(3, -2, 5)
	a2 := core.NewVector(-2, 3, 1)
	a3 := a1.AddVector(*a2)

	if !a3.IsEqual(*core.NewPoint(1, 1, 6)) {
		t.Errorf("got: %+v, want: (1, 1, 6, 1)", a3)
	}

	// Case 2: Adding two vectors gives a new vector
	c2v1 := core.NewVector(3, -2, 5)
	c2v2 := core.NewVector(-2, 3, 1)
	c2v3 := core.AddVectors([]core.Vector{*c2v1, *c2v2})
	if !c2v3.IsEqual(*core.NewVector(1, 1, 6)) {
		t.Errorf("got: %+v, want: (1, 1, 6, 0)", c2v3)
	}

}

func TestSubtractTuples(t *testing.T) {

	// Case 1: Subtract two points, gives a new vector
	c1p1 := core.NewPoint(3, 2, 1)
	c1p2 := core.NewPoint(5, 6, 7)
	c1v1 := c1p1.Subtract(*c1p2)

	if !c1v1.IsEqual(*core.NewVector(-2, -4, -6)) {
		t.Errorf("got: %+v, want: (-2, -4, -6, 0)", c1v1)
	}

	// Case 2: Subtract a vector from a point, gives a new point
	// This is equivalent to walking from a point in the direction of the vector
	c2p1 := core.NewPoint(3, 2, 1)
	c2v1 := core.NewVector(5, 6, 7)
	c2p2 := c2p1.SubtractVector(*c2v1)
	if !c2p2.IsEqual(*core.NewPoint(-2, -4, -6)) {
		t.Errorf("got: %+v, want: (-2, -4, -6, 1)", c2p2)
	}

	// Case 3: Subtract two vectors, gives a new vector
	c3v1 := core.NewVector(3, 2, 1)
	c3v2 := core.NewVector(5, 6, 7)
	c3v3 := core.SubtractVectors([]core.Vector{*c3v1, *c3v2})
	if !c3v3.IsEqual(*core.NewVector(-2, -4, -6)) {
		t.Errorf("got: %+v, want: (-2, -4, -6, 0)", c3v3)
	}

	// Case 4: Subtract a vector from itself, gives a zero vector
	c4v1 := core.NewVector(3, 2, 1)
	c4v2 := core.SubtractVectors([]core.Vector{*c4v1, *c4v1})
	if !c4v2.IsEqual(*core.NewVector(0, 0, 0)) {
		t.Errorf("got: %+v, want: (0, 0, 0, 0)", c4v2)
	}

}

func TestNegateVector(t *testing.T) {
	a1 := core.NewVector(1, -2, 3)

	if !a1.Negate().IsEqual(*core.NewVector(-1, 2, -3)) {
		t.Errorf("got: %+v, want: (-1, 2, -3, 0)", a1.Negate())
	}
}

func TestScalarMuliply(t *testing.T) {
	p1 := core.NewPoint(1, -2, 3)
	p2 := p1.ScalarMultiply(3)
	if !p2.IsEqual(*core.NewPoint(3, -6, 9)) {
		t.Errorf("got: %+v, want: (3, -6, 9, 1)", p2)
	}
	v1 := core.NewVector(1, -2, 3)
	v2 := v1.ScalarMultiply(3)
	if !v2.IsEqual(*core.NewVector(3, -6, 9)) {
		t.Errorf("got: %+v, want: (3, -6, 9, 0)", v2)
	}

}

func TestScalarDivide(t *testing.T) {
	p1 := core.NewPoint(1, -2, 3)
	p2 := p1.ScalarDivide(2)
	if !p2.IsEqual(*core.NewPoint(0.5, -1, 1.5)) {
		t.Errorf("got: %+v, want: (0.5, -1, 1.5, 1)", p2)
	}
	v1 := core.NewVector(1, -2, 3)
	v2 := v1.ScalarDivide(2)
	if !v2.IsEqual(*core.NewVector(0.5, -1, 1.5)) {
		t.Errorf("got: %+v, want: (0.5, -1, 1.5, 0)", v2)
	}

}

func TestVectorMagnitude(t *testing.T) {

	v := core.NewVector(1, 0, 0)
	if !core.IsFloatEqual(v.Magnitude(), 1.0) {
		t.Errorf("got: %f, want: 1.0", v.Magnitude())
	}

	v = core.NewVector(0, 1, 0)
	if !core.IsFloatEqual(v.Magnitude(), 1.0) {
		t.Errorf("got: %f, want: 1.0", v.Magnitude())
	}

	v = core.NewVector(0, 0, 1)
	if !core.IsFloatEqual(v.Magnitude(), 1.0) {
		t.Errorf("got: %f, want: 1.0", v.Magnitude())
	}

	v = core.NewVector(1, 2, 3)
	if !core.IsFloatEqual(v.Magnitude(), math.Sqrt(14)) {
		t.Errorf("got: %f, want: %f", v.Magnitude(), math.Sqrt(14))
	}

	v = core.NewVector(-1, -2, -3)
	if !core.IsFloatEqual(v.Magnitude(), math.Sqrt(14)) {
		t.Errorf("got: %f, want: %f", v.Magnitude(), math.Sqrt(14))
	}

}

func TestNormalizeVector(t *testing.T) {
	var v, v_normalized *core.Vector

	v = core.NewVector(4, 0, 0)
	v_normalized = v.Normalize()
	if !core.IsFloatEqual(v_normalized.X, 1.0) {
		t.Errorf("got: %f, want: 1.0", v_normalized.X)
	}
	if !core.IsFloatEqual(v_normalized.Y, 0.0) {
		t.Errorf("got: %f, want: 0.0", v_normalized.Y)
	}
	if !core.IsFloatEqual(v_normalized.Z, 0.0) {
		t.Errorf("got: %f, want: 0.0", v_normalized.Z)
	}

	v = core.NewVector(1, 2, 3)
	v_normalized = v.Normalize()
	if !core.IsFloatEqual(v_normalized.X, 0.26726) {
		t.Errorf("got: %f, want: 0.26726", v_normalized.X)
	}
	if !core.IsFloatEqual(v_normalized.Y, 0.53452) {
		t.Errorf("got: %f, want: 0.53452", v_normalized.Y)
	}
	if !core.IsFloatEqual(v_normalized.Z, 0.80178) {
		t.Errorf("got: %f, want: 0.80178", v_normalized.Z)
	}

	// Check the magnitude of the normalized vector
	v = core.NewVector(1, 2, 3)
	if !core.IsFloatEqual(v.Normalize().Magnitude(), 1.0) {
		t.Errorf("got: %f, want: 1.0", v.Normalize().Magnitude())
	}
}

func TestDotProduct(t *testing.T) {
	a := core.NewVector(1, 2, 3)
	b := core.NewVector(2, 3, 4)

	dotProduct := a.DotProduct(*b)

	if dotProduct != 20.0 {
		t.Errorf("got: %f, want: 20.0", dotProduct)
	}

}

func TestCrossProduct(t *testing.T) {
	a := core.NewVector(1, 2, 3)
	b := core.NewVector(2, 3, 4)

	abCrossProduct := a.CrossProduct(*b)

	if !abCrossProduct.IsEqual(*core.NewVector(-1, 2, -1)) {
		t.Errorf("got: %+v, want: (-1, 2, -1)", abCrossProduct)
	}

	baCrossProduct := b.CrossProduct(*a)

	if !baCrossProduct.IsEqual(*core.NewVector(1, -2, 1)) {
		t.Errorf("got: %+v, want: (-1, 2, -1)", baCrossProduct)
	}

}

/* ------------- Canvas --------------- */

func TestNewCanvas(t *testing.T) {
	canvas := rendering.NewCanvas(10, 20, *color.Black)

	for x := 0; x < canvas.Width; x++ {
		for y := 0; y < canvas.Height; y++ {
			if canvas.Color[x][y] != *color.Black {
				t.Errorf("got: %+v, want: %+v", canvas.Color[x][y], *color.Black)
			}
		}
	}
}

func TestCanvasWritePixel(t *testing.T) {
	canvas := rendering.NewCanvas(80, 80, *color.Black)
	canvas.WritePixel(2, 3, *color.Red)

	if canvas.PixelAt(2, 3) != *color.Red {
		t.Errorf("got: %+v, want: %+v", canvas.PixelAt(2, 3), *color.Red)
	}

}

/* ------------- Matrix --------------- */

func TestNewMatrix(t *testing.T) {
	/*
		Creates a matrix like below:

		1.00   2.00   3.00   4.00
		5.50   6.50   7.50   8.50
		9.00  10.00  11.00  12.00
		13.50  14.50  15.50  16.50

		Note that the Columns are sent as arrays
	*/
	matrix := core.NewMatrix(4, 4, [][]float64{{1, 2, 3, 4}, {5.5, 6.5, 7.5, 8.5}, {9, 10, 11, 12}, {13.5, 14.5, 15.5, 16.5}})

	if matrix.Value[0][0] != 1 {
		t.Errorf("got: %f, want: %f", matrix.Value[0][0], 1.0)
	}

	if matrix.Value[0][3] != 4 {
		t.Errorf("got: %f, want: %f", matrix.Value[0][3], 4.0)
	}

	if matrix.Value[1][0] != 5.5 {
		t.Errorf("got: %f, want: %f", matrix.Value[1][0], 5.5)
	}

	if matrix.Value[1][2] != 7.5 {
		t.Errorf("got: %f, want: %f", matrix.Value[1][2], 7.5)
	}

	if matrix.Value[2][2] != 11 {
		t.Errorf("got: %f, want: %f", matrix.Value[2][2], 11.0)
	}

	if matrix.Value[3][0] != 13.5 {
		t.Errorf("got: %f, want: %f", matrix.Value[3][0], 13.5)
	}

	if matrix.Value[3][2] != 15.5 {
		t.Errorf("got: %f, want: %f", matrix.Value[3][2], 15.5)
	}
}

func Test2x2Matrix(t *testing.T) {
	matrix := core.NewMatrix(2, 2, [][]float64{{-3, 5}, {1, -2}})

	value := matrix.Value[0][0]
	expectedValue := -3.0
	if value != expectedValue {
		t.Errorf("got: %f, want: %f", value, expectedValue)
	}

	value = matrix.Value[0][1]
	expectedValue = 5.0
	if value != expectedValue {
		t.Errorf("got: %f, want: %f", value, expectedValue)
	}

	value = matrix.Value[1][0]
	expectedValue = 1.0
	if value != expectedValue {
		t.Errorf("got: %f, want: %f", value, expectedValue)
	}

	value = matrix.Value[1][1]
	expectedValue = -2.0
	if value != expectedValue {
		t.Errorf("got: %f, want: %f", value, expectedValue)
	}

}

func Test3x3Matrix(t *testing.T) {
	matrix := core.NewMatrix(3, 3, [][]float64{{-3, 5, 0}, {1, -2, -7}, {0, 1, 1}})

	value := matrix.Value[0][0]
	expectedValue := -3.0
	if value != expectedValue {
		t.Errorf("got: %f, want: %f", value, expectedValue)
	}

	value = matrix.Value[1][1]
	expectedValue = -2.0
	if value != expectedValue {
		t.Errorf("got: %f, want: %f", value, expectedValue)
	}

	value = matrix.Value[2][2]
	expectedValue = 1.0
	if value != expectedValue {
		t.Errorf("got: %f, want: %f", value, expectedValue)
	}

}

func TestMatrixEquality(t *testing.T) {
	m1 := core.NewMatrix(4, 4, [][]float64{{1, 5, 9, 5}, {2, 6, 8, 4}, {3, 7, 7, 3}, {4, 8, 6, 2}})
	m2 := core.NewMatrix(4, 4, [][]float64{{1, 5, 9, 5}, {2, 6, 8, 4}, {3, 7, 7, 3}, {4, 8, 6, 2}})
	if !m1.IsEqual(*m2) {
		t.Errorf("expected matrices to be equal - but they are not")
		m1.PrintMatrix()
		fmt.Println("-----------")
		m2.PrintMatrix()

	}

	m3 := core.NewMatrix(4, 4, [][]float64{{1, 5, 9, 5}, {2, 6, 8, 4}, {3, 7, 7, 3}, {4, 8, 6, 2}})
	m4 := core.NewMatrix(4, 4, [][]float64{{2, 6, 8, 4}, {3, 7, 7, 3}, {1, 5, 9, 5}, {4, 8, 6, 2}})
	if m3.IsEqual(*m4) {
		t.Errorf("expected matrices to be not equal - but they are")
		m3.PrintMatrix()
		fmt.Println("-----------")
		m4.PrintMatrix()

	}

}

func TestMatrixMultiply_InvalidDimensions(t *testing.T) {
	m1 := core.NewMatrix(2, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	m2 := core.NewMatrix(2, 2, [][]float64{
		{7, 8},
		{9, 10},
	})

	result := m1.Multiply(*m2)

	if result.Rows != 1 || result.Columns != 1 || math.IsNaN(result.Value[0][0]) == false {
		t.Errorf("Expected NaNMatrix, but got %v", result.Value)
	}
}

func TestMatrixMultiply(t *testing.T) {
	m1 := core.NewMatrix(4, 4, [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})

	m2 := core.NewMatrix(4, 4, [][]float64{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	})

	result := m1.Multiply(*m2)

	expectedMatrix := core.NewMatrix(4, 4, [][]float64{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	})

	if !result.IsEqual(*expectedMatrix) {
		t.Errorf("Multiplication result does not mathch: ")
		fmt.Println("expected: ")
		expectedMatrix.PrintMatrix()
		fmt.Println("\ngot: ")
		result.PrintMatrix()

	}
}

func TestMatrixMultiply_Tuple(t *testing.T) {
	m1 := core.NewMatrix(4, 4, [][]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	})

	result := m1.MultiplyTuple([4]float64{1, 2, 3, 1})

	expectedMatrix := core.NewMatrix(4, 1, [][]float64{
		{18},
		{24},
		{33},
		{1},
	})

	if !result.IsEqual(*expectedMatrix) {
		t.Errorf("Multiplication result does not mathch: ")
		fmt.Println("expected: ")
		expectedMatrix.PrintMatrix()
		fmt.Println("\ngot: ")
		result.PrintMatrix()

	}
}

func TestMatrixMultiply_IdentityMatrix(t *testing.T) {
	m1 := core.NewMatrix(4, 4, [][]float64{
		{1, 2, 3, 5},
		{4, 5, 6, 0},
		{7, 8, 9, 16},
		{-2, 1, 4.4, 11},
	})

	expected := m1

	result := m1.Multiply(*core.IdentityMatrix())

	if !result.IsEqual(*expected) {
		t.Errorf("Expected %v, but got %v", expected.Value, result.Value)
	}
}

func TestMatrixMultiply_Transpose(t *testing.T) {
	m1 := core.NewMatrix(4, 4, [][]float64{
		{0, 9, 3, 0},
		{9, 8, 0, 8},
		{1, 8, 5, 3},
		{0, 0, 5, 8},
	})

	expected := core.NewMatrix(4, 4, [][]float64{
		{0, 9, 1, 0},
		{9, 8, 8, 0},
		{3, 0, 5, 5},
		{0, 8, 3, 8},
	})

	result := m1.Transpose()

	if !result.IsEqual(*expected) {
		t.Errorf("Expected %v, but got %v", expected.Value, result.Value)
	}
}

func TestMatrixMultiply_TransposeIdentity(t *testing.T) {
	m1 := core.IdentityMatrix()

	result := m1.Transpose()

	if !result.IsEqual(*core.IdentityMatrix()) {
		t.Errorf("Expected %v, but got %v", core.IdentityMatrix().Value, result.Value)
	}
}

func TestMatrix_2x2Determinant(t *testing.T) {
	m := core.NewMatrix(2, 2, [][]float64{
		{1, 5},
		{-3, 2},
	})
	result := core.Determinant2(*m)
	expected := 17.0
	if !core.IsFloatEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestMatrix_3x3SubMatrix(t *testing.T) {
	m1 := core.NewMatrix(3, 3, [][]float64{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3},
	})

	expected := core.NewMatrix(2, 2, [][]float64{
		{-3, 2},
		{0, 6},
	})

	result := m1.SubMatrix(0, 2)

	if !result.IsEqual(*expected) {
		t.Errorf("Expected %v, but got %v", expected.Value, result.Value)
	}
}

func TestMatrix_4x4SubMatrix(t *testing.T) {
	m1 := core.NewMatrix(4, 4, [][]float64{
		{1, 2, 3, 4},
		{2, 4, -4, 2},
		{8, 6, 4, 1},
		{0, 10, 11, 1},
	})

	expected := core.NewMatrix(3, 3, [][]float64{
		{1, 3, 4},
		{2, -4, 2},
		{0, 11, 1},
	})

	result := m1.SubMatrix(2, 1)

	if !result.IsEqual(*expected) {
		t.Errorf("Expected %v, but got %v", expected.Value, result.Value)
	}
}

func TestMatrix_3x3Minor(t *testing.T) {
	m := core.NewMatrix(3, 3, [][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})
	result := core.Minor3(*m, 1, 0)
	expected := 25.0
	if !core.IsFloatEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestMatrix_3x3Cofactor(t *testing.T) {
	var result, expected float64
	m := core.NewMatrix(3, 3, [][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})
	result = core.Cofactor3(*m, 0, 0)
	expected = -12.0
	if !core.IsFloatEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	result = core.Cofactor3(*m, 1, 0)
	expected = -25.0
	if !core.IsFloatEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestMatrix_3x3Determinant(t *testing.T) {
	// Given the following 3x3 matrix A
	matrix := core.NewMatrix(3, 3, [][]float64{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	})

	// Test cofactor(A, 0, 0)
	result := core.Cofactor3(*matrix, 0, 0)
	expected := 56.0
	if !core.IsFloatEqual(result, expected) {
		t.Errorf("Expected cofactor(A, 0, 0) = %v, but got %v", expected, result)
	}

	// Test cofactor(A, 0, 1)
	result = core.Cofactor3(*matrix, 0, 1)
	expected = 12.0
	if !core.IsFloatEqual(result, expected) {
		t.Errorf("Expected cofactor(A, 0, 1) = %v, but got %v", expected, result)
	}

	// Test cofactor(A, 0, 2)
	result = core.Cofactor3(*matrix, 0, 2)
	expected = -46.0
	if !core.IsFloatEqual(result, expected) {
		t.Errorf("Expected cofactor(A, 0, 2) = %v, but got %v", expected, result)
	}

	// Test determinant(A)
	result = core.Determinant3(*matrix)
	expected = -196.0
	if !core.IsFloatEqual(result, expected) {
		t.Errorf("Expected determinant(A) = %v, but got %v", expected, result)
	}
}

func TestMatrix_4x4Determinant(t *testing.T) {
	// Given the following 4x4 matrix A
	matrix := core.NewMatrix(4, 4, [][]float64{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	})

	// Test cofactor(A, 0, 0)
	// subM1 := matrix.SubMatrix(0, 0)
	result := core.Cofactor4(*matrix, 0, 0)
	expected := 690.0
	if !core.IsFloatEqual(result, expected) {
		t.Errorf("Expected cofactor(A, 0, 0) = %v, but got %v", expected, result)
	}

	// Test cofactor(A, 0, 1)
	result = core.Cofactor4(*matrix, 0, 1)
	expected = 447.0
	if !core.IsFloatEqual(result, expected) {
		t.Errorf("Expected cofactor(A, 0, 1) = %v, but got %v", expected, result)
	}

	// Test cofactor(A, 0, 2)
	result = core.Cofactor4(*matrix, 0, 2)
	expected = 210.0
	if !core.IsFloatEqual(result, expected) {
		t.Errorf("Expected cofactor(A, 0, 2) = %v, but got %v", expected, result)
	}

	// Test cofactor(A, 0, 3)
	result = core.Cofactor4(*matrix, 0, 3)
	expected = 51.0
	if !core.IsFloatEqual(result, expected) {
		t.Errorf("Expected cofactor(A, 0, 3) = %v, but got %v", expected, result)
	}

	// Test determinant(A)
	result = matrix.Determinant4()
	expected = -4071.0
	if !core.IsFloatEqual(result, expected) {
		t.Errorf("Expected determinant(A) = %v, but got %v", expected, result)
	}
}

func TestMatrix_InvertibleMatrix(t *testing.T) {
	// Scenario: Testing an invertible matrix for invertibility
	// Det is -2120.0
	matrix := core.NewMatrix(4, 4, [][]float64{
		{6, 4, 4, 4},
		{5, 5, 7, 6},
		{4, -9, 3, -7},
		{9, 1, 7, -6},
	})

	// Test invertibility
	if !matrix.IsInvertible() {
		t.Errorf("Expected matrix to be invertible, but it is not")
	}

	// Scenario: Testing a noninvertible matrix for invertibility
	// Det is 0
	matrix2 := core.NewMatrix(4, 4, [][]float64{
		{-4, 2, -2, -3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{0, 0, 0, 0},
	})

	// Test invertibility
	if matrix2.IsInvertible() {
		t.Errorf("Expected matrix to be noninvertible, but it is invertible")
	}

}

func TestMatrix_Inverse(t *testing.T) {
	// Given the following 4x4 matrix A
	matrix := core.NewMatrix(4, 4, [][]float64{
		{8, -5, 9, 2},
		{7, 5, 6, 1},
		{-6, 0, 9, 6},
		{-3, 0, -9, -4},
	})

	// Expected inverse(A)
	expected := core.NewMatrix(4, 4, [][]float64{
		{-0.15385, -0.15385, -0.28205, -0.53846},
		{-0.07692, 0.12308, 0.02564, 0.03077},
		{0.35897, 0.35897, 0.43590, 0.92308},
		{-0.69231, -0.69231, -0.76923, -1.92308},
	})

	// Calculate inverse(A)
	result := matrix.Inverse()

	// Compare the result with the expected matrix
	if !result.IsEqual(*expected) {
		t.Errorf("Expected inverse(A) = %v, but got %v", expected.Value, result.Value)
	}

	// Given the following 4x4 matrix A
	matrix1 := core.NewMatrix(4, 4, [][]float64{
		{9, 3, 0, 9},
		{-5, -2, -6, -3},
		{-4, 9, 6, 4},
		{-7, 6, 6, 2},
	})

	// Expected inverse(A)
	expected1 := core.NewMatrix(4, 4, [][]float64{
		{-0.04074, -0.07778, 0.14444, -0.22222},
		{-0.07778, 0.03333, 0.36667, -0.33333},
		{-0.02901, -0.14630, -0.10926, 0.12963},
		{0.17778, 0.06667, -0.26667, 0.33333},
	})

	// Calculate inverse(A)
	result1 := matrix1.Inverse()

	// Compare the result with the expected matrix
	if !result1.IsEqual(*expected1) {
		t.Errorf("Expected inverse(A) = %v, but got %v", expected1.Value, result1.Value)
	}
}

func TestMatrix_ProductWithInverse(t *testing.T) {
	// Test that if "C ← A * B" then  "C * (1/B) = A"
	// Given the following 4x4 matrix A
	matrixA := core.NewMatrix(4, 4, [][]float64{
		{3, -9, 7, 3},
		{3, -8, 2, -9},
		{-4, 4, 4, 1},
		{-6, 5, -1, 1},
	})

	// And the following 4x4 matrix B
	matrixB := core.NewMatrix(4, 4, [][]float64{
		{8, 2, 2, 2},
		{3, -1, 7, 0},
		{7, 0, 5, 4},
		{6, -2, 0, 5},
	})

	// Calculate C ← A * B
	matrixC := matrixA.Multiply(*matrixB)

	// Calculate C * inverse(B)
	matrixBInverse := matrixB.Inverse()
	result := matrixC.Multiply(*matrixBInverse)

	// Compare the result with matrix A
	if !result.IsEqual(*matrixA) {
		t.Errorf("Expected C * inverse(B) = A, but got %v", result.Value)
	}
}

/* ------------- Transformations --------------- */
func TestTranslatePoint(t *testing.T) {
	p := core.NewPoint(-3, 4, 5)
	result := p.Translate(5, -3, 2)
	if !result.IsEqual(*core.NewPoint(2, 1, 7)) {
		t.Errorf("got: %+v, want: (2, 1, 7)", result)
	}

	p1 := core.NewPoint(-3, 4, 5)
	tM := core.TranslationM(5, -3, 2).Inverse()
	result1 := tM.Multiply(*p1.ToMatrix()).ToPoint()
	if !result1.IsEqual(*core.NewPoint(-8, 7, 3)) {
		t.Errorf("got: %+v, want: (-8, 7, 3)", result)
	}

}

func TestScalingMatrixAppliedToPoint(t *testing.T) {
	// Given p ← point(-4, 6, 8)
	p := core.NewPoint(-4, 6, 8)

	// When scaling is applied with factors (2, 3, 4)
	result := p.Scale(2, 3, 4)

	// Then the result should be point(-8, 18, 32)
	expected := core.NewPoint(-8, 18, 32)

	if !result.IsEqual(*expected) {
		t.Errorf("Expected scaling result = %v, but got %v", expected, result)
	}
}

func TestScalingMatrixAppliedToVector(t *testing.T) {
	// Given v ← vector(-4, 6, 8)
	v := core.NewVector(-4, 6, 8)

	// When scaling is applied with factors (2, 3, 4)
	result := v.Scale(2, 3, 4)

	// Then the result should be vector(-8, 18, 32)
	expected := core.NewVector(-8, 18, 32)

	if !result.IsEqual(*expected) {
		t.Errorf("Expected scaling result = %v, but got %v", expected, result)
	}
}

func TestMultiplyingByInverseOfScalingMatrix(t *testing.T) {
	// Given transform ← scaling(2, 3, 4)
	transform := core.ScaleM(2, 3, 4)

	// And inv ← inverse(transform)
	inv := transform.Inverse()

	// And v ← vector(-4, 6, 8)
	v := core.NewVector(-4, 6, 8)

	// Then inv * v = vector(-2, 2, 2)
	result := inv.Multiply(*v.ToMatrix()).ToVector()
	expected := core.NewVector(-2, 2, 2)

	if !result.IsEqual(*expected) {
		t.Errorf("Expected inv * v = %v, but got %v", expected, result)
	}
}

func TestReflectionIsScalingByNegativeValue(t *testing.T) {
	// Given p ← point(2, 3, 4)
	p := core.NewPoint(2, 3, 4)

	// When scaling is applied with factors (-1, 1, 1)
	result := p.Scale(-1, 1, 1)

	// Then the result should be point (-2, 3, 4), point is reflected on x axis
	expected := core.NewPoint(-2, 3, 4)

	if !result.IsEqual(*expected) {
		t.Errorf("Expected scaling result = %v, but got %v", expected, result)
	}
}

func TestRotatingPointAroundXAxis(t *testing.T) {
	// Given p ← point(0, 1, 0)
	p := core.NewPoint(0, 1, 0)

	// And half_quarter ← rotation_x(π / 4)
	halfQuarter := core.RotateXM(math.Pi / 4)

	// And full_quarter ← rotation_x(π / 2)
	fullQuarter := core.RotateXM(math.Pi / 2)

	// Then half_quarter * p = point(0, √2/2, √2/2)
	resultHalfQuarter := halfQuarter.Multiply(*p.ToMatrix()).ToPoint()
	expectedHalfQuarter := core.NewPoint(0, math.Sqrt(2)/2, math.Sqrt(2)/2)

	if !resultHalfQuarter.IsEqual(*expectedHalfQuarter) {
		t.Errorf("Expected half_quarter * p = %v, but got %v", expectedHalfQuarter, resultHalfQuarter)
	}

	// And full_quarter * p = point(0, 0, 1)
	resultFullQuarter := fullQuarter.Multiply(*p.ToMatrix()).ToPoint()
	expectedFullQuarter := core.NewPoint(0, 0, 1)

	if !resultFullQuarter.IsEqual(*expectedFullQuarter) {
		t.Errorf("Expected full_quarter * p = %v, but got %v", expectedFullQuarter, resultFullQuarter)
	}
}

func TestInverseOfXRotationRotatesOppositeDirection(t *testing.T) {
	// Given p ← point(0, 1, 0)
	p := core.NewPoint(0, 1, 0)

	// And half_quarter ← rotation_x(π / 4)
	halfQuarter := core.RotateXM(math.Pi / 4)

	// And inv ← inverse(half_quarter)
	inv := halfQuarter.Inverse()

	// Then inv * p = point(0, √2/2, -√2/2)
	result := inv.Multiply(*p.ToMatrix()).ToPoint()
	expected := core.NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)

	if !result.IsEqual(*expected) {
		t.Errorf("Expected inv * p = %v, but got %v", expected, result)
	}
}

func TestRotatingPointAroundYAxis(t *testing.T) {
	// Given p ← point(0, 0, 1)
	p := core.NewPoint(0, 0, 1)

	// And half_quarter ← rotation_y(π / 4)
	halfQuarter := core.RotateYM(math.Pi / 4)

	// And full_quarter ← rotation_y(π / 2)
	fullQuarter := core.RotateYM(math.Pi / 2)

	// Then half_quarter * p = point(√2/2, 0, √2/2)
	resultHalfQuarter := halfQuarter.Multiply(*p.ToMatrix()).ToPoint()
	expectedHalfQuarter := core.NewPoint(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)

	if !resultHalfQuarter.IsEqual(*expectedHalfQuarter) {
		t.Errorf("Expected half_quarter * p = %v, but got %v", expectedHalfQuarter, resultHalfQuarter)
	}

	// And full_quarter * p = point(1, 0, 0)
	resultFullQuarter := fullQuarter.Multiply(*p.ToMatrix()).ToPoint()
	expectedFullQuarter := core.NewPoint(1, 0, 0)

	if !resultFullQuarter.IsEqual(*expectedFullQuarter) {
		t.Errorf("Expected full_quarter * p = %v, but got %v", expectedFullQuarter, resultFullQuarter)
	}
}

func TestRotatingPointAroundZAxis(t *testing.T) {
	// Given p ← point(0, 1, 0)
	p := core.NewPoint(0, 1, 0)

	// And half_quarter ← rotation_z(π / 4)
	halfQuarter := core.RotateZM(math.Pi / 4)

	// And full_quarter ← rotation_z(π / 2)
	fullQuarter := core.RotateZM(math.Pi / 2)

	// Then half_quarter * p = point(-√2/2, √2/2, 0)
	resultHalfQuarter := halfQuarter.Multiply(*p.ToMatrix()).ToPoint()
	expectedHalfQuarter := core.NewPoint(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)

	if !resultHalfQuarter.IsEqual(*expectedHalfQuarter) {
		t.Errorf("Expected half_quarter * p = %v, but got %v", expectedHalfQuarter, resultHalfQuarter)
	}

	// And full_quarter * p = point(-1, 0, 0)
	resultFullQuarter := fullQuarter.Multiply(*p.ToMatrix()).ToPoint()
	expectedFullQuarter := core.NewPoint(-1, 0, 0)

	if !resultFullQuarter.IsEqual(*expectedFullQuarter) {
		t.Errorf("Expected full_quarter * p = %v, but got %v", expectedFullQuarter, resultFullQuarter)
	}
}

func TestShearingTransformation(t *testing.T) {
	// Scenario: A shearing transformation moves x in proportion to z
	// Given transform ← shearing(0, 1, 0, 0, 0, 0)
	// And p ← point(2, 3, 4)
	p := core.NewPoint(2, 3, 4)
	result := p.Shear(0, 1, 0, 0, 0, 0)
	expected := core.NewPoint(6, 3, 4)

	if !result.IsEqual(*expected) {
		t.Errorf("Expected transform * p = %v, but got %v", expected, result)
	}

	// Scenario: A shearing transformation moves y in proportion to x
	// Given transform ← shearing(0, 0, 1, 0, 0, 0)
	// And p ← point(2, 3, 4)
	p = core.NewPoint(2, 3, 4)
	result = p.Shear(0, 0, 1, 0, 0, 0)
	expected = core.NewPoint(2, 5, 4)

	if !result.IsEqual(*expected) {
		t.Errorf("Expected transform * p = %v, but got %v", expected, result)
	}

	// Scenario: A shearing transformation moves y in proportion to z
	// Given transform ← shearing(0, 0, 0, 1, 0, 0)
	// And p ← point(2, 3, 4)
	p = core.NewPoint(2, 3, 4)
	result = p.Shear(0, 0, 0, 1, 0, 0)
	expected = core.NewPoint(2, 7, 4)

	if !result.IsEqual(*expected) {
		t.Errorf("Expected transform * p = %v, but got %v", expected, result)
	}

	// Scenario: A shearing transformation moves z in proportion to x
	// Given transform ← shearing(0, 0, 0, 0, 1, 0)
	// And p ← point(2, 3, 4)
	p = core.NewPoint(2, 3, 4)
	result = p.Shear(0, 0, 0, 0, 1, 0)
	expected = core.NewPoint(2, 3, 6)

	if !result.IsEqual(*expected) {
		t.Errorf("Expected transform * p = %v, but got %v", expected, result)
	}

	// Scenario: A shearing transformation moves z in proportion to y
	// Given transform ← shearing(0, 0, 0, 0, 0, 1)
	// And p ← point(2, 3, 4)
	p = core.NewPoint(2, 3, 4)
	result = p.Shear(0, 0, 0, 0, 0, 1)
	expected = core.NewPoint(2, 3, 7)

	if !result.IsEqual(*expected) {
		t.Errorf("Expected transform * p = %v, but got %v", expected, result)
	}
}

func TestChainingTransformations(t *testing.T) {
	//Scenario: Individual transformations are applied in sequence
	p := core.NewPoint(1, 0, 1)
	pM := p.ToMatrix()
	rotateXM := core.RotateXM(math.Pi / 2)
	scaleM := core.ScaleM(5, 5, 5)
	translationM := core.TranslationM(10, 5, 7)

	// apply rotation first
	pRotateM := rotateXM.Multiply(*pM)
	pRotate := pRotateM.ToPoint()
	expected := core.NewPoint(1, -1, 0)
	if !pRotate.IsEqual(*expected) {
		t.Errorf("Expected transform * p = %v, but got %v", expected, pRotate)
	}

	// then apply scaling
	pRotateAndScaleM := scaleM.Multiply(*pRotateM)
	pRotateAndScale := pRotateAndScaleM.ToPoint()
	expected = core.NewPoint(5, -5, 0)
	if !pRotateAndScale.IsEqual(*expected) {
		t.Errorf("Expected transform * p = %v, but got %v", expected, pRotateAndScale)
	}

	// then apply translation
	pRotateAndScaleAndTranslateM := translationM.Multiply(*pRotateAndScaleM)
	pRotateAndScaleAndTranslate := pRotateAndScaleAndTranslateM.ToPoint()
	expected = core.NewPoint(15, 0, 7)
	if !pRotateAndScaleAndTranslate.IsEqual(*expected) {
		t.Errorf("Expected transform * p = %v, but got %v", expected, pRotateAndScaleAndTranslate)
	}

	// Scenario: Chained transformations must be applied in reverse order
	pRotateAndScaleAndTranslateChainedM := translationM.Multiply(*scaleM).Multiply(*rotateXM).Multiply(*pM)
	pRotateAndScaleAndTranslateChained := pRotateAndScaleAndTranslateChainedM.ToPoint()
	expected = core.NewPoint(15, 0, 7)
	if !pRotateAndScaleAndTranslateChained.IsEqual(*expected) {
		t.Errorf("Expected transform * p = %v, but got %v", expected, pRotateAndScaleAndTranslateChained)
	}

	// Scenario: Chained Transformations using ChainedTransforms()
	chainedTransformM := core.ChainTransforms([]*core.Matrix{
		pM,
		core.RotateXM(math.Pi / 2),
		core.ScaleM(5, 5, 5),
		core.TranslationM(10, 5, 7),
	})
	result := chainedTransformM.ToPoint()
	expected = core.NewPoint(15, 0, 7)
	if !result.IsEqual(*expected) {
		t.Errorf("Expected transform * p = %v, but got %v", expected, result)
	}

}

/* ------------- Rays --------------- */
func TestRayPosition(t *testing.T) {
	// Given r ← ray(point(2, 3, 4), vector(1, 0, 0))
	origin := core.NewPoint(2, 3, 4)
	direction := core.NewVector(1, 0, 0)
	r := rayt.Ray{Origin: *origin, Direction: *direction}

	// Then position(r, 0) = point(2, 3, 4)
	result := r.Position(0)
	expected := core.NewPoint(2, 3, 4)
	if !result.IsEqual(*expected) {
		t.Errorf("Expected position(r, 0) = %v, but got %v", expected, result)
	}

	// And position(r, 1) = point(3, 3, 4)
	result = r.Position(1)
	expected = core.NewPoint(3, 3, 4)
	if !result.IsEqual(*expected) {
		t.Errorf("Expected position(r, 1) = %v, but got %v", expected, result)
	}

	// And position(r, -1) = point(1, 3, 4)
	result = r.Position(-1)
	expected = core.NewPoint(1, 3, 4)
	if !result.IsEqual(*expected) {
		t.Errorf("Expected position(r, -1) = %v, but got %v", expected, result)
	}

	// And position(r, 2.5) = point(4.5, 3, 4)
	result = r.Position(2.5)
	expected = core.NewPoint(4.5, 3, 4)
	if !result.IsEqual(*expected) {
		t.Errorf("Expected position(r, 2.5) = %v, but got %v", expected, result)
	}
}

func TestRayIntersectsSphereAtTwoPoints(t *testing.T) {
	// Scenario: A ray intersects a sphere at two points
	// Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
	r := rayt.Ray{
		Origin:    *core.NewPoint(0, 0, -5),
		Direction: *core.NewVector(0, 0, 1),
	}

	// And s ← sphere()
	s := shape.UnitSphere()

	// When xs ← intersect(s, r)
	xs := r.IntersectSphere(*s)

	// Then xs.count = 2
	if len(xs) != 2 {
		t.Errorf("Expected xs.count = 2, but got %d", len(xs))
	}

	// And xs[0] = 4.0
	if xs[0].T != 4.0 {
		t.Errorf("Expected xs[0] = 4.0, but got %v", xs[0])
	}

	// And xs[1] = 6.0
	if xs[1].T != 6.0 {
		t.Errorf("Expected xs[1] = 6.0, but got %v", xs[1])
	}
}

func TestRayIntersectsSphereAtTangent(t *testing.T) {
	// Scenario: A ray intersects a sphere at a tangent
	// Given r ← ray(point(0, 1, -5), vector(0, 0, 1))
	r := rayt.Ray{
		Origin:    *core.NewPoint(0, 1, -5),
		Direction: *core.NewVector(0, 0, 1),
	}

	// And s ← sphere()
	s := shape.UnitSphere()

	// When xs ← intersect(s, r)
	xs := r.IntersectSphere(*s)

	// Then xs.count = 0
	if len(xs) != 2 {
		t.Errorf("Expected xs.count = 2, but got %d", len(xs))
	}

	// And xs[0] = 5.0
	if xs[0].T != 5.0 {
		t.Errorf("Expected xs[0] = 5.0, but got %v", xs[0])
	}

	// And xs[1] = 5.0
	if xs[1].T != 5.0 {
		t.Errorf("Expected xs[1] = 5.0, but got %v", xs[1])
	}
}

func TestRayMissesSphere(t *testing.T) {
	// Scenario: A ray misses a sphere
	// Given r ← ray(point(0, 2, -5), vector(0, 0, 1))
	r := rayt.Ray{
		Origin:    *core.NewPoint(0, 2, -5),
		Direction: *core.NewVector(0, 0, 1),
	}

	// And s ← sphere()
	s := shape.UnitSphere()

	// When xs ← intersect(s, r)
	xs := r.IntersectSphere(*s)

	// Then xs.count = 0
	if len(xs) != 0 {
		t.Errorf("Expected xs.count = 0, but got %d", len(xs))
	}
}

func TestRayOriginatesInsideSphere(t *testing.T) {
	// Scenario: A ray originates inside a sphere
	// Given r ← ray(point(0, 0, 0), vector(0, 0, 1))
	r := rayt.Ray{
		Origin:    *core.NewPoint(0, 0, 0),
		Direction: *core.NewVector(0, 0, 1),
	}

	// And s ← sphere()
	s := shape.UnitSphere()

	// When xs ← intersect(s, r)
	xs := r.IntersectSphere(*s)

	// Then xs.count = 2
	if len(xs) != 2 {
		t.Errorf("Expected xs.count = 2, but got %d", len(xs))
	}

	// And xs[0] = -1.0
	if xs[0].T != -1.0 {
		t.Errorf("Expected xs[0] = -1.0, but got %v", xs[0])
	}

	// And xs[1] = 1.0
	if xs[1].T != 1.0 {
		t.Errorf("Expected xs[1] = 1.0, but got %v", xs[1])
	}
}

func TestSphereIsBehindRay(t *testing.T) {
	// Scenario: A sphere is behind a ray
	// Given r ← ray(point(0, 0, 5), vector(0, 0, 1))
	r := rayt.Ray{
		Origin:    *core.NewPoint(0, 0, 5),
		Direction: *core.NewVector(0, 0, 1),
	}

	// And s ← sphere()
	s := shape.UnitSphere()

	// When xs ← intersect(s, r)
	xs := r.IntersectSphere(*s)

	// Then xs.count = 2
	if len(xs) != 2 {
		t.Errorf("Expected xs.count = 2, but got %d", len(xs))
	}

	// And xs[0] = -6.0
	if xs[0].T != -6.0 {
		t.Errorf("Expected xs[0] = -6.0, but got %v", xs[0])
	}

	// And xs[1] = -4.0
	if xs[1].T != -4.0 {
		t.Errorf("Expected xs[1] = -4.0, but got %v", xs[1])
	}
}

func TestHit_AllPositiveT(t *testing.T) {
	s := shape.UnitSphere()
	i1 := rayt.NewIntersection(1, *s)
	i2 := rayt.NewIntersection(2, *s)
	xs := []rayt.Intersection{i2, i1}

	hit := rayt.Ray{}.Hit(xs)
	if hit == nil {
		t.Errorf("Expect hit to be i1 (%v), but got nil value", i1)
	} else if hit.T != i1.T {
		t.Errorf("Expected hit to be i1 (%v), but got %v", i1, hit)
	}
}

func TestHit_SomeNegativeT(t *testing.T) {
	s := shape.UnitSphere()
	i1 := rayt.NewIntersection(-1, *s)
	i2 := rayt.NewIntersection(1, *s)
	xs := []rayt.Intersection{i2, i1}

	hit := rayt.Ray{}.Hit(xs)
	if hit == nil {
		t.Errorf("Expect hit to be i2 (%v), but got nil value", i2)
	} else if hit.T != i2.T {
		t.Errorf("Expected hit to be i2 (%v), but got %v", i2, hit)
	}
}

func TestHit_AllNegativeT(t *testing.T) {
	s := shape.UnitSphere()
	i1 := rayt.NewIntersection(-2, *s)
	i2 := rayt.NewIntersection(-1, *s)
	xs := []rayt.Intersection{i2, i1}

	hit := rayt.Ray{}.Hit(xs)

	if hit != nil {
		t.Errorf("Expected hit to be nil, but got %v", hit)
	}
}

func TestHit_LowestNonnegativeIntersection(t *testing.T) {
	s := shape.UnitSphere()
	i1 := rayt.NewIntersection(5, *s)
	i2 := rayt.NewIntersection(7, *s)
	i3 := rayt.NewIntersection(-3, *s)
	i4 := rayt.NewIntersection(2, *s)
	xs := []rayt.Intersection{i1, i2, i3, i4}

	hit := rayt.Ray{}.Hit(xs)
	if hit == nil {
		t.Errorf("Expect hit to be i4 (%v), but got nil value", i4)
	} else if hit.T != i4.T {
		t.Errorf("Expected hit to be i4 (%v), but got %v", i4, hit)
	}
}

func TestTranslatingRay(t *testing.T) {
	// Scenario: Translating a ray
	// Given r ← ray(point(1, 2, 3), vector(0, 1, 0))
	r := rayt.Ray{
		Origin:    *core.NewPoint(1, 2, 3),
		Direction: *core.NewVector(0, 1, 0),
	}
	// And m ← translation(3, 4, 5)
	m := core.TranslationM(3, 4, 5)
	// When r2 ← transform(r, m)
	r2 := r.Transform(m)
	// Then r2.origin = point(4, 6, 8)
	expectedOrigin := core.NewPoint(4, 6, 8)
	if !r2.Origin.IsEqual(*expectedOrigin) {
		t.Errorf("Expected r2.origin = %v, but got %v", expectedOrigin, r2.Origin)
	}
	// And r2.direction = vector(0, 1, 0)
	expectedDirection := core.NewVector(0, 1, 0)
	if !r2.Direction.IsEqual(*expectedDirection) {
		t.Errorf("Expected r2.direction = %v, but got %v", expectedDirection, r2.Direction)
	}
}

func TestScalingRay(t *testing.T) {
	// Scenario: Scaling a ray
	// Given r ← ray(point(1, 2, 3), vector(0, 1, 0))
	r := rayt.Ray{
		Origin:    *core.NewPoint(1, 2, 3),
		Direction: *core.NewVector(0, 1, 0),
	}
	// And m ← scaling(2, 3, 4)
	m := core.ScaleM(2, 3, 4)
	// When r2 ← transform(r, m)
	r2 := r.Transform(m)
	// Then r2.origin = point(2, 6, 12)
	expectedOrigin := core.NewPoint(2, 6, 12)
	if !r2.Origin.IsEqual(*expectedOrigin) {
		t.Errorf("Expected r2.origin = %v, but got %v", expectedOrigin, r2.Origin)
	}
	// And r2.direction = vector(0, 3, 0)
	expectedDirection := core.NewVector(0, 3, 0)
	if !r2.Direction.IsEqual(*expectedDirection) {
		t.Errorf("Expected r2.direction = %v, but got %v", expectedDirection, r2.Direction)
	}
}

func TestSphereDefaultTransformation(t *testing.T) {
	// Scenario: A sphere's default transformation
	// Given s ← sphere()
	s := shape.UnitSphere()
	// Then s.transform = identity_matrix
	if !s.Transform.IsEqual(*core.IdentityMatrix()) {
		t.Errorf("Expected sphere's default transform to be identity matrix, but got %v", s.Transform.Value)
	}
}

func TestChangingSphereTransformation(t *testing.T) {
	// Scenario: Changing a sphere's transformation
	// Given s ← sphere()
	s := shape.UnitSphere()
	// And t ← translation(2, 3, 4)
	tm := core.TranslationM(2, 3, 4)
	// When set_transform(s, t)
	s.Transform = *tm
	// Then s.transform = t
	if !s.Transform.IsEqual(*tm) {
		t.Errorf("Expected sphere's transform to be %v, but got %v", tm.Value, s.Transform.Value)
	}
}

func TestIntersectingScaledSphereWithRay(t *testing.T) {
	// Scenario: Intersecting a scaled sphere with a ray
	// Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
	r := rayt.Ray{
		Origin:    *core.NewPoint(0, 0, -5),
		Direction: *core.NewVector(0, 0, 1),
	}
	// And s ← sphere()
	s := shape.UnitSphere()
	// When set_transform(s, scaling(2, 2, 2))
	s.Transform = *core.ScaleM(2, 2, 2)

	xs := r.IntersectSphere(*s)
	// Then xs.count = 2
	if len(xs) != 2 {
		t.Errorf("Expected xs.count = 2, but got %d", len(xs))
	}
	// And xs[0].t = 3
	if !core.IsFloatEqual(xs[0].T, 3.0) {
		t.Errorf("Expected xs[0].t = 3, but got %v", xs[0].T)
	}
	// And xs[1].t = 7
	if !core.IsFloatEqual(xs[1].T, 7.0) {
		t.Errorf("Expected xs[1].t = 7, but got %v", xs[1].T)
	}
}

func TestIntersectingTranslatedSphereWithRay(t *testing.T) {
	// Scenario: Intersecting a translated sphere with a ray
	// Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
	r := rayt.Ray{
		Origin:    *core.NewPoint(0, 0, -5),
		Direction: *core.NewVector(0, 0, 1),
	}
	// And s ← sphere()
	s := shape.UnitSphere()
	// When set_transform(s, translation(5, 0, 0))
	s.Transform = *core.TranslationM(5, 0, 0)

	xs := r.IntersectSphere(*s)
	// Then xs.count = 0
	if len(xs) != 0 {
		t.Errorf("Expected xs.count = 0, but got %d", len(xs))
	}
}

/* ------------- Lighting and Shading --------------- */

func TestSphereNormalAt(t *testing.T) {
	// Scenario: The normal on a sphere at a point on the x axis
	// Given s ← sphere()
	s := shape.UnitSphere()
	// When n ← normal_at(s, point(1, 0, 0))
	n := lighting.NormalAt(*s, *core.NewPoint(1, 0, 0))
	// Then n = vector(1, 0, 0)
	expected := core.NewVector(1, 0, 0)
	if !n.IsEqual(*expected) {
		t.Errorf("Expected normal = %v, but got %v", expected, n)
	}

	// Scenario: The normal on a sphere at a point on the y axis
	// Given s ← sphere()
	s = shape.UnitSphere()
	// When n ← normal_at(s, point(0, 1, 0))
	n = lighting.NormalAt(*s, *core.NewPoint(0, 1, 0))
	// Then n = vector(0, 1, 0)
	expected = core.NewVector(0, 1, 0)
	if !n.IsEqual(*expected) {
		t.Errorf("Expected normal = %v, but got %v", expected, n)
	}

	// Scenario: The normal on a sphere at a point on the z axis
	// Given s ← sphere()
	s = shape.UnitSphere()
	// When n ← normal_at(s, point(0, 0, 1))
	n = lighting.NormalAt(*s, *core.NewPoint(0, 0, 1))
	// Then n = vector(0, 0, 1)
	expected = core.NewVector(0, 0, 1)
	if !n.IsEqual(*expected) {
		t.Errorf("Expected normal = %v, but got %v", expected, n)
	}

	// Scenario: The normal on a sphere at a nonaxial point
	// Given s ← sphere()
	s = shape.UnitSphere()
	// When n ← normal_at(s, point(√3/3, √3/3, √3/3))
	sqrtThird := math.Sqrt(3) / 3
	n = lighting.NormalAt(*s, *core.NewPoint(sqrtThird, sqrtThird, sqrtThird))
	// Then n = vector(√3/3, √3/3, √3/3)
	expected = core.NewVector(sqrtThird, sqrtThird, sqrtThird)
	if !n.IsEqual(*expected) {
		t.Errorf("Expected normal = %v, but got %v", expected, n)
	}
}

func TestSphereNormalIsNormalized(t *testing.T) {
	// Scenario: The normal is a normalized vector
	// Given s ← sphere()
	s := shape.UnitSphere()
	// When n ← normal_at(s, point(√3/3, √3/3, √3/3))
	sqrtThird := math.Sqrt(3) / 3
	n := lighting.NormalAt(*s, *core.NewPoint(sqrtThird, sqrtThird, sqrtThird))
	// Then n = normalize(n)
	normalized := n.Normalize()
	if !n.IsEqual(*normalized) {
		t.Errorf("Expected normal to be normalized, but n = %v and normalize(n) = %v", n, normalized)
	}

	// Also verify that the magnitude is 1
	if !core.IsFloatEqual(n.Magnitude(), 1.0) {
		t.Errorf("Expected normal magnitude to be 1.0, but got %v", n.Magnitude())
	}
}

func TestComputingNormalOnTranslatedSphere(t *testing.T) {
	// Scenario: Computing the normal on a translated sphere
	// Given s ← sphere()
	s := shape.UnitSphere()
	// And set_transform(s, translation(0, 1, 0))
	s.Transform = *core.ChainTransforms([]*core.Matrix{
		core.TranslationM(0, 1, 0),
	})
	// When n ← normal_at(s, point(0, 1.70711, -0.70711))
	n := lighting.NormalAt(*s, *core.NewPoint(0, 1.70711, -0.70711))
	// Then n = vector(0, 0.70711, -0.70711)
	expected := core.NewVector(0, 0.70711, -0.70711)
	if !n.IsEqual(*expected) {
		t.Errorf("Expected normal = %v, but got %v", expected, n)
	}
}

func TestComputingNormalOnTransformedSphere(t *testing.T) {
	// Scenario: Computing the normal on a transformed sphere
	// Given s ← sphere()
	s := shape.UnitSphere()
	// And m ← scaling(1, 0.5, 1) * rotation_z(π/5)
	// And set_transform(s, m)
	s.Transform = *core.ChainTransforms([]*core.Matrix{
		core.RotateZM(math.Pi / 5),
		core.ScaleM(1, 0.5, 1),
	})
	// When n ← normal_at(s, point(0, √2/2, -√2/2))
	sqrtHalf := math.Sqrt(2) / 2
	n := lighting.NormalAt(*s, *core.NewPoint(0, sqrtHalf, -sqrtHalf))
	// Then n = vector(0, 0.97014, -0.24254)
	expected := core.NewVector(0, 0.97014, -0.24254)
	if !n.IsEqual(*expected) {
		t.Errorf("Expected normal = %v, but got %v", expected, n)
	}
}

func TestReflectVector(t *testing.T) {
	// Scenario: Reflecting a vector approaching at 45°
	// Given v ← vector(1, -1, 0)
	v := core.NewVector(1, -1, 0)
	// And n ← vector(0, 1, 0)
	n := core.NewVector(0, 1, 0)
	// When r ← reflect(v, n)
	r := lighting.Reflect(*v, *n)
	// Then r = vector(1, 1, 0)
	expected := core.NewVector(1, 1, 0)
	if !r.IsEqual(*expected) {
		t.Errorf("Expected reflect(v, n) = %v, but got %v", expected, r)
	}

	// Scenario: Reflecting a vector off a slanted surface
	// Given v ← vector(0, -1, 0)
	v = core.NewVector(0, -1, 0)
	// And n ← vector(√2/2, √2/2, 0)
	sqrtHalf := math.Sqrt(2) / 2
	n = core.NewVector(sqrtHalf, sqrtHalf, 0)
	// When r ← reflect(v, n)
	r = lighting.Reflect(*v, *n)
	// Then r = vector(1, 0, 0)
	expected = core.NewVector(1, 0, 0)
	if !r.IsEqual(*expected) {
		t.Errorf("Expected reflect(v, n) = %v, but got %v", expected, r)
	}
}

func TestLighting(t *testing.T) {
	// Setup common values for all scenarios
	m := material.DefaultMaterial()
	position := core.NewPoint(0, 0, 0)

	// Scenario: Lighting with the eye between the light and the surface
	eyev := core.NewVector(0, 0, -1)
	normalv := core.NewVector(0, 0, -1)
	light := lighting.NewLight(*color.NewColor(1, 1, 1), *core.NewPoint(0, 0, -10))
	result := lighting.Lighting(m, light, *position, *eyev, *normalv)
	expected := color.NewColor(1.9, 1.9, 1.9)
	if !result.IsEqual(*expected) {
		t.Errorf("Expected lighting result = %v, but got %v", expected, result)
	}

	// Scenario: Lighting with the eye between light and surface, eye offset 45°
	sqrtHalf := math.Sqrt(2) / 2
	eyev = core.NewVector(0, sqrtHalf, -sqrtHalf)
	normalv = core.NewVector(0, 0, -1)
	light = lighting.NewLight(*color.NewColor(1, 1, 1), *core.NewPoint(0, 0, -10))
	result = lighting.Lighting(m, light, *position, *eyev, *normalv)
	expected = color.NewColor(1.0, 1.0, 1.0)
	if !result.IsEqual(*expected) {
		t.Errorf("Expected lighting result = %v, but got %v", expected, result)
	}

	// Scenario: Lighting with eye opposite surface, light offset 45°
	eyev = core.NewVector(0, 0, -1)
	normalv = core.NewVector(0, 0, -1)
	light = lighting.NewLight(*color.NewColor(1, 1, 1), *core.NewPoint(0, 10, -10))
	result = lighting.Lighting(m, light, *position, *eyev, *normalv)
	expected = color.NewColor(0.7364, 0.7364, 0.7364)
	if !result.IsEqual(*expected) {
		t.Errorf("Expected lighting result = %v, but got %v", expected, result)
	}

	// Scenario: Lighting with eye in the path of the reflection vector
	eyev = core.NewVector(0, -sqrtHalf, -sqrtHalf)
	normalv = core.NewVector(0, 0, -1)
	light = lighting.NewLight(*color.NewColor(1, 1, 1), *core.NewPoint(0, 10, -10))
	result = lighting.Lighting(m, light, *position, *eyev, *normalv)
	expected = color.NewColor(1.6364, 1.6364, 1.6364)
	if !result.IsEqual(*expected) {
		t.Errorf("Expected lighting result = %v, but got %v", expected, result)
	}

	// Scenario: Lighting with the light behind the surface
	eyev = core.NewVector(0, 0, -1)
	normalv = core.NewVector(0, 0, -1)
	light = lighting.NewLight(*color.NewColor(1, 1, 1), *core.NewPoint(0, 0, 10))
	result = lighting.Lighting(m, light, *position, *eyev, *normalv)
	expected = color.NewColor(0.1, 0.1, 0.1)
	if !result.IsEqual(*expected) {
		t.Errorf("Expected lighting result = %v, but got %v", expected, result)
	}
}
