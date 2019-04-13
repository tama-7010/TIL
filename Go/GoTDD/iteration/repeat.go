package iteration

const repeatCount = 5

func Repeat(character string) (string, int) {
	var repeated string
	var repeatedTime int
	for i := 0; i < repeatCount; i++ {
		repeated += character
		repeatedTime += 1
	}
	return repeated, repeatedTime
}
