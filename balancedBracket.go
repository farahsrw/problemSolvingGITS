package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reader bufio.Reader
	var inputLine string
	var input string
	var err error

	reader = *bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan inputan: ")
	inputLine, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi kesalahan saat membaca input:", err)
		return
	}

	input = strings.TrimSpace(inputLine)
	fmt.Printf("Apakah balanced bracket? %s", balancedBracket(input))
}

func balancedBracket(input string) string {
	var stack []rune
	var karakter, top rune
	var i int
	input = strings.ReplaceAll(input, " ", "")

	for i = 0; i < len(input); i++ {
		karakter = rune(input[i])

		if karakter == '(' || karakter == '[' || karakter == '{' {
			stack = append(stack, karakter)
		}

		if karakter == ')' || karakter == ']' || karakter == '}' {
			if len(stack) == 0 {
				return "NO"
			}

			top = stack[len(stack)-1]

			if (karakter == ')' && top != '(') ||
				(karakter == ']' && top != '[') ||
				(karakter == '}' && top != '{') {
				return "NO"
			}

			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return "YES"
	} else {
		return "NO"
	}
}
