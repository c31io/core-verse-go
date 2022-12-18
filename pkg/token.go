package cvg

type Token struct {
	name   TokenName
	lexeme []rune
}

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

func getTokenName(token []rune) TokenName {
	return Identifier
}
