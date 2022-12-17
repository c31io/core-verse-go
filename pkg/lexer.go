package cvg

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
		"tuple", "array", ",", "Length",
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
		// choices and sequences
		"|", "..", ";",
	}
	Literals = []string{
		// zero values
		"false?",
	}
)

type Token struct {
	name   TokenName
	lexeme string
}

func (inter *Interpreter) Lexer() {
	for {
		cmd := <-inter.cmdChan
		inter.sendToken(Token{0, cmd})
	}
}

func (inter *Interpreter) sendToken(token Token) {
	inter.tokenChan <- token
}
