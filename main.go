package main

import (
	"dl/lexer/tag"
	"dl/lexer/token"
	"fmt"
)

func main() {
	t := token.New(tag.ID, "max")

	fmt.Println(t.String())
}
