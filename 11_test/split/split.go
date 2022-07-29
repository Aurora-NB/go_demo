package split

import "strings"

func Split(s, seq string) []string {
	count := strings.Count(s, seq)
	res := make([]string, 0, count+1)
	if len(seq) == 0 {
		res = append(res, s)
		return res
	}
	for index := strings.Index(s, seq); index >= 0; index = strings.Index(s, seq) {
		r := s[:index]
		if len(r) != 0 {
			res = append(res, r)
		}
		s = s[index+len(seq):]
	}
	if len(s) > 0 {
		res = append(res, s)
	}
	return res
}

func Add(a, b int) int {
	return a + b
}
