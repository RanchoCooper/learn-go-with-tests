package iteration

func Repeat(c string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result = result + c
	}
	return result
}
