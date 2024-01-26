package main

import (
	"fmt"
	"os"
	"tinylang/cmd/lex"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <cmd>\ncmds:\n", os.Args[0]);
		fmt.Println(lex.Usage())
		os.Exit(1)
	}

	switch os.Args[1] {
		
		case "lex":
			lex.Lex()
	}
}