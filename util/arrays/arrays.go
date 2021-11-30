package arrays

func AppendN(arr []int, num, n int) []int {
	for i := 0; i < n; i++ {
		arr = append(arr, num)
	}
	return arr
}

func Contains(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}
