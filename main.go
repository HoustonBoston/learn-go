package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	// m "math"
	// "net/http"
	// _ "net/http/pprof"
	"os"
	// "strconv"
)

func main() {
	for {
		PrintPrompt()
		ScanText()
	}
}

func PrintPrompt() {
	wd, _ := os.Getwd()

	var prompt string
	prompt = wd + "$ "

	fmt.Printf("%s", prompt) // part one done
}

func ScanText() {
	scanner := bufio.NewScanner(os.Stdin)
	hasInput := scanner.Scan()  // return false if EOF is detected 
	in := scanner.Text()
	err := scanner.Err()

	if err == io.EOF || hasInput == false {
		fmt.Print("\n")
		os.Exit(0)
	}

	// part 3
	if hasInput == true && in != "" {
		// split into tokens
		input_tokens := strings.Split(in, " ")
		ValidateTokens(input_tokens)
	} 
}

func ValidateTokens(token_arr []string) {
	token := token_arr[0]
	argCount := len(token_arr) - 1 // ex: cd /bin/ <- counts as one argument

	switch token {
		case "exit": 
			os.Exit(0)
		case "cd":
			Change_dir(token_arr[1], argCount)
		default:
			fmt.Println("Unknown command: ", token)
	}
}

func Change_dir(dir string, argCount int) {
	err := os.Chdir(dir)

	if argCount != 1 {  // takes exactly one
		fmt.Println("cd [dir-to-go]")
		return
	}

	if err != nil {
		fmt.Println(err)
	}
}
