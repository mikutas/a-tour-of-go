package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var matrix [][]uint8
	var row []uint8
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			row = append(row, uint8(i^j))
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func main() {
	pic.Show(Pic)
}
