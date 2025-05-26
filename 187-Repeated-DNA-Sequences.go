package main

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 11 {
		return []string{}
	}

	var a int

	m := make(map[int]int, len(s)-9)

    // fill the window
	for i := 0; i <= 9; i++ {
		a = (a << 2) | toBits(s[i])
	}
	m[a]++

    // slide
	for i := 10; i < len(s); i++ {
		a = (a << 2 & 1048575) | toBits(s[i])
		m[a]++
	}

	res := make([]string, 0, len(m))
	for k, v := range m {
		if v > 1 {
			res = append(res, toString(k))
		}
	}

	return res
}

func toBits(b byte) int {
	return int(b & 6 >> 1)
}

func toString(x int) string {
	res := make([]byte, 0, 10)
	for i := 9; i >= 0; i-- {
		v := 3 << (i * 2) & x >> (i * 2)
		switch v {
		case 0:
			res = append(res, 65)
		case 1:
			res = append(res, 67)
		case 2:
			res = append(res, 84)
		case 3:
			res = append(res, 71)
		}
	}
	return string(res)
}