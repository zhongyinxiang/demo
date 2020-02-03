package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abcdEFGHijklMNOPQlasd"
	//	result := processLetter(str)
	//	fmt.Println(result)
	result := getfunc(str, processLetter)
	fmt.Println(result)
}

func processLetter(str string) string {
	result := ""
	for i, value := range str {
		if i%2 == 0 {
			result += strings.ToUpper(string(value))
		} else {
			result += strings.ToLower(string(value))
		}
	}
	return result
}
func getfunc(str string, myfunc func(string) string) string {
	return myfunc(str)
}
