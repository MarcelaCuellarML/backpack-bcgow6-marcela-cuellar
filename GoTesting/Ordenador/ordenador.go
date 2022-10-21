package ordenador

func SliceOrdenador(nums []int) []int {
	sliceNumbers := nums
	n := len(sliceNumbers)
	for {
		organizado := false
		for i := 1; i < n; i++ {
			if sliceNumbers[i-1] > sliceNumbers[i] {
				temp := sliceNumbers[i]
				sliceNumbers[i] = sliceNumbers[i-1]
				sliceNumbers[i-1] = temp
				organizado = true

			}
		}
		if !organizado {
			break
		}
	}
	return sliceNumbers
}
