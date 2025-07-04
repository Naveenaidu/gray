package main

import (
	"fmt"
	"math"
	"testing"
)

func TestAddTuples(t *testing.T) {
	// Case 1: Adding vector to point gives a new point
	a1 := NewPoint(3, -2, 5)
	a2 := NewVector(-2, 3, 1)
	a3 := a1.AddVector(*a2)

	if !a3.IsEqual(*NewPoint(1, 1, 6)) {
		t.Errorf("got: %+v, want: (1, 1, 6, 1)", a3)
	}

	// Case 2: Adding two vectors gives a new vector
	c2v1 := NewVector(3, -2, 5)
	c2v2 := NewVector(-2, 3, 1)
	c2v3 := AddVectors([]Vector{*c2v1, *c2v2})
	if !c2v3.IsEqual(*NewVector(1, 1, 6)) {
		t.Errorf("got: %+v, want: (1, 1, 6, 0)", c2v3)
	}

}

func TestSubtractTuples(t *testing.T) {

	// Case 1: Subtract two points, gives a new vector
	c1p1 := NewPoint(3, 2, 1)
	c1p2 := NewPoint(5, 6, 7)
	c1v1 := c1p1.Subtract(*c1p2)

	if !c1v1.IsEqual(*NewVector(-2, -4, -6)) {
		t.Errorf("got: %+v, want: (-2, -4, -6, 0)", c1v1)
	}

	// Case 2: Subtract a vector from a point, gives a new point
	// This is equivalent to walking from a point in the direction of the vector
	c2p1 := NewPoint(3, 2, 1)
	c2v1 := NewVector(5, 6, 7)
	c2p2 := c2p1.SubtractVector(*c2v1)
	if !c2p2.IsEqual(*NewPoint(-2, -4, -6)) {
		t.Errorf("got: %+v, want: (-2, -4, -6, 1)", c2p2)
	}

	// Case 3: Subtract two vectors, gives a new vector
	c3v1 := NewVector(3, 2, 1)
	c3v2 := NewVector(5, 6, 7)
	c3v3 := SubtractVectors([]Vector{*c3v1, *c3v2})
	if !c3v3.IsEqual(*NewVector(-2, -4, -6)) {
		t.Errorf("got: %+v, want: (-2, -4, -6, 0)", c3v3)
	}

	// Case 4: Subtract a vector from itself, gives a zero vector
	c4v1 := NewVector(3, 2, 1)
	c4v2 := SubtractVectors([]Vector{*c4v1, *c4v1})
	if !c4v2.IsEqual(*NewVector(0, 0, 0)) {
		t.Errorf("got: %+v, want: (0, 0, 0, 0)", c4v2)
	}

}

func TestNegateVector(t *testing.T) {
	a1 := NewVector(1, -2, 3)

	if !a1.Negate().IsEqual(*NewVector(-1, 2, -3)) {
		t.Errorf("got: %+v, want: (-1, 2, -3, 0)", a1.Negate())
	}
}

func TestScalarMuliply(t *testing.T) {
	p1 := NewPoint(1, -2, 3)
	p2 := p1.ScalarMultiply(3)
	if !p2.IsEqual(*NewPoint(3, -6, 9)) {
		t.Errorf("got: %+v, want: (3, -6, 9, 1)", p2)
	}
	v1 := NewVector(1, -2, 3)
	v2 := v1.ScalarMultiply(3)
	if !v2.IsEqual(*NewVector(3, -6, 9)) {
		t.Errorf("got: %+v, want: (3, -6, 9, 0)", v2)
	}

}

func TestScalarDivide(t *testing.T) {
	p1 := NewPoint(1, -2, 3)
	p2 := p1.ScalarDivide(2)
	if !p2.IsEqual(*NewPoint(0.5, -1, 1.5)) {
		t.Errorf("got: %+v, want: (0.5, -1, 1.5, 1)", p2)
	}
	v1 := NewVector(1, -2, 3)
	v2 := v1.ScalarDivide(2)
	if !v2.IsEqual(*NewVector(0.5, -1, 1.5)) {
		t.Errorf("got: %+v, want: (0.5, -1, 1.5, 0)", v2)
	}

}

