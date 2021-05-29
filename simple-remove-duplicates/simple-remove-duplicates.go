package simpleRemoveDuplicates

func Solve(arr []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range arr {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
