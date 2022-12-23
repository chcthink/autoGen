package cmd

import (
	"errors"
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

func IsSupportType(dataType string) (ts ColumnType, err error) {
	ts, ok := sql2tsType[dataType]
	if !ok {
		err = errors.New("暂不支持" + dataType)
		return
	}
	return
}