func TestVectorMagnitude(t *testing.T) {

	v := NewVector(1, 0, 0)
	if !isFloatEqual(v.Magnitude(), 1.0) {
		t.Errorf("got: %f, want: 1.0", v.Magnitude())
	}

	v = NewVector(0, 1, 0)
	if !isFloatEqual(v.Magnitude(), 1.0) {
		t.Errorf("got: %f, want: 1.0", v.Magnitude())
	}

	v = NewVector(0, 0, 1)
	if !isFloatEqual(v.Magnitude(), 1.0) {
		t.Errorf("got: %f, want: 1.0", v.Magnitude())
	}

	v = NewVector(1, 2, 3)
	if !isFloatEqual(v.Magnitude(), math.Sqrt(14)) {
		t.Errorf("got: %f, want: %f", v.Magnitude(), math.Sqrt(14))
	}

	v = NewVector(-1, -2, -3)
	if !isFloatEqual(v.Magnitude(), math.Sqrt(14)) {
		t.Errorf("got: %f, want: %f", v.Magnitude(), math.Sqrt(14))
	}

}

func TestNormalizeVector(t *testing.T) {
	var v, v_normalized *Vector

	v = NewVector(4, 0, 0)
	v_normalized = v.Normalize()
	if !isFloatEqual(v_normalized.x, 1.0) {
		t.Errorf("got: %f, want: 1.0", v_normalized.x)
	}
	if !isFloatEqual(v_normalized.y, 0.0) {
		t.Errorf("got: %f, want: 0.0", v_normalized.y)
	}
	if !isFloatEqual(v_normalized.z, 0.0) {
		t.Errorf("got: %f, want: 0.0", v_normalized.z)
	}

	v = NewVector(1, 2, 3)
	v_normalized = v.Normalize()
	if !isFloatEqual(v_normalized.x, 0.26726) {
		t.Errorf("got: %f, want: 0.26726", v_normalized.x)
	}
	if !isFloatEqual(v_normalized.y, 0.53452) {
		t.Errorf("got: %f, want: 0.53452", v_normalized.y)
	}
	if !isFloatEqual(v_normalized.z, 0.80178) {
		t.Errorf("got: %f, want: 0.80178", v_normalized.z)
	}

	// Check the magnitude of the normalized vector
	v = NewVector(1, 2, 3)
	if !isFloatEqual(v.Normalize().Magnitude(), 1.0) {
		t.Errorf("got: %f, want: 1.0", v.Normalize().Magnitude())
	}
}

func TestDotProduct(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)

	dotProduct := a.DotProduct(*b)

	if dotProduct != 20.0 {
		t.Errorf("got: %f, want: 20.0", dotProduct)
	}

}

func TestCrossProduct(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)

	abCrossProduct := a.CrossProduct(*b)

	if !abCrossProduct.IsEqual(*NewVector(-1, 2, -1)) {
		t.Errorf("got: %+v, want: (-1, 2, -1)", abCrossProduct)
	}

	baCrossProduct := b.CrossProduct(*a)

	if !baCrossProduct.IsEqual(*NewVector(1, -2, 1)) {
		t.Errorf("got: %+v, want: (-1, 2, -1)", baCrossProduct)
	}

}

/* ------------- Canvas --------------- */

func TestNewCanvas(t *testing.T) {
	canvas := NewCanvas(10, 20, *Black)

	for x := 0; x < canvas.width; x++ {
		for y := 0; y < canvas.height; y++ {
			if canvas.color[x][y] != *Black {
				t.Errorf("got: %+v, want: %+v", canvas.color[x][y], *Black)
			}
		}
	}
}

