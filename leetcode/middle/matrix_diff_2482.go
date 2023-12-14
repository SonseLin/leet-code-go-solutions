package middle

import "fmt"

func getNumsFromRow(row []int) (zeroes, ones int) {
	for _, elem := range row {
		if elem == 1 {
			ones++
		} else {
			zeroes++
		}
	}
	return
}

func rotateMatrix(matrix [][]int) [][]int {
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

func onesMinusZeros(grid [][]int) [][]int {
	var result [][]int
	rotated := rotateMatrix(grid)
	for _, row := range grid {
		zR, oR := getNumsFromRow(row)
		var line []int
		for c, _ := range rotated {
			zC, oC := getNumsFromRow(rotated[c])
			line = append(line, oR+oC-zR-zC)
		}
		result = append(result, line)
	}
	return result
}

func TestOMZ() {
	grid := [][]int{
		{1, 1, 1},
		{1, 1, 1},
	}
	gridBigger := [][]int{
		{0, 1, 1},
		{1, 0, 1},
		{0, 0, 1},
	}
	fmt.Println(onesMinusZeros(grid))
	fmt.Println(onesMinusZeros(gridBigger))
}
