package stringsplit

import "regexp"

func StringSplit(str string) []string {
	out := []string{}
	for i := 0; i < len(str); i = i + 2 {
		tok := string(str[i])
		if len(str) > i+1 {
			tok += string(str[i+1])
		} else {
			tok += "_"
		}
		out = append(out, tok)
	}

	return out
}

func Solution(str string) []string {
	var res []string
	if len(str)%2 != 0 {
		str += "_"
	}
	for i := 0; i < len(str); i += 2 {
		res = append(res, str[i:i+2])
	}
	return res
}

func RegexpSolution(str string) []string {
	return regexp.MustCompile(".{2}").FindAllString(str+"_", -1)
}
