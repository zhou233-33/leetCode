package hash

func IsAnagram(s string, t string) bool {
	letterCount := map[int32]int{}
	for _, l := range s {
		letterCount[l] = letterCount[l] + 1
	}
	for _, l := range t {
		letterCount[l] = letterCount[l] - 1
	}
	for _, val := range letterCount {
		if val != 0 {
			return false
		}
	}
	return true
}
