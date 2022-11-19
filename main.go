package main

import (
	"fmt"
	"os"
	"strings"
)

func Wc(arg string) int {
	if arg == "" {
		return 0
	} else {
		s := strings.Split(arg, " ")
		cnt := 0
		for i := 0; i < len(s); i++ {
			if len(s[i]) > 0 {
				cnt++
			}
		}
		return cnt
	}
}

func main() {
	toGetAllArgs := os.Args[1:]
	argsCount := len(toGetAllArgs)

	if argsCount > 0 {
		wordCount := 0
		for i := 1; i < argsCount+1; i++ {
			wordCount += Wc(os.Args[i])
		}
		fmt.Println(wordCount)
		return
	}
	fmt.Println(argsCount)
	return
}
