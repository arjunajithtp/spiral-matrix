package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateMatrix(1, 1))
}

const (
	leftToRight = 0
	topToBottom = 1
	rightToLeft = 2
	bottomToTop = 3
)

func generateMatrix(colLen, rowLen int) [][]int {
	matrix := make([][]int, colLen)
	for i := 0; i < len(matrix); i++ {
		row := make([]int, rowLen)
		matrix[i] = row
	}

	top, left := 0, 0
	bottom, right := colLen-1, rowLen-1
	dir := leftToRight
	value := 1

	for top <= bottom && left <= right {
		switch dir {
		case leftToRight:
			for i := left; i <= right; i++ {
				matrix[top][i] = value
				value++
			}
			top++
			dir = topToBottom
		case topToBottom:
			for i := top; i <= bottom; i++ {
				matrix[i][right] = value
				value++
			}
			right--
			dir = rightToLeft
		case rightToLeft:
			for i := right; i >= left; i-- {
				matrix[bottom][i] = value
				value++
			}
			bottom--
			dir = bottomToTop
		case bottomToTop:
			for i := bottom; i >= top; i-- {
				matrix[i][left] = value
				value++
			}
			left++
			dir = leftToRight
		}
	}
	return matrix
}