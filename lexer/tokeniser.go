package lexer

import "github.com/jecrocker/goson/token"

// NextToken - Get the next token
func (l *Lexer) NextToken() token.Node {
	var node token.Node
	l.skipWhitespace()

	startChar := l.currentChar
	startLine := l.currentLine

	switch l.ch {
	case 0:
		node.Type = token.EOF
		node.Literal = ""
	case '{':
		node.Type = token.LPAREN
		node.Literal = "{"
	case '}':
		node.Type = token.RPAREN
		node.Literal = "}"
	case '[':
		node.Type = token.LBRACKET
		node.Literal = "["
	case ']':
		node.Type = token.RBRACKET
		node.Literal = "]"
	case ',':
		node.Type = token.COMMA
		node.Literal = ","
	case ':':
		node.Type = token.COLON
		node.Literal = ":"
	case '"':
		node.Type = token.STRING
		node.Literal = l.readString()
	case '-', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
		node.Type = token.NUMBER
		node.Literal = l.readNumber()
	default:
		tok := l.readIdentifier()
		node.Type = token.LookupType(tok)
		node.Literal = tok
	}

	l.readChar()

	node.Char = startChar
	node.Line = startLine

	return node
}
