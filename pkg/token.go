package cvg

type Token struct {
	name   tokenName
	lexeme []rune
}

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
	tokenLambda // =>
	tokenBind   // :
	tokenUnify  // =
	// choices, sequences and tuples
	tokenRange    // ..
	tokenChoise   // |
	tokenSequence // ;
	tokenComma    // ,
	//// literals
	// zero values
	tokenFail // false?
)

func getTokenName(token []rune) tokenName {
	return 0
}
