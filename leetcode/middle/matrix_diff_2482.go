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
	//new_row, new_col := 0, 0
	for i, row := range matrix {
		var resRow []int
		for j, _ := range row {
			resRow = append(resRow, matrix[len(row)-j-1][i])
		}
		result = append(result, resRow)
	}
	//for i, row := range matrix {
	//	new_col = i
	//	for j, _ := range row {
	//		new_row = j
	//	}
	//}
	//for i := 0; i < new_row; i++ {
	//	var nr []int
	//	for j := 0; j < new_col; j++ {
	//		nr = append(nr, matrix[j][i])
	//	}
	//	result = append(result, nr)
	//}
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
	//grid := [][]int{
	//	{1, 1, 1},
	//	{1, 1, 1},
	//}
	gridBigger := [][]int{
		{0, 1, 1},
		{1, 0, 1},
		{0, 0, 1},
	}
	//fmt.Println(onesMinusZeros(grid))
	fmt.Println(rotateMatrix(gridBigger))
	fmt.Println(gridBigger)
	fmt.Println(onesMinusZeros(gridBigger))
}
