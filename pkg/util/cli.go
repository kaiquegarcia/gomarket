package util

import (
	"bufio"
	"fmt"
	"os"
)

const (
	LINE_SEPARATOR = "----------------------------"
)

func PrintLineSeparator() {
	fmt.Println(LINE_SEPARATOR)
}

func AskCLI(text string) string {
	return askCLI(text+"\n", "> ")
}

func FinishByTooManyErrors() {
	fmt.Println("I'm sorry, I can't do this anymore")
	FinishCLI()
	panic("")
}

func FinishCLI() {
	askCLI("program finished, please press ENTER to stop it running ", "")
}

func askCLI(text string, scanPreffix string) string {
	fmt.Print(text)
	if scanPreffix != "" {
		fmt.Print(scanPreffix)
	}

	buf := bufio.NewScanner(os.Stdin)
	if buf.Scan() {
		return buf.Text()
	}

	if err := buf.Err(); err != nil {
		panic(err.Error())
	}

	return ""
}
