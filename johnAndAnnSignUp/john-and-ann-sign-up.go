package johnAndAnnSignUp

import "fmt"

func Ann(n int) []int {
	// your code
	var days []int
	for i := 1; i <= n/2; i++ {
		days = append(days, i)
		days = append(days, i)
	}
	return days
}
func John(n int) []int {
	// your code
	var days []int
	for i := 0; i <= (n/2)+1; i++ {
    if i % 2 == 0 {
      days = append(days, i)
		  days = append(days, i)
    } else {
      days = append(days, i)
    }
		
	}
	return days
}
func SumJohn(n int) int {
	// your code
  var sum int
  fmt.Printf("DEBUG: %v\n", John(n))
  for _,v := range John(n) {
    sum += v
  }
	return sum
}
func SumAnn(n int) int {
	// your code
  var sum int
  for _,v := range Ann(n) {
    sum += v
  }
	return sum
}
