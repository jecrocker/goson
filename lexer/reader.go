package lexer

import (
	"fmt"
)

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.buffer) {
		l.ch = 0
	} else {
		l.ch = l.buffer[l.nextPosition]
	}

	l.position = l.nextPosition
	l.nextPosition++
	l.currentChar++
}

func (l *Lexer) readString() string {
	str := "\""
	l.readChar()
	for l.ch != '"' && l.ch != 0 {
		str = fmt.Sprintf("%s%c", str, l.ch)
		if l.ch == '\\' {
			l.readChar()
			str = fmt.Sprintf("%s%c", str, l.ch)
		}
		l.readChar()
	}
	l.readChar()
	str = fmt.Sprintf("%s\"", str)
	return str
}

func (l *Lexer) readIdentifier() string {
	ident := ""
	for isLetter(l.ch) {
		ident = fmt.Sprintf("%s%c", ident, l.ch)
		l.readChar()
	}

	return ident
}

func (l *Lexer) readNumber() string {
	number := ""

	for isDigit(l.ch) {
		number = fmt.Sprintf("%s%c", number, l.ch)
		l.readChar()
	}

	return number
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' || l.ch == '\r' {
			l.currentLine++
			l.currentChar = 0
		}
		l.readChar()
	}
}

func (l *Lexer) peekChar() rune {
	if l.nextPosition >= len(l.buffer) {
		return 0
	}
	return l.buffer[l.nextPosition]
}
