package cmd

import (
	"strings"
	"unicode"
)

func toCamel(s string) (tName string) {
	sList := strings.Split(s, "_")
	for i := range sList {
		var r = []rune(sList[i])
		if sList[i] == "id" {
			tName += "ID"
		} else if unicode.IsLower(r[0]) {
			r[0] -= 32
			tName += string(r)
		}
	}
	return
}
