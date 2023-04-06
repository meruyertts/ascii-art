package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	input := os.Args[1:]
	if len(input) == 0 {
		fmt.Println("no input was entered")
		return
	}
	myStr := input[0]
	if !isASCII(myStr) {
		fmt.Println("non-ASCII character was entered")
		return
	}

	if len(myStr) == 0 {
		return
	}

	if lineCounter() < 855 {
		fmt.Println("file does not contain all characters")
		return
	}
	re := regexp.MustCompile(`\\n`)
	newStr := re.Split(myStr, -1)
	for i := 0; i < len(newStr); i++ {
		if len(newStr[i]) > 0 {
			printWord(newStr[i])
		}
	}

	for i := len(myStr) - 2; i >= 0; i -= 2 {
		if myStr[i:i+2] != "\\n" {
			break
		}
		fmt.Println()

	}
}

func lineCounter() int {
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("no file was found")
		os.Exit(0)
	}
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	return lineCount
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func printWord(s string) {
	firstline := ""
	myline := ""
	isitmyfile := true
	for line := 2; line <= 9; line++ {
		for _, char := range s {
			myrune := int(char)
			for i := 32; i <= 126; i++ {
				j := (i - 32) * 9
				if myrune == i {
					firstline, isitmyfile = readExactLine("standard.txt", line+j)
					if !isitmyfile {
						return
					}
					myline += firstline

				}
			}
		}
		fmt.Println(strings.ReplaceAll(myline, "\n", ""))
		myline = ""
	}
}

func readExactLine(fileName string, lineNumber int) (string, bool) {
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	br := bufio.NewReader(inputFile)
	for i := 1; i < lineNumber; i++ {
		_, _ = br.ReadString('\n')
	}
	str, err := br.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}

	return str, true
}
