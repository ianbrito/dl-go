package main

import (
	"dl/lexer"
	"dl/lexer/tag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("prog.dl")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	l := lexer.New(file)
	token := l.NextToken()

	for token.Tag() != tag.EOF {
		fmt.Println(token.String())
		token = l.NextToken()
	}
}
