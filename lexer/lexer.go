package lexer

// Lexer - the lexer object
type Lexer struct {
	buffer       []rune
	ch           rune
	position     int
	nextPosition int
	currentChar  int
	currentLine  int
}

// New - create a new lexer
func New(obj string) *Lexer {
	l := &Lexer{
		buffer: []rune(obj),
	}

	l.readChar()

	return l
}
