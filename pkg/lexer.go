package cvg

import (
	"fmt"
	"strings"
	"unicode"
)

// Greedy states of the lexing loop in LineLexer().
type lexerState int

const (
	// whitespaces are described by isSpace().
	whiteSpace lexerState = iota
	// Word is a string of unicode letters, '_', and '?'.
	word
	// Tokens that are made of symbols are listed in the file token.go.
	symbol
)

// The LineLexer breaks the line into tokens.
func (inter *Interpreter) LineLexer(line *string, ch chan Token) {
	// The states:
	s := []rune(*line)
	start, end := 0, 0
	state := whiteSpace
	// The closures:
	unknowChar := func() {
		inter.print("Unknown character:\n" +
			*line + "\n" +
			strings.Repeat(" ", end) + "^")
		inter.clearParser <- struct{}{}
	}
	unknowToken := func(name *string) {
		inter.print("Unknown token: " + *name)
		inter.clearParser <- struct{}{}
	}
	sendToken := func() {
		if state != whiteSpace {
			lexeme := string(s[start:end])
			name := getTokenName(&lexeme)
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
			// Break the lexeme if
			if ss := string(s[start : end+1]); !(ss == "<=" ||
				ss == ">=" ||
				ss == "=>" ||
				ss == "..") {
				sendToken()
			}
		}
	}
	// The state machine:
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
	// The main lexing loop:
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
	// Check if the main loop stopped without errors.
	if !breakLoop {
		// EOL was removed by line scanning.
		// Put it back as a visible lexeme.
		ch <- Token{tokenEOL, "EOL"}
	}
	close(ch)
}

// Wrapper function to make LineLexer() unit testable.
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

// Send token to its channel.
func (inter *Interpreter) sendToken(token Token) {
	inter.tokenChan <- token
}

// Pretty print Interpreter{}.
func (inter *Interpreter) showState() {
	fmt.Print(
		inter.Sprint(
			strings.ReplaceAll(
				strings.TrimSpace(
					fmt.Sprintf("%+v\n", inter)),
				" ",
				"\n")))
}

// If a rune is one of the whitespaces (' ', '\t', '\t', '\n').
func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r' || r == '\n'
}

// If a rune is a symbol used by a separator or an operator.
func isSymbol(r rune) bool {
	symbols := "(){}[]+-*/<>=:.|;,"
	for _, s := range symbols {
		if s == r {
			return true
		}
	}
	return false
}

// If a rune is a letter, '_', '?', or a digit.
func isAlphaNumeric(r rune) bool {
	return isAlpha(r) || unicode.IsDigit(r)
}

// If a rune is a letter, '_', or '?'.
// The '?' is preserved for zero-value literals.
func isAlpha(r rune) bool {
	return r == '_' || r == '?' || unicode.IsLetter(r)
}
