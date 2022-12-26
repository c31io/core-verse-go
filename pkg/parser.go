package cvg

import "fmt"

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
				err := root.buildExpression(tokens)
				if err != nil {
					inter.print(err.Error())
				} else {
					root.Floater()
					//root.Rewriter()
				}
				root, tokens = nil, nil
			}
		} else {
			tokens = append(tokens, newToken)
		}
	}
}

// Parse tokens into an expression.
func (node *Expression) buildExpression(tokens []Token) error {
	fmt.Printf("%v\n", tokens)
	return nil
}
