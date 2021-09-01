package main

import (
	"fmt"
	"github.com/qwerty22121998/term"
)

func main() {
	term.Cls()
	term.Move(0, 0)
	fmt.Print(1)
	term.Down(1)
	fmt.Println(2)
	fmt.Print(3)
	term.Left(1)
	term.ClsToEOL()
	term.NextLine()
}
