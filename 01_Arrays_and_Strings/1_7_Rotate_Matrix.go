// Problem: Rotate NxN Matrix in-place: image pixel = 4 bytes, rotate = 90o degree
// Hint:
// #51: rotate layer by layer
// #100: Swapping value of 4 arrays
package main

import (
	"fmt"
	"reflect"
)

func rotateMatrix(matrix [][]int) {
	total_rows := len(matrix)
	total_cols := total_rows
	temp_matrix := make([][]int, total_rows)
	for row := range temp_matrix {
		temp_matrix[row] = make([]int, total_cols)
	}

	for row := range matrix {
		fmt.Printf("Row:%v\n", matrix[row])
		for col := range matrix[row] {
			temp_matrix[col][total_rows-1-row] = matrix[row][col]
			fmt.Printf("rơw: %d col:%d -> %d => matrix[%d][%d]:%d\n", row, col, matrix[row][col], col, total_rows-1-row, temp_matrix[col][total_rows-1-row])
		}
	}

	for row := range matrix {
		matrix[row] = temp_matrix[row]
	}

	fmt.Printf("new matrix:%v\n", matrix)

}

// Implement solution in book
func rotateMatrixLayer(matrix [][]int) {
	n := len(matrix)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		fmt.Printf("layer:%d/%d first:%d last:%d\n", layer, n/2, first, last)
		for i := first; i < last; i++ {
			offset := i - first
			fmt.Printf("i:%d offset:%d\n", i, offset)
			top := matrix[first][i] // save top
			// left -> top
			matrix[first][i] = matrix[last-offset][first]
			// bottom -> left
			matrix[last-offset][first] = matrix[last][last-offset]
			// right -> bottom
			matrix[last][last-offset] = matrix[i][last]
			// top -> right
			matrix[i][last] = top

		}
	}
}

func mainRotateMatrix() {
	cases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			[][]int{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
		},
	}
	for _, c := range cases {
		fmt.Printf("Input: %v, expected: %v\n", c.input, c.expected)
		rotateMatrixLayer(c.input)
		fmt.Printf("Rotated matrix: %v\n", c.input)
		if !reflect.DeepEqual(c.input, c.expected) {
			fmt.Printf("Expected: %v, actual: %v\n", c.expected, c.input)
		} else {
			fmt.Println("OK")
		}
	}
}
