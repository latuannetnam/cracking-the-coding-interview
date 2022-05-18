// Problem: Give MxN matrix, check if item in matrix = 0 => Row & Col of item = 0
// Hint:
// #17: Try finding the cells with zeros first before making any change to matrix
// #74: Optmize memory space = O(N) instead of O(N^2) -> optmize list of cells are zero
// #102: Use memory space of O(1) to store list of rows,cols to be zeroed -> use matrix itself to for data storage
package main

import (
	"fmt"
	"reflect"
)

func ZeroMatrix(matrix [][]int) {
	zero_rows := map[int]bool{}
	zero_cols := map[int]bool{}
	// Find rows and cols of zero cells
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == 0 {
				// zero_rows = append(zero_rows, row)
				// zero_cols = append(zero_cols, col)
				if !zero_rows[row] {
					zero_rows[row] = true
				}
				if !zero_cols[col] {
					zero_cols[col] = true
				}

			}
		}
	}

	fmt.Printf("Zero rows:%v zero cols:%v\n", zero_rows, zero_cols)

	// Zero all rows in list
	for row := range zero_rows {
		for col := range matrix[row] {
			matrix[row][col] = 0
		}
	}

	// Zero all cols in list
	for row := range matrix {
		for col := range zero_cols {
			matrix[row][col] = 0
		}
	}

}
func mainZeroMatrix() {
	cases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 0, 6},
				{7, 8, 9},
			},
			[][]int{
				{1, 0, 3},
				{0, 0, 0},
				{7, 0, 9},
			},
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 0},
			},
			[][]int{
				{1, 2, 0},
				{4, 5, 0},
				{0, 0, 0},
			},
		},
		{
			[][]int{
				{1, 2, 3},
				{0, 5, 6},
				{7, 8, 9},
			},
			[][]int{
				{0, 2, 3},
				{0, 0, 0},
				{0, 8, 9},
			},
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 0, 6},
				{7, 8, 0},
			},
			[][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			[][]int{
				{1, 0, 3},
				{4, 0, 6},
				{7, 8, 9},
			},
			[][]int{
				{0, 0, 0},
				{0, 0, 0},
				{7, 0, 9},
			},
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
	}
	for _, c := range cases {
		fmt.Printf("Input: %v, expected: %v\n", c.input, c.expected)
		ZeroMatrix(c.input)
		fmt.Printf("Zero matrix: %v\n", c.input)
		if !reflect.DeepEqual(c.input, c.expected) {
			fmt.Printf("Expected: %v, actual: %v\n", c.expected, c.input)
		} else {
			fmt.Println("OK!")
		}
	}
}
