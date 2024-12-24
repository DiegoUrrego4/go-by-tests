package loops

func Repeat(character rune, repeatCount int) string {
	if repeatCount == 0 {
		return ""
	}

	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += string(character)
	}

	return repeated
}
