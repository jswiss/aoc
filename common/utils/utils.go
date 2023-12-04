package utils

func SliceSum(sl []int) int {
	sum := 0
	for i := 0; i < len(sl); i++ {
		sum += sl[i]
	}
	return sum
}

// Ints returns a unique subset of the int slice provided.
func Ints(input []int) []int {
	u := make([]int, 0, len(input))
	m := make(map[int]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}
