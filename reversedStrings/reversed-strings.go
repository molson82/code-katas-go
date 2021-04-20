package reversedStrings

import "strings"

func Solution(word string) string {
	strSlice := strings.Split(word, "")
	var newStr string
	for i := len(strSlice) - 1; i >= 0; i-- {
		newStr = newStr + strSlice[i]
	}

	return newStr
}
