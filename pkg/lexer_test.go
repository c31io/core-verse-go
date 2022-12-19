package cvg

import "testing"

type lex struct {
	in  string
	out []Token
}

var lexList []lex = []lex{
	{"3", []Token{{tokenLitNumber, "3"}, {tokenEOL, "EOL"}}},
	{"3+7", []Token{{tokenLitNumber, "3"}, {tokenPlus, "+"}, {tokenLitNumber, "7"}, {tokenEOL, "EOL"}}},
}

func TestLineLexer(t *testing.T) {
	inter := Interpreter{}
	inter.Init(1)
	for _, v := range lexList {
		c := make(chan Token, 1)
		go inter.LineLexer(&v.in, c)
		j := 0
		for out := range c {
			if out.name != v.out[j].name || out.lexeme != v.out[j].lexeme {
				t.Fatalf(`Want %v %v, got %v %v`, v.out[j].name, v.out[j].lexeme, out.name, out.lexeme)
			}
			j++
		}
	}
}
