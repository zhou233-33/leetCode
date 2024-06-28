package string

func StrStr(haystack string, needle string) int {
	next := GetNext(needle)
	nByte := []byte(needle)
	hByte := []byte(haystack)
	j := -1
	for i := 0; i < len(haystack); i++ {
		for j > -1 && nByte[j+1] != hByte[i] {
			j = next[j]
		}
		if nByte[j+1] == hByte[i] {
			j++
		}
		if j == len(needle)-1 {
			return i - len(needle) + 1
		}
	}
	return -1
}
func GetNext(s string) []int {
	//next中的值表示当前的值和从头开始的第几个值是相同的。
	//因此当匹配的时候有值不相等的时候,需要拿它之前一位的next值。跳转到对应值后，下次比较从该值后面的一个数进行比较。
	//j表示已经匹配相等的最后一位
	j := -1
	next := make([]int, len(s))
	byteS := []byte(s)
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > -1 && byteS[i] != byteS[j+1] {
			j = next[j]
		}
		if byteS[j+1] == byteS[i] {
			j++
		}
		next[i] = j
	}
	return next
}

func GetRepeatStr(s string) bool {
	next := GetNext(s)
	n := len(s)
	if next[n-1] != -1 && n%(n-1-next[n-1]) == 0 {
		return true
	}
	return false
}
