package cvg

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// Interpreter{} holds states and communicating channels.
type Interpreter struct {
	initialized bool
	chanBufSize int
	prompt      *string
	cmdChan     chan string
	tokenChan   chan Token
	clearParser chan struct{}
}

// Late initialization. Some functions do not require an initialized
// interpreter, e.g. LineLexer(). chanBufSize should always be positive.
func (inter *Interpreter) Init(chanBufSize int) {
	inter.chanBufSize = chanBufSize
	inter.cmdChan = make(chan string, chanBufSize)
	inter.tokenChan = make(chan Token, chanBufSize)
	inter.clearParser = make(chan struct{}, chanBufSize)
	////////////// From the slides of Haskell Exchange 2022 talk //////////////
	// Verse is lenient but not strict:                                      //
	// - Like strict: everything gets evaluated in the end                   //
	// - Like lazy: functions can be called before the argument has a value  //
	///////////////// This is the reason why I have to go now /////////////////
	go inter.lexer()
	go inter.parser()
	inter.initialized = true
}

// Run interpreter with nonempty prompt and os.Stdin for interactive mode, or
// run with an empty string "" as prompt to handle file input for batch mode.
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

// Send prompt to stdout.
func (inter *Interpreter) showPrompt() {
	fmt.Print(*inter.prompt)
}

// Print to stdout without messing user input.
func (inter *Interpreter) print(a ...any) {
	fmt.Print(inter.Sprint(a...))
}

// Prepare message to stdout without messing user input.
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

// Send command to lexer().
func (inter *Interpreter) sendCmd(cmd string) {
	inter.cmdChan <- cmd
}
