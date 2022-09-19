package main

// 1.7 Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is
// 4 bytes, write a method to rotate the image by 90 degrees. Can you do this in place?
func rotateMatrix() [][]string {
	matrix := [][]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
	}
	newMatrix := make([][]string, 4)
	for i := range newMatrix {
		newMatrix[i] = make([]string, 4)
	}

	for i := range matrix {
		for j := range matrix[i] {
			newMatrix[j][len(matrix)-i-1] = matrix[i][j]
		}
	}
	return newMatrix
}

// 1.8 Zero Matrix. Write an algorithm such that if an element in an MxN matrix is 0,
// its entire row and column set to 0

func zeroMatrix(row, column int) [][]int8 {
	var matrix = make([][]int8, row)
	for i := range matrix {
		matrix[i] = make([]int8, column)
	}

	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = 0
		}
	}

	return matrix
}
