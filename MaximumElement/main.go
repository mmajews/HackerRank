package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	const commandSeparator = " "
	reader := bufio.NewReader(os.Stdin)
	numberOfLinesReadFromInput, _ := reader.ReadString('\n')
	numberOfLinesToBeReceived, _ := strconv.Atoi(numberOfLinesReadFromInput)
	numberOfLinesToBeReceived = trimString(numberOfLinesToBeReceived)

	for i := 0; i < numberOfLinesToBeReceived; i++ {
		line, _ := reader.ReadString('\n')
		var twoStrings[]string
		twoStrings = strings.Split(line, commandSeparator)
		fmt.Println(len(twoStrings))
	}
}

func trimString(stringToTrim string) string {
	var trimmedString string
	trimmedString = strings.TrimSpace(stringToTrim)
	return trimmedString
}
