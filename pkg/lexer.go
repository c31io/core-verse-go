package cvg

import (
	"fmt"
	"strings"
)

type TokenName int

const (
	Identifier TokenName = iota
	Keyword
	Separator
	Operator
	Literal
	Comment
)

var (
	Keywords = []string{
		// conditions
		"if", "then", "else",
		// loops
		"for", "do",
		// types
		"int",
		// tuples
		"tuple", "array", "Length",
	}
	Separators = []string{
		// brackets
		"(", ")", "{", "}", "[", "]",
		// EOL, time to execute?
		"EOL",
	}
	Operators = []string{
		// arithmetics
		"+", "-", "*", "/",
		// comparisons
		"<", ">", "<=", ">=",
		// lambdas, binds and unifications
		"=>", ":", "=",
		// choices, sequences and tuples
		"..", "|", ";", ",",
	}
	Literals = []string{
		// zero values
		"false?",
	}
	WhiteSpaces = []string{
		" ",
		"\t",
		"\r",
		"\n",
	}
)

type Token struct {
	name   TokenName
	lexeme []byte
}

func (inter *Interpreter) LineLexer(line *string) {
	s := []byte(*line)
	start, end := 0, 0
	for start != len(s) {
		end++
		break // TODO lexing
	}
}

func (inter *Interpreter) Lexer() {
	for {
		cmd := <-inter.cmdChan
		if strings.ToLower(strings.TrimSpace(cmd)) == "?" {
			inter.showState()
		}
		inter.LineLexer(&cmd)
		// EOL was removed by line scanning.
		// Put it back as a visible lexeme.
		inter.sendToken(Token{Separator, []byte("EOL")})
	}
}

func (inter *Interpreter) sendToken(token Token) {
	inter.tokenChan <- token
}

func (inter *Interpreter) showState() {
	fmt.Print(
		inter.Sprint(
			strings.ReplaceAll(
				strings.TrimSpace(
					fmt.Sprintf("%+v\n", inter)),
				" ",
				"\n")))
}
