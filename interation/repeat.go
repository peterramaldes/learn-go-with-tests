package interation

func Repeat(chr string, qtd int) string {
	var result string
	for i := 0; i < qtd; i++ {
		result += chr
	}
	return result
}
