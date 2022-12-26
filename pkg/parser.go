package cvg

import (
	"fmt"
	"strconv"
)

// Break at the EOL without a semicolon before it.
func (inter *Interpreter) parser() {
	var root *Expression
	var tokens []Token
	for {
		if root == nil {
			root = &Expression{}
			tokens = []Token{}
		}
		newToken := <-inter.tokenChan
		if newToken.name == tokenEOL {
			if tokens[len(tokens)-1].name != tokenSequence {
				err := root.buildExpression(inter, tokens)
				if err != nil {
					inter.print(err.Error())
				} else {
					root.Floater()
					root.Rewriter()
				}
				root, tokens = nil, nil
			}
		} else {
			tokens = append(tokens, newToken)
		}
	}
}

// Parse tokens into an expression.
func (expr *Expression) buildExpression(inter *Interpreter, tokens []Token) error {
	fmt.Printf("%v\n", tokens)
	expr.inter = inter
	// Is it a value?
	if isInteger(tokens) {
		vi, err := strconv.ParseInt(tokens[0].lexeme, 10, 0)
		if err != nil {
			return ErrorParsingNumber{}
		}
		expr.exprType = exprValueInt
		expr.valueInt = &Value[int]{}
		expr.valueInt.NewValue([]int{int(vi)}, valueInteger)
	}
	if isFloat(tokens) {
		vf, err := strconv.ParseFloat(
			tokens[0].lexeme+
				tokens[1].lexeme+
				tokens[2].lexeme,
			64)
		if err != nil {
			return ErrorParsingNumber{}
		}
		expr.exprType = exprValueFloat
		expr.valueFloat = &Value[float64]{}
		expr.valueFloat.NewValue([]float64{vf}, valueFloat)
	}
	// Is it a sequence?
	// Is it a scope?
	// Is it a fail?
	// Is it a for-do?
	// Is it a if-else?
	// Is it choices?
	// Is it an application?
	// Is it a unify?
	return nil
}

func isInteger(tokens []Token) bool {
	return len(tokens) == 1 && tokens[0].name == tokenDecimal
}

func isFloat(tokens []Token) bool {
	return len(tokens) == 3 && tokens[0].name == tokenDecimal &&
		tokens[1].name == tokenPoint && tokens[2].name == tokenDecimal
}

func isSequence(tokens []Token) bool {
	return true
}
