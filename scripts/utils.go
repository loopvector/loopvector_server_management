package main

import (
	"bufio"
	"strings"
)

func getInputString(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		return ""
	}
	input := strings.TrimSpace(scanner.Text())
	return input
}
