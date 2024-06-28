package string

import "fmt"

//ReverseString翻转字符串
//ReverseStringFrontK翻转前k个字符串

// 双指针遍历，实现O(1)
func ReverseString(s []byte) []byte {
	left := 0
	right := len(s) - 1
	for left < right {
		s[right], s[left] = s[left], s[right]
		left++
		right--
	}
	return s
}

func ReverseStringFrontK(s string, k int) string {
	reverse := func(s string, front int, end int) string {
		tmp := []byte(s)
		for front < end {
			tmp[front], tmp[end] = tmp[end], tmp[front]
			front++
			end--
		}
		return string(tmp)
	}
	start := 0
	middle := k - 1
	end := 2*k - 1
	for start < len(s) {
		if middle >= len(s) {
			middle = len(s)
		}
		if end >= len(s) {
			end = len(s)
		}
		s = reverse(s, start, middle)
		start += 2 * k
		middle += 2 * k
		end += 2 * k
	}
	return s
}

func ReverseWordInString(s string) string {
	sByte := []byte(s)
	slow := 0
	for i := 0; i < len(s); i++ {
		if sByte[i] != ' ' {
			//添加一个空格，接下来继续交换
			if slow != 0 {
				sByte[slow] = ' '
				slow++
			}
			for i < len(s) && sByte[i] != ' ' {
				sByte[i], sByte[slow] = sByte[slow], sByte[i]
				i++
				slow++
			}
		}
	}

	reverse := func(s []byte, front, end int) []byte {
		tmp := s
		for front < end {
			tmp[front], tmp[end] = tmp[end], tmp[front]
			front++
			end--
		}
		return tmp
	}
	sByte = sByte[:slow]
	fmt.Printf("space delete:%v\n", string(sByte))
	sByte = reverse(sByte, 0, len(sByte)-1)
	fmt.Printf("reverse :%v\n", string(sByte))
	front := 0
	for i := 0; i <= len(sByte); i++ {
		if i == len(sByte) || sByte[i] == ' ' {
			sByte = reverse(sByte, front, i-1)
			front = i + 1
		}
	}
	return string(sByte)
}

func RightReverseString(s string, k int) string {
	sByte := []byte(s)
	reverse := func(s []byte, front, end int) []byte {
		tmp := s
		for front < end {
			tmp[front], tmp[end] = tmp[end], tmp[front]
			front++
			end--
		}
		return tmp
	}
	sByte = reverse(sByte, 0, len(sByte)-1-k)
	sByte = reverse(sByte, len(sByte)-k, len(sByte)-1)
	sByte = reverse(sByte, 0, len(sByte)-1)
	return string(sByte)
}