func TestCanvasWritePixel(t *testing.T) {
	canvas := NewCanvas(80, 80, *Black)
	canvas.WritePixel(2, 3, *Red)

	if canvas.PixelAt(2, 3) != *Red {
		t.Errorf("got: %+v, want: %+v", canvas.PixelAt(2, 3), *Red)
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

		Note that the columns are sent as arrays
	*/
	matrix := NewMatrix(4, 4, [][]float64{{1, 2, 3, 4}, {5.5, 6.5, 7.5, 8.5}, {9, 10, 11, 12}, {13.5, 14.5, 15.5, 16.5}})

	if matrix.value[0][0] != 1 {
		t.Errorf("got: %f, want: %f", matrix.value[0][0], 1.0)
	}

	if matrix.value[0][3] != 4 {
		t.Errorf("got: %f, want: %f", matrix.value[0][3], 4.0)
	}

	if matrix.value[1][0] != 5.5 {
		t.Errorf("got: %f, want: %f", matrix.value[1][0], 5.5)
	}

	if matrix.value[1][2] != 7.5 {
		t.Errorf("got: %f, want: %f", matrix.value[1][2], 7.5)
	}

	if matrix.value[2][2] != 11 {
		t.Errorf("got: %f, want: %f", matrix.value[2][2], 11.0)
	}

	if matrix.value[3][0] != 13.5 {
		t.Errorf("got: %f, want: %f", matrix.value[3][0], 13.5)
	}

	if matrix.value[3][2] != 15.5 {
		t.Errorf("got: %f, want: %f", matrix.value[3][2], 15.5)
	}
}

func Test2x2Matrix(t *testing.T) {
	matrix := NewMatrix(2, 2, [][]float64{{-3, 5}, {1, -2}})

	value := matrix.value[0][0]
	expectedValue := -3.0
	if value != expectedValue {
		t.Errorf("got: %f, want: %f", value, expectedValue)
	}

	value = matrix.value[0][1]
	expectedValue = 5.0
	if value != expectedValue {
		t.Errorf("got: %f, want: %f", value, expectedValue)
	}

	value = matrix.value[1][0]
	expectedValue = 1.0
	if value != expectedValue {
		t.Errorf("got: %f, want: %f", value, expectedValue)
	}

	value = matrix.value[1][1]
	expectedValue = -2.0
	if value != expectedValue {
		t.Errorf("got: %f, want: %f", value, expectedValue)
	}

}

func Test3x3Matrix(t *testing.T) {
	matrix := NewMatrix(3, 3, [][]float64{{-3, 5, 0}, {1, -2, -7}, {0, 1, 1}})

	value := matrix.value[0][0]
	expectedValue := -3.0
	if value != expectedValue {
		t.Errorf("got: %f, want: %f", value, expectedValue)
	}

	value = matrix.value[1][1]
	expectedValue = -2.0
	if value != expectedValue {
		t.Errorf("got: %f, want: %f", value, expectedValue)
	}

	value = matrix.value[2][2]
	expectedValue = 1.0
	if value != expectedValue {
		t.Errorf("got: %f, want: %f", value, expectedValue)
	}

}

func TestMatrixEquality(t *testing.T) {
	m1 := NewMatrix(4, 4, [][]float64{{1, 5, 9, 5}, {2, 6, 8, 4}, {3, 7, 7, 3}, {4, 8, 6, 2}})
	m2 := NewMatrix(4, 4, [][]float64{{1, 5, 9, 5}, {2, 6, 8, 4}, {3, 7, 7, 3}, {4, 8, 6, 2}})
	if !m1.IsEqual(*m2) {
		t.Errorf("expected matrices to be equal - but they are not")
		m1.PrintMatrix()
		fmt.Println("-----------")
		m2.PrintMatrix()

	}

	m3 := NewMatrix(4, 4, [][]float64{{1, 5, 9, 5}, {2, 6, 8, 4}, {3, 7, 7, 3}, {4, 8, 6, 2}})
	m4 := NewMatrix(4, 4, [][]float64{{2, 6, 8, 4}, {3, 7, 7, 3}, {1, 5, 9, 5}, {4, 8, 6, 2}})
	if m3.IsEqual(*m4) {
		t.Errorf("expected matrices to be not equal - but they are")
		m3.PrintMatrix()
		fmt.Println("-----------")
		m4.PrintMatrix()

	}

}

func TestMatrixMultiply_InvalidDimensions(t *testing.T) {
	m1 := NewMatrix(2, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	m2 := NewMatrix(2, 2, [][]float64{
		{7, 8},
		{9, 10},
	})

	result := m1.Multiply(*m2)

	if result.rows != 1 || result.columns != 1 || math.IsNaN(result.value[0][0]) == false {
		t.Errorf("Expected NaNMatrix, but got %v", result.value)
	}
}

func TestMatrixMultiply(t *testing.T) {
	m1 := NewMatrix(4, 4, [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})

	m2 := NewMatrix(4, 4, [][]float64{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	})

	result := m1.Multiply(*m2)

	expectedMatrix := NewMatrix(4, 4, [][]float64{
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
	m1 := NewMatrix(4, 4, [][]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	})

	result := m1.MultiplyTuple([4]float64{1, 2, 3, 1})

	expectedMatrix := NewMatrix(4, 1, [][]float64{
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
	m1 := NewMatrix(4, 4, [][]float64{
		{1, 2, 3, 5},
		{4, 5, 6, 0},
		{7, 8, 9, 16},
		{-2, 1, 4.4, 11},
	})

	expected := m1

	result := m1.Multiply(*IdentityMatrix())

	if !result.IsEqual(*expected) {
		t.Errorf("Expected %v, but got %v", expected.value, result.value)
	}
}

func TestMatrixMultiply_Transpose(t *testing.T) {
	m1 := NewMatrix(4, 4, [][]float64{
		{0, 9, 3, 0},
		{9, 8, 0, 8},
		{1, 8, 5, 3},
		{0, 0, 5, 8},
	})

	expected := NewMatrix(4, 4, [][]float64{
		{0, 9, 1, 0},
		{9, 8, 8, 0},
		{3, 0, 5, 5},
		{0, 8, 3, 8},
	})

	result := m1.Transpose()

	if !result.IsEqual(*expected) {
		t.Errorf("Expected %v, but got %v", expected.value, result.value)
	}
}

func TestMatrixMultiply_TransposeIdentity(t *testing.T) {
	m1 := IdentityMatrix()

	result := m1.Transpose()

	if !result.IsEqual(*IdentityMatrix()) {
		t.Errorf("Expected %v, but got %v", IdentityMatrix().value, result.value)
	}
}

func TestMatrix_2x2Determinant(t *testing.T) {
	m := NewMatrix(2, 2, [][]float64{
		{1, 5},
		{-3, 2},
	})
	result := Determinant2(*m)
	expected := 17.0
	if !isFloatEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestMatrix_3x3SubMatrix(t *testing.T) {
	m1 := NewMatrix(3, 3, [][]float64{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3},
	})

	expected := NewMatrix(2, 2, [][]float64{
		{-3, 2},
		{0, 6},
	})

	result := m1.SubMatrix(0, 2)

	if !result.IsEqual(*expected) {
		t.Errorf("Expected %v, but got %v", expected.value, result.value)
	}
}

func TestMatrix_4x4SubMatrix(t *testing.T) {
	m1 := NewMatrix(4, 4, [][]float64{
		{1, 2, 3, 4},
		{2, 4, -4, 2},
		{8, 6, 4, 1},
		{0, 10, 11, 1},
	})

	expected := NewMatrix(3, 3, [][]float64{
		{1, 3, 4},
		{2, -4, 2},
		{0, 11, 1},
	})

	result := m1.SubMatrix(2, 1)

	if !result.IsEqual(*expected) {
		t.Errorf("Expected %v, but got %v", expected.value, result.value)
	}
}

func TestMatrix_3x3Minor(t *testing.T) {
	m := NewMatrix(3, 3, [][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})
	result := Minor3(*m, 1, 0)
	expected := 25.0
	if !isFloatEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestMatrix_3x3Cofactor(t *testing.T) {
	var result, expected float64
	m := NewMatrix(3, 3, [][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})
	result = Cofactor3(*m, 0, 0)
	expected = -12.0
	if !isFloatEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	result = Cofactor3(*m, 1, 0)
	expected = -25.0
	if !isFloatEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestMatrix_3x3Determinant(t *testing.T) {
	// Given the following 3x3 matrix A
	matrix := NewMatrix(3, 3, [][]float64{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	})

	// Test cofactor(A, 0, 0)
	result := Cofactor3(*matrix, 0, 0)
	expected := 56.0
	if !isFloatEqual(result, expected) {
		t.Errorf("Expected cofactor(A, 0, 0) = %v, but got %v", expected, result)
	}

	// Test cofactor(A, 0, 1)
	result = Cofactor3(*matrix, 0, 1)
	expected = 12.0
	if !isFloatEqual(result, expected) {
		t.Errorf("Expected cofactor(A, 0, 1) = %v, but got %v", expected, result)
	}

	// Test cofactor(A, 0, 2)
	result = Cofactor3(*matrix, 0, 2)
	expected = -46.0
	if !isFloatEqual(result, expected) {
		t.Errorf("Expected cofactor(A, 0, 2) = %v, but got %v", expected, result)
	}

	// Test determinant(A)
	result = Determinant3(*matrix)
	expected = -196.0
	if !isFloatEqual(result, expected) {
		t.Errorf("Expected determinant(A) = %v, but got %v", expected, result)
	}
}

func TestMatrix_4x4Determinant(t *testing.T) {
	// Given the following 4x4 matrix A
	matrix := NewMatrix(4, 4, [][]float64{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	})

	// Test cofactor(A, 0, 0)
	// subM1 := matrix.SubMatrix(0, 0)
	result := Cofactor4(*matrix, 0, 0)
	expected := 690.0
	if !isFloatEqual(result, expected) {
		t.Errorf("Expected cofactor(A, 0, 0) = %v, but got %v", expected, result)
	}

	// Test cofactor(A, 0, 1)
	result = Cofactor4(*matrix, 0, 1)
	expected = 447.0
	if !isFloatEqual(result, expected) {
		t.Errorf("Expected cofactor(A, 0, 1) = %v, but got %v", expected, result)
	}

	// Test cofactor(A, 0, 2)
	result = Cofactor4(*matrix, 0, 2)
	expected = 210.0
	if !isFloatEqual(result, expected) {
		t.Errorf("Expected cofactor(A, 0, 2) = %v, but got %v", expected, result)
	}

	// Test cofactor(A, 0, 3)
	result = Cofactor4(*matrix, 0, 3)
	expected = 51.0
	if !isFloatEqual(result, expected) {
		t.Errorf("Expected cofactor(A, 0, 3) = %v, but got %v", expected, result)
	}

	// Test determinant(A)
	result = matrix.Determinant4()
	expected = -4071.0
	if !isFloatEqual(result, expected) {
		t.Errorf("Expected determinant(A) = %v, but got %v", expected, result)
	}
}

func TestMatrix_InvertibleMatrix(t *testing.T) {
	// Scenario: Testing an invertible matrix for invertibility
	// Det is -2120.0
	matrix := NewMatrix(4, 4, [][]float64{
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
	matrix2 := NewMatrix(4, 4, [][]float64{
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
	matrix := NewMatrix(4, 4, [][]float64{
		{8, -5, 9, 2},
		{7, 5, 6, 1},
		{-6, 0, 9, 6},
		{-3, 0, -9, -4},
	})

	// Expected inverse(A)
	expected := NewMatrix(4, 4, [][]float64{
		{-0.15385, -0.15385, -0.28205, -0.53846},
		{-0.07692, 0.12308, 0.02564, 0.03077},
		{0.35897, 0.35897, 0.43590, 0.92308},
		{-0.69231, -0.69231, -0.76923, -1.92308},
	})

	// Calculate inverse(A)
	result := matrix.Inverse()

	// Compare the result with the expected matrix
	if !result.IsEqual(*expected) {
		t.Errorf("Expected inverse(A) = %v, but got %v", expected.value, result.value)
	}

	// Given the following 4x4 matrix A
	matrix1 := NewMatrix(4, 4, [][]float64{
		{9, 3, 0, 9},
		{-5, -2, -6, -3},
		{-4, 9, 6, 4},
		{-7, 6, 6, 2},
	})

	// Expected inverse(A)
	expected1 := NewMatrix(4, 4, [][]float64{
		{-0.04074, -0.07778, 0.14444, -0.22222},
		{-0.07778, 0.03333, 0.36667, -0.33333},
		{-0.02901, -0.14630, -0.10926, 0.12963},
		{0.17778, 0.06667, -0.26667, 0.33333},
	})

	// Calculate inverse(A)
	result1 := matrix1.Inverse()

	// Compare the result with the expected matrix
	if !result1.IsEqual(*expected1) {
		t.Errorf("Expected inverse(A) = %v, but got %v", expected1.value, result1.value)
	}
}

func TestMatrix_ProductWithInverse(t *testing.T) {
	// Test that if "C ← A * B" then  "C * (1/B) = A"
	// Given the following 4x4 matrix A
	matrixA := NewMatrix(4, 4, [][]float64{
		{3, -9, 7, 3},
		{3, -8, 2, -9},
		{-4, 4, 4, 1},
		{-6, 5, -1, 1},
	})

	// And the following 4x4 matrix B
	matrixB := NewMatrix(4, 4, [][]float64{
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
		t.Errorf("Expected C * inverse(B) = A, but got %v", result.value)
	}
}
