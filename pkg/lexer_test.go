package cvg

import "testing"

// Holds testing pairs.
type lex struct {
	in  string
	out []Token
}

// Examples from the slides. URL: https://simon.peytonjones.org/assets/pdfs/haskell-exchange-22.pdf
var lexList []lex = []lex{
	{"3", []Token{{tokenLitNumber, "3"}, {tokenEOL, "EOL"}}},
	{"3+7", []Token{{tokenLitNumber, "3"}, {tokenPlus, "+"}, {tokenLitNumber, "7"}, {tokenEOL, "EOL"}}},
}

// Test code from the slides.
func TestLineLexer(t *testing.T) {
	inter := Interpreter{}
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
