package cvg

import (
	"strconv"
	"unicode"
)

// Token{} holds a token's type and its piece of string.
type Token struct {
	name   tokenName
	lexeme string
}

// Keywords, separators, operators, literals, and variables.
type tokenName int

const (
	//// keywords

	// conditions

	tokenIf   tokenName = iota // if
	tokenThen                  // then
	tokenElse                  // else

	// loops

	tokenFor // for
	tokenDo  // do

	// types

	tokenInt // int

	// tuples

	tokenTuple  // tuple
	tokenArray  // array
	tokenLength // Length

	//// separators

	tokenParenL // (
	tokenParenR // )
	tokenSqBraL // [
	tokenSqBraR // ]
	tokenCurlyL // {
	tokenCurlyR // }
	tokenEOL    // \n

	//// operators

	// arithmetics

	tokenPlus     // +
	tokenMinus    // -
	tokenMultiply // *
	tokenDivide   // /

	// comparisons

	tokenLe  // <
	tokenGr  // >
	tokenLeq // <=
	tokenGeq // >=

	// lambdas, binds and unifications

	tokenLambda     // =>
	tokenScope      // :
	tokenUnify      // =
	tokenScopeUnify // :=

	// choices, sequences and tuples

	tokenRange    // ..
	tokenChoise   // |
	tokenSequence // ;
	tokenComma    // ,

	//// literals

	// numbers

	tokenLitNumber // 42

	// zero values

	tokenFail // false?

	//// variables

	tokenVar
)

// What kind of token is this string?
func getTokenName(s *string) tokenName {
	switch *s {
	case "if":
		return tokenIf
	case "then":
		return tokenThen
	case "else":
		return tokenElse
	case "for":
		return tokenFor
	case "do":
		return tokenDo
	case "int":
		return tokenInt
	case "tuple":
		return tokenTuple
	case "array":
		return tokenArray
	case "Length":
		return tokenLength
	case "(":
		return tokenParenL
	case ")":
		return tokenParenR
	case "[":
		return tokenSqBraL
	case "]":
		return tokenSqBraR
	case "{":
		return tokenCurlyL
	case "}":
		return tokenCurlyR
	case "\n":
		return tokenEOL
	case "+":
		return tokenPlus
	case "-":
		return tokenMinus
	case "*":
		return tokenMultiply
	case "/":
		return tokenDivide
	case "<":
		return tokenLe
	case ">":
		return tokenGr
	case "<=":
		return tokenLeq
	case ">=":
		return tokenGeq
	case "=>":
		return tokenLambda
	case ":":
		return tokenScope
	case "=":
		return tokenUnify
	case ":=":
		return tokenScopeUnify
	case "..":
		return tokenRange
	case "|":
		return tokenChoise
	case ";":
		return tokenSequence
	case ",":
		return tokenComma
	case "false?":
		return tokenFail
	default:
		r := []rune(*s)[0]
		if unicode.IsDigit(r) {
			_, err := strconv.ParseFloat(*s, 64)
			if err != nil {
				return -1
			}
			return tokenLitNumber
		} else if isAlpha(r) {
			return tokenVar
		} else {
			return -1
		}
	}
}
