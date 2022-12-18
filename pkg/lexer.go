package cvg

import (
	"fmt"
	"strings"
	"unicode"
)

type lexerState int

const (
	whiteSpace lexerState = iota
	word
	symbol
)

func (inter *Interpreter) LineLexer(line *string, ch chan Token) {
	s := []rune(*line)
	start, end := 0, 0
	state := whiteSpace
	sendToken := func(nextState lexerState) {
		if state != nextState {
			if state != whiteSpace {
				lexeme := string(s[start:end])
				token := Token{getTokenName([]rune(lexeme)), []rune(lexeme)}
				ch <- token
			}
			start = end
			state = nextState
		}
	}
	unknowSymbol := func() {
		inter.print("Lexer: Unknown character in string:\n" + *line + "\n" + strings.Repeat(" ", end) + "^")
		inter.clearParser <- struct{}{}
	}
	breakLoop := false
	handleLastRune := func() {
		lastRune := s[end]
		switch {
		case isSpace(lastRune):
			sendToken(whiteSpace)
		case isSymbol(lastRune):
			sendToken(symbol)
		case isAlphaNumeric(lastRune):
			sendToken(word)
		default:
			unknowSymbol()
			breakLoop = true
		}
	}
	for start != len(s) {
		handleLastRune()
		end++
		if end == len(s) {
			sendToken(whiteSpace)
			break
		}
		if breakLoop {
			break
		}
	}
	close(ch)
}

func (inter *Interpreter) lexer() {
	for {
		cmd := <-inter.cmdChan
		if strings.ToLower(strings.TrimSpace(cmd)) == "?" {
			inter.showState()
		}
		tokens := make(chan Token, inter.chanBufSize)
		go inter.LineLexer(&cmd, tokens)
		for token := range tokens {
			inter.sendToken(token)
		}
		// EOL was removed by line scanning.
		// Put it back as a visible lexeme.
		inter.sendToken(Token{0, []rune("EOL")})
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

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r' || r == '\n'
}

func isSymbol(r rune) bool {
	symbols := "(){}[]+-*/<>=:.|;,"
	for _, s := range symbols {
		if s == r {
			return true
		}
	}
	return false
}

func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}
