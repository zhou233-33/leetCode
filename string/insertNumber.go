package string

//将字符串的数字替换成“number”字符串

func InsertNumberString(s string) string {
	insertNumber := []byte{'n', 'u', 'm', 'b', 'e', 'r'}
	sByte := []byte(s)
	for i := 0; i < len(sByte); i++ {
		if sByte[i] >= '0' && sByte[i] <= '9' {
			sByte = append(sByte[:i], append(insertNumber, sByte[i+1:]...)...)
		}
	}
	return string(sByte)
}
