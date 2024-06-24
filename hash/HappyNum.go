package hash

import "fmt"

func IsHappyNum(n int) bool {
	ExistNum := map[int]bool{}
	for !ExistNum[n] {
		ExistNum[n] = true
		fmt.Println(n)
		n = CalculateNum(n)
		if n == 1 {
			return true
		}
	}
	return false
}
func CalculateNum(n int) int {
	result := 0
	for n != 0 {
		a := n % 10
		n = n / 10
		result += a * a
	}
	return result
}
