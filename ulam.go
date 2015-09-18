package main

import (
	"flag"
	"fmt"
	"math"
)

var (
	matrix     [][]int // there is no spoon
	startSize  = 2     // we'll start with a 2x2
	targetSize int     // how big our matrix should be
)

func main() {
	// get target matrix size
	flag.IntVar(&targetSize, "d", 2, "Spiral size")
	flag.Parse()

	// sanity
	if targetSize < 2 {
		if targetSize == 1 {
			fmt.Println("1")
		}
		return
	}

	// allocate matrix
	matrix = make([][]int, startSize, targetSize)
	for i := range matrix {
		matrix[i] = make([]int, startSize, targetSize)
	}

	// init matrix
	initialize()

	// 1. add top row
	// 2. rotate
	// 3. wash/rinse/repeat
	for i := 0; i < (targetSize-2)*2; i++ {
		shift()
		rotate()
	}

	// print it
	printMatrix()
}

func printMatrix() {
	// TODO: implement stringer interface
	for _, r := range matrix {
		fmt.Println(r)
	}
}

// hard-coded 2x2 spiral
func initialize() {
	// 4 1
	// 3 2
	matrix[0][0] = 4
	matrix[0][1] = 1
	matrix[1][0] = 3
	matrix[1][1] = 2
}

// shift matrix down one row and add another to the "top"
func shift() {
	numRows := len(matrix)
	numCols := len(matrix[0])

	// create new "bottom" row (copy of existing "bottom" row)
	newRow := make([]int, numCols)
	for i := 0; i < numCols; i++ {
		newRow[i] = matrix[numRows-1][i]
	}

	// append new row
	matrix = append(matrix, newRow)

	// iterate from bottom row and copy to row above
	for i := numRows; i > 0; i-- {
		for j, v := range matrix[i-1] {
			matrix[i][j] = v
		}
	}

	// add new "top" row starting with the max value of the existing
	// "top" row + 1
	nextVal := int(math.Max(float64(matrix[1][0]), float64(matrix[1][len(matrix[1])-1]))) + 1
	for i := 0; i < numCols; i++ {
		matrix[0][i] = nextVal + i
	}
}

// rotate matrix counter-clockwise
func rotate() {
	// allocate newMatrix with swapped # of rows/cols
	numRows := len(matrix)
	numCols := len(matrix[0])
	newMatrix := make([][]int, numCols, targetSize)
	for i := range newMatrix {
		newMatrix[i] = make([]int, numRows, targetSize)
	}

	// flippity-do-daw
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			newMatrix[numCols-j-1][i] = matrix[i][j]
		}
	}

	// assign new matrix
	matrix = newMatrix
}
