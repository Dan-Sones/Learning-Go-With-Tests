package iteration

func Repeat(character string, charCount int) string {
	var repeated string
	for i := 0; i < charCount; i++ {
		repeated += character
	}
	return repeated
}
