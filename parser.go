package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(params string, parts int) []string {
	divider := " "
	return strings.SplitN(params, divider, parts)
}

func include(s []string, e string) bool {
	for _, a := range s {
		if strings.Compare(a, e) == 0 {
			return true
		}
	}
	return false
}

func repl() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">: ")
	text, _ := reader.ReadString('\n')
	return text
}

func str2Int(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
