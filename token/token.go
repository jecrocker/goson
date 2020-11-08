package token

// Node - the token node
type Node struct {
	Type    Type
	Literal string
	Char    int
	Line    int
}

// Type - The type of the token
type Type string

const (
	LPAREN   = "LPAREN"
	RPAREN   = "RPAREN"
	LBRACKET = "LBRACKET"
	RBRACKET = "RBRACKET"
	COMMA    = "COMMA"
	COLON    = "COLON"

	TBOOL = "TBOOL"
	FBOOL = "FBOOL"
	NULL  = "NULL"

	STRING  = "STRING"
	NUMBER  = "NUMBER"
	INVALID = "INVALID"
	EOF     = "EOF"
)

var keywords = map[string]Type{
	"true":  TBOOL,
	"false": FBOOL,
	"null":  NULL,
}

// LookupType - lookup the type of a literal
func LookupType(word string) Type {
	if tok, ok := keywords[word]; ok {
		return tok
	}

	return INVALID
}
