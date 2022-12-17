package main

import (
	"fmt"
	"os"

	cvg "github.com/c31io/core-verse-go/pkg"
)

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
	inter.Init()
	err := inter.Run(&prompt, os.Stdin)
	if err != nil {
		fmt.Println(err.Error())
	}
}
