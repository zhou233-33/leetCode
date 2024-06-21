package array

func GenerateMatrix(n int) [][]int {
	var direction = [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	x, y := 0, -1
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	var step = 1
	var directionIndex = 0
	for step <= n*n {
		nextX := x + direction[directionIndex][0]
		nextY := y + direction[directionIndex][1]
		if nextX >= 0 && nextX < n && nextY >= 0 && nextY < n && matrix[nextX][nextY] == 0 {
			matrix[nextX][nextY] = step
			x, y = nextX, nextY
			step++
		} else {
			directionIndex = (directionIndex + 1) % 4
		}
	}
	return matrix
}
