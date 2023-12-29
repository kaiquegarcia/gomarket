package util

import "fmt"

// TryableFunc is a function that returns true if it could run successfully or false if not. If you return a filled string, it will be prompted before recalling the function.
type TryableFunc func() (bool, string)

// Try will try to run the {callable} until it succeed or for {maxAttempts} times. If any of the attempts succeed, it will panic.
func Try(maxAttempts int, callable TryableFunc) {
	attempts := 0
	for {
		success, msg := callable()
		if success {
			return
		}

		if attempts >= maxAttempts {
			FinishByTooManyErrors()
		}

		fmt.Println(msg)
		attempts++
	}
}
