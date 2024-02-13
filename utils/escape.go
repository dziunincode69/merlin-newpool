package utils

func EscapeMarkdownV2(text string) string {
	specialChars := []rune{'_', '*', '[', ']', '(', ')', '~', '`', '>', '#', '+', '-', '=', '|', '{', '}', '.', '!'}
	var escapedText string
	for _, char := range text {
		for _, specialChar := range specialChars {
			if char == specialChar {
				escapedText += "\\" // Add a backslash before the special character
				break
			}
		}
		escapedText += string(char)
	}
	return escapedText
}
