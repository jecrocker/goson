package lexer

import "unicode"

func isLetter(check rune) bool {
	return unicode.IsLetter(check)
}

func isDigit(check rune) bool {
	return unicode.IsDigit(check) || check == '.' || check == '-' || check == '+' || check == 'e' || check == 'E'
}
