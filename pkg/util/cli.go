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

func Ask(text string) string {
	fmt.Println(text)
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()

	if err := buf.Err(); err != nil {
		panic(err.Error())
	}

	return buf.Text()
}
