package lexer

import (
	"reflect"
	"testing"

	"github.com/jecrocker/goson/token"
)

func TestLexer_NextToken(t *testing.T) {
	lexer := New(`{}[],:"abc\"" true false null -1.234567e+1`)

	tests := []struct {
		name string
		l    *Lexer
		want token.Node
	}{
		{
			name: "LPAREN",
			l:    lexer,
			want: token.Node{
				Type:    token.LPAREN,
				Char:    1,
				Line:    0,
				Literal: "{",
			},
		},
		{
			name: "RPAREN",
			l:    lexer,
			want: token.Node{
				Type:    token.RPAREN,
				Char:    2,
				Line:    0,
				Literal: "}",
			},
		},
		{
			name: "LBRACKET",
			l:    lexer,
			want: token.Node{
				Type:    token.LBRACKET,
				Char:    3,
				Line:    0,
				Literal: "[",
			},
		},
		{
			name: "RBRACKET",
			l:    lexer,
			want: token.Node{
				Type:    token.RBRACKET,
				Char:    4,
				Line:    0,
				Literal: "]",
			},
		},
		{
			name: "COMMA",
			l:    lexer,
			want: token.Node{
				Type:    token.COMMA,
				Char:    5,
				Line:    0,
				Literal: ",",
			},
		},
		{
			name: "COLON",
			l:    lexer,
			want: token.Node{
				Type:    token.COLON,
				Char:    6,
				Line:    0,
				Literal: ":",
			},
		},
		{
			name: "STRING",
			l:    lexer,
			want: token.Node{
				Type:    token.STRING,
				Char:    7,
				Line:    0,
				Literal: "\"abc\\\"\"",
			},
		},
		{
			name: "TBOOL",
			l:    lexer,
			want: token.Node{
				Type:    token.TBOOL,
				Char:    15,
				Line:    0,
				Literal: "true",
			},
		},
		{
			name: "FBOOL",
			l:    lexer,
			want: token.Node{
				Type:    token.FBOOL,
				Char:    20,
				Line:    0,
				Literal: "false",
			},
		},
		{
			name: "NULL",
			l:    lexer,
			want: token.Node{
				Type:    token.NULL,
				Char:    26,
				Line:    0,
				Literal: "null",
			},
		},
		{
			name: "NUMBER",
			l:    lexer,
			want: token.Node{
				Type:    token.NUMBER,
				Char:    31,
				Line:    0,
				Literal: "-1.234567e+1",
			},
		},
		{
			name: "EOF",
			l:    lexer,
			want: token.Node{
				Type:    token.EOF,
				Char:    44,
				Line:    0,
				Literal: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.NextToken(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lexer.NextToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
