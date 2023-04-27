package main

import (
	"dl/lexer"
	"dl/parser"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("prog.dl")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	l := lexer.New(file)

	p := parser.New(l)

	p.Parse()

	fmt.Println("finalizado")
}
