package easy

import "fmt"

func transpose(matrix [][]int) [][]int {
	var result [][]int
	// cols matrix
	rows := len(matrix[0])
	// rows matrix
	cols := len(matrix)
	for r := 0; r < rows; r++ {
		var row []int
		for c := 0; c < cols; c++ {
			row = append(row, matrix[cols-c-1][r])
		}
		result = append(result, row)
	}
	return result
}

func reverse(matrix [][]int) [][]int {
	for i, _ := range matrix {
		for c := 0; c < len(matrix[0])/2; c++ {
			tmp := matrix[i][c]
			matrix[i][c] = matrix[i][len(matrix[0])-1-c]
			matrix[i][len(matrix[0])-1-c] = tmp
		}
	}
	return matrix
}

func TestTP() {
	matr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(reverse(transpose(matr)))
}
