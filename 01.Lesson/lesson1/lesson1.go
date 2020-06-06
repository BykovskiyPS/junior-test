package lesson1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // map number in pos
	result := [2]int{0, 0}
	for index, value := range nums {
		newIndex, ok := m[target-value]
		if ok {
			result[0] = newIndex
			result[1] = index
			return result[:]
		}
		m[value] = index
	}
	return result[:]
}
