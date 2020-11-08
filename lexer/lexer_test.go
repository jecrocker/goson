package lexer

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *Lexer
	}{
		{
			name: "Setup Correctly",
			args: args{
				input: "some_identifier",
			},
			want: &Lexer{
				buffer:       []rune("some_identifier"),
				ch:           's',
				position:     0,
				nextPosition: 1,
				currentChar:  1,
				currentLine:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
