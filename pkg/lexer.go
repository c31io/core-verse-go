package cvg

import (
	"fmt"
	"strings"
	"unicode"
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
)

type Token struct {
	name   TokenName
	lexeme []rune
}

type LexerState int

const (
	WhiteSpace LexerState = iota
	Word
	Symbol
	Value
)

func getTokenName(token []rune) TokenName {
	return Identifier
}

func (inter *Interpreter) LineLexer(line *string, ch chan Token) {
	s := []rune(*line)
	start, end := 0, 0
	state := WhiteSpace
	sendToken := func() {
		lexeme := string(s[start:end])
		token := Token{getTokenName([]rune(lexeme)), []rune(lexeme)}
		ch <- token
		start = end
	}
	unknowSymbol := func() {
		inter.print("Lexer: Unknown Symbol in string:\n" + *line + "\n" + strings.Repeat(" ", start) + "^")
		// TODO clear parser
	}
loop:
	for start != len(s) {
		switch state {
		case WhiteSpace:
			if unicode.IsSpace(s[end]) {
				start = end
			} else if unicode.IsLetter(s[end]) || s[end] == '_' {
				start = end
				state = Word
			} else if unicode.IsSymbol(s[end]) && s[end] != '_' {
				start = end
				state = Symbol
			} else if unicode.IsDigit(s[end]) {
				start = end
				state = Value
			} else {
				unknowSymbol()
				break loop
			}
		case Word:
			if unicode.IsSpace(s[end]) {
				sendToken()
				state = WhiteSpace
			} else if unicode.IsLetter(s[end]) || s[end] == '_' || unicode.IsDigit(s[end]) {
			} else if unicode.IsSymbol(s[end]) && s[end] != '_' {
				sendToken()
				state = Symbol
			} else if unicode.IsDigit(s[end]) {
				sendToken()
				state = Value
			} else {
				unknowSymbol()
				break loop
			}
		case Symbol:
			if unicode.IsSpace(s[end]) {
				sendToken()
				state = WhiteSpace
			} else if unicode.IsLetter(s[end]) || s[end] == '_' {
				sendToken()
				state = Word
			} else if unicode.IsSymbol(s[end]) && s[end] != '_' {
			} else if unicode.IsDigit(s[end]) {
				sendToken()
				state = Value
			} else {
				unknowSymbol()
				break loop
			}
		case Value:
			if unicode.IsSpace(s[end]) {
				sendToken()
				state = WhiteSpace
			} else if unicode.IsLetter(s[end]) || s[end] == '_' {
				sendToken()
				state = Word
			} else if unicode.IsSymbol(s[end]) && s[end] != '_' {
				sendToken()
				state = Symbol
			} else if unicode.IsDigit(s[end]) {
			} else {
				unknowSymbol()
				break loop
			}
		}
		end++
		if end == len(s) {
			sendToken()
			break
		}
	}
	close(ch)
}

func (inter *Interpreter) Lexer() {
	for {
		cmd := <-inter.cmdChan
		if strings.ToLower(strings.TrimSpace(cmd)) == "?" {
			inter.showState()
		}
		tokens := make(chan Token)
		go inter.LineLexer(&cmd, tokens)
		for token := range tokens {
			inter.sendToken(token)
		}
		// EOL was removed by line scanning.
		// Put it back as a visible lexeme.
		inter.sendToken(Token{Separator, []rune("EOL")})
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
