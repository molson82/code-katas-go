package consecutiveStrings

func LongestConsec(strarr []string, k int) string {
	var maxlen int
	var maxWord string
	for i := 0; i <= len(strarr)-k; i++ {
		var compareStr string
		for j := 0; j < k; j++ {
			compareStr += strarr[i+j]
		}
		if len(compareStr) > maxlen {
			maxlen = len(compareStr)
			maxWord = compareStr
		}
	}

	return maxWord
}
