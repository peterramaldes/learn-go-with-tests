package arrays

func Sum(numbers []int) (sum int) {
	for _, v := range numbers {
		sum += v
	}
	return
}

func SumAll(arr ...[]int) (sum []int) {
	for _, a := range arr {
		sum = append(sum, Sum(a))
	}
	return
}
