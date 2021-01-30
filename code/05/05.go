package main

// 替换空格
import (
	"strings"
)

func replaceSpace(s string) string {
	var str strings.Builder
	for _, i := range s {
		if string(i) == " " {
			str.WriteString("%20")
		} else {
			str.WriteString(string(i))
		}
	}
	return str.String()
}

func replaceSpace2(s string) string {
	var str string = ""
	for _, v := range s {
		if v == ' ' {
			str += "%20"
		} else {
			str += string(v)
		}
	}
	return str
}
