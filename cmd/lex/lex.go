package lex

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"tinylang/pkg/lexer"
	"tinylang/pkg/token"
)

var LexCmd *flag.FlagSet

func Lex() {
	LexCmd = flag.NewFlagSet("lex", flag.ExitOnError)
	filename := LexCmd.String("i", "", "Input Source (Defaults To Stdin)")
	output := LexCmd.String("o", "", "Output Path (Defaults To False)")
	pretty := LexCmd.Bool("pretty", false, "Set to use Pretty Printing (Defaults To False)")
	LexCmd.Parse(os.Args[2:])

	var sourceBytes []byte
	var err error

	if len(*filename) > 0 {
		sourceBytes, err = os.ReadFile(*filename)
	} else {
		sourceBytes, err = io.ReadAll(os.Stdin)
	}

	if (err != nil) {
		log.Fatal(err)
	}
	
	l := lexer.New(string(sourceBytes), filename)

	tokens := []token.Token{}

	for tok := l.NextToken(); tok.Type != token.EOI; tok = l.NextToken() {
		tokens = append(tokens, tok)
	}

	tokens = append(tokens, l.NextToken())

	var buffer []byte

	if *pretty {
		buffer, err = json.MarshalIndent(tokens, "", "\t")
	} else {
		buffer, err = json.Marshal(tokens)
	}

	

	if (err != nil) {
		log.Fatalln(err)
	}

	if len(*output) == 0 {
		fmt.Println(string(buffer))
	} else {
		os.WriteFile(*output, buffer, os.ModePerm)
	}
}

func Usage() string {
	var s bytes.Buffer

	s.WriteString("- lex\n")
	s.WriteString("  -i <Input file>\n")

	return s.String()
}