package main

import (
	"bufio"
	"fmt"
	"go-lox/pkg/parser"
	"go-lox/pkg/tokens"
	"go-lox/pkg/utils"
	"io/ioutil"
	"os"
)

// run runs a line of lox
func run(s string) {
	t := tokens.NewTokenizer(s)
	t.ScanTokens()
	tokenList := t.GetTokenList()
	//fmt.Println(tokenList)
	Parser := parser.NewParser(tokenList)
	expr := Parser.Parse()
	Interpreter := parser.NewInterpreter()
	Interpreter.Interpret(expr)
}

func runFile(path string) {
	content, err := ioutil.ReadFile(path) // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	run(string(content))
}

func runPrompt() {

	utils.PrintPromptInit()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">>> ")
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			continue
		} else {
			run(line)
		}
	}
}

func main() {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage: go-lox [script]")
		return
	} else if len(args) == 2 {
		runFile(args[1])
	} else {
		runPrompt()
	}
}
