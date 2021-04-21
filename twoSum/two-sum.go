package twoSum

func TwoSum(numbers []int, target int) [2]int {
  for i:=0; i<len(numbers); i++ {
    for j:=i; j<len(numbers); j++ {
      if i == j {
        continue
      }
      if target == numbers[i] + numbers[j] {
        return [2]int{i,j}
      }
    }
  }

  return [2]int{}
}
