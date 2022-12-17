package cvg

import "sync"

var AstMutex sync.Mutex

type AstNode struct {
	token    *Token
	children []Token
}

func (inter *Interpreter) Parser() {
	for {
		token := <-inter.tokenChan
		inter.print(string(token.lexeme))
	}
}