package arrays

func Sum(numbers []int) (sum int) {
	for _, v := range numbers {
		sum += v
	}
	return
}

func SumAllTails(arr ...[]int) (sums []int) {
	for _, a := range arr {
		if len(a) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(a[1:]))
		}
	}
	return
}
