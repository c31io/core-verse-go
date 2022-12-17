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
	tokenChan   chan Token
}

func (inter *Interpreter) Init() {
	inter.cmdChan = make(chan string)
	inter.tokenChan = make(chan Token)
	go inter.Lexer()
	go inter.Parser()
	inter.initialized = true
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

func (inter *Interpreter) showPrompt() {
	fmt.Print(*inter.prompt)
}

func (inter *Interpreter) print(a ...any) {
	newLine := func() string {
		if *inter.prompt == "" {
			return ""
		} else {
			return "\n"
		}
	}()
	fmt.Print(newLine, fmt.Sprint(a...), newLine, *inter.prompt)
}

func (inter *Interpreter) sendCmd(cmd string) {
	inter.cmdChan <- cmd
}
