package main

import (
	"fmt"
	"os"

	cvg "github.com/c31io/core-verse-go/pkg"
)

// If argument exists, run in interactive mode.
func main() {
	prompt := func() string {
		if len(os.Args) > 1 {
			return "> "
		} else {
			return ""
		}
	}()
	if len(prompt) != 0 {
		println("Core Verse interpreter in Go")
	}
	inter := new(cvg.Interpreter)
	inter.Init(64)
	err := inter.Run(&prompt, os.Stdin)
	if err != nil {
		fmt.Println(err.Error())
	}
}
