package cvg

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Interpreter struct {
	initialized bool
	chanBufSize int
	prompt      *string
	cmdChan     chan string
	tokenChan   chan Token
	clearParser chan struct{}
	integers    map[string]int
	tuples      map[string][]int
	choices     map[string][]int
	astRoot     *AstNode
}

func (inter *Interpreter) Init(chanBufSize int) {
	inter.chanBufSize = chanBufSize
	inter.cmdChan = make(chan string, chanBufSize)
	inter.tokenChan = make(chan Token, chanBufSize)
	inter.clearParser = make(chan struct{}, chanBufSize)
	inter.integers = make(map[string]int)
	inter.tuples = make(map[string][]int)
	inter.choices = make(map[string][]int)
	inter.astRoot = new(AstNode)
	////////////// From the slides of Haskell Exchange 2022 talk //////////////
	// Verse is lenient but not strict:                                      //
	// - Like strict: everything gets evaluated in the end                   //
	// - Like lazy: functions can be called before the argument has a value  //
	///////////////// This is the reason why I have to go now /////////////////
	go inter.lexer()
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
	// TODO wait for active rewrites
	fmt.Println()
	return nil
}

func (inter *Interpreter) showPrompt() {
	fmt.Print(*inter.prompt)
}

func (inter *Interpreter) print(a ...any) {
	fmt.Print(inter.Sprint(a...))
}

func (inter *Interpreter) Sprint(a ...any) string {
	newLine := func() string {
		if *inter.prompt == "" {
			return ""
		} else {
			return "\n"
		}
	}()
	return fmt.Sprint(newLine, fmt.Sprint(a...), newLine, *inter.prompt)
}

func (inter *Interpreter) sendCmd(cmd string) {
	inter.cmdChan <- cmd
}
