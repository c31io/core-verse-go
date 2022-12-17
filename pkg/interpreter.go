package cvg

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Interpreter struct {
	initialized bool
	prompt      *string
	cmdChan     chan string
}

func (inter *Interpreter) Init() {
	inter.cmdChan = make(chan string)
	go inter.Lexer()
	inter.initialized = true
}

func (inter *Interpreter) showPrompt() {
	fmt.Print(*inter.prompt)
}

func (inter *Interpreter) println(a ...any) {
	fmt.Print("\n", fmt.Sprint(a...), "\n", *inter.prompt)
}

func (inter *Interpreter) Run(prompt *string, input *os.File) error {
	if !inter.initialized {
		return errors.New("Interpreter uninitialized")
	}
	inter.prompt = prompt
	scanner := bufio.NewScanner(os.Stdin)
	inter.showPrompt()
	for scanner.Scan() {
		inter.showPrompt()
		inter.sendCmd(scanner.Text())
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	fmt.Println()
	return nil
}

func (inter *Interpreter) sendCmd(cmd string) {
	inter.cmdChan <- cmd
}

func (inter *Interpreter) Lexer() {
	for {
		cmd := <-inter.cmdChan
		inter.println(cmd)
	}
}
