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
	unknowChar := func() {
		inter.print("Unknown character:\n" + *line + "\n" + strings.Repeat(" ", end) + "^")
		inter.clearParser <- struct{}{}
	}
	unknowToken := func(name *string) {
		inter.print("Unknown token: " + *name)
		inter.clearParser <- struct{}{}
	}
	sendToken := func() {
		if state != whiteSpace {
			lexeme := string(s[start:end])
			name := getTokenName(lexeme)
			if name != -1 {
				token := Token{name, lexeme}
				ch <- token
			} else {
				unknowToken(&lexeme)
			}
		}
		start = end
	}
	sendTokenMaybe := func(nextState lexerState) {
		if state != nextState {
			sendToken()
			state = nextState
		} else if nextState == symbol {
			if ss := string(s[start : end+1]); !(ss == "<=" ||
				ss == ">=" ||
				ss == "=>" ||
				ss == "..") {
				sendToken()
			}
		}
	}
	breakLoop := false
	handleLastRune := func() {
		lastRune := s[end]
		switch {
		case isSpace(lastRune):
			sendTokenMaybe(whiteSpace)
		case isSymbol(lastRune):
			sendTokenMaybe(symbol)
		case isAlphaNumeric(lastRune):
			sendTokenMaybe(word)
		default:
			unknowChar()
			breakLoop = true
		}
	}
	for start != len(s) {
		handleLastRune()
		end++
		if breakLoop {
			break
		}
		if end == len(s) {
			sendTokenMaybe(whiteSpace)
			break
		}
	}
	if !breakLoop {
		// EOL was removed by line scanning.
		// Put it back as a visible lexeme.
		ch <- Token{tokenEOL, "EOL"}
	}
	close(ch)
}

func (inter *Interpreter) lexer() {
	for {
		cmd := <-inter.cmdChan
		if strings.ToLower(strings.TrimSpace(cmd)) == "??" {
			inter.showState()
		}
		tokens := make(chan Token, inter.chanBufSize)
		go inter.LineLexer(&cmd, tokens)
		for token := range tokens {
			inter.sendToken(token)
		}
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
	return isAlpha(r) || unicode.IsDigit(r)
}

func isAlpha(r rune) bool {
	return r == '_' || r == '?' || unicode.IsLetter(r)
}
