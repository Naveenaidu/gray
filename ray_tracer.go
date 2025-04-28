package main

import (
	"fmt"

	"rsc.io/quote"
)

type Tuple struct {
	x, y, z, w float32
}

func point(x float32, y float32, z float32) Tuple {
	return Tuple{x, y, z, 1.0}
}

func isPoint(t Tuple) bool {
	return t.w == 1.0
}

func vector(x float32, y float32, z float32) Tuple {
	return Tuple{x, y, z, 0.0}
}

func isVector(t Tuple) bool {
	return t.w == 0.0
}

func main() {
	fmt.Println(quote.Glass())
}
