package cvg

import "sync"

var AstMutex sync.Mutex

type AstNode struct {
	token    *Token
	children []Token
	rewriter *Rewriter
}

func (inter *Interpreter) parser() {
	for {
		token := <-inter.tokenChan
		inter.print(token.name, "\t", string(token.lexeme))
	}
}
