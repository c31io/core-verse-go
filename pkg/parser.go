package cvg

import (
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
			if len(tokens) > 0 && tokens[len(tokens)-1].name != tokenSequence {
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
func (expr *Expression) buildExpression(
	inter *Interpreter, tokens []Token) error {
	if !goodBrackets(tokens) {
		return ErrorBracketsNotMatch{}
	}
	expr.inter = inter
	switch {
	case isInteger(tokens):
		vi, err := strconv.ParseInt(tokens[0].lexeme, 10, 0)
		if err != nil {
			return ErrorParsingNumber{}
		}
		expr.exprType = exprValueInt
		expr.valueInt = &Value[int]{}
		expr.valueInt.NewValue([]int{int(vi)}, valueInteger)
	case isFloat(tokens):
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
	case isSequence(tokens):
		expr.exprType = exprSequence
		expr.innerExprs = []Expression{}
		bracketCount := 0
		head := 0
		for index, token := range tokens {
			switch token.name {
			case tokenParenL, tokenSqBraL, tokenCurlyL:
				bracketCount++
			case tokenParenR, tokenSqBraR, tokenCurlyR:
				bracketCount--
			case tokenSequence:
				if bracketCount == 0 {
					innerExpr := Expression{outerExpr: expr}
					expr.innerExprs = append(expr.innerExprs, innerExpr)
					expr.innerExprs[len(expr.innerExprs)-1].buildExpression(inter, tokens[head:index])
					head = index + 1
				}
			default:
				if index == len(tokens)-1 {
					index += 1
					innerExpr := Expression{outerExpr: expr}
					expr.innerExprs = append(expr.innerExprs, innerExpr)
					expr.innerExprs[len(expr.innerExprs)-1].buildExpression(inter, tokens[head:index])
				}
			}
		}
		// scope?
		// fail?
		// for-do?
		// if-else?
		// choices?
		// application?
		// unify?
	default:
		return ErrorUnknownExpression{}
	}
	return nil
}

func goodBrackets(tokens []Token) bool {
	bracketStack := make([]tokenName, len(tokens))
	height := 0
	for _, token := range tokens {
		switch token.name {
		case tokenParenL:
			bracketStack[height] = tokenParenL
			height++
		case tokenParenR:
			if height < 1 || bracketStack[height-1] != tokenParenL {
				return false
			}
			height--
		case tokenSqBraL:
			bracketStack[height] = tokenSqBraL
			height++
		case tokenSqBraR:
			if height < 1 || bracketStack[height-1] != tokenSqBraL {
				return false
			}
			height--
		case tokenCurlyL:
			bracketStack[height] = tokenCurlyL
			height++
		case tokenCurlyR:
			if height < 1 || bracketStack[height-1] != tokenCurlyL {
				return false
			}
			height--
		}
	}
	return height == 0
}

func isInteger(tokens []Token) bool {
	return len(tokens) == 1 && tokens[0].name == tokenDecimal
}

func isFloat(tokens []Token) bool {
	return len(tokens) == 3 && tokens[0].name == tokenDecimal &&
		tokens[1].name == tokenPoint && tokens[2].name == tokenDecimal
}

func isSequence(tokens []Token) bool {
	bracketCount := 0
	for _, token := range tokens {
		switch token.name {
		case tokenParenL, tokenSqBraL, tokenCurlyL:
			bracketCount++
		case tokenParenR, tokenSqBraR, tokenCurlyR:
			bracketCount--
		case tokenSequence:
			if bracketCount == 0 {
				return true
			}
		}
	}
	return false
}
