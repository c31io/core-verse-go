package cvg

func (inter *Interpreter) Parser() {
	for {
		token := <-inter.tokenChan
		inter.print(token.lexeme)
	}
}
