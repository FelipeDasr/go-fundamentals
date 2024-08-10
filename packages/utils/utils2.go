package utils

// Sum four numbers
func SumFourNumbers(x int, y int, z int, a int) int {
	return SumToNumbers(x, y) + SumToNumbers(z, a)
}