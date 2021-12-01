package iteration

const repeatCount = 4

func Repeat(s string) string {
	var result string
	result = ""
	for i := 0; i < repeatCount; i++ {
		result += s
	}
	return result
}
