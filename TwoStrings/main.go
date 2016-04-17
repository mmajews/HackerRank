package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	receivedText, _ := reader.ReadString('\n')

	var numberOfTestCases int
	receivedText = trimString(receivedText)
	numberOfTestCases, err := strconv.Atoi(receivedText)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	for i := numberOfTestCases; i > 0; i-- {
		var twoStrings [2]string
		twoStrings = getTwoStringsFromInput(reader)
		detectIfCommonSubStringExists(twoStrings[0], twoStrings[1])
	}
}

func trimString(stringToTrim string) string {
	var trimmedString string
	trimmedString = strings.TrimSpace(stringToTrim)
	return trimmedString
}

func getTwoStringsFromInput(reader *bufio.Reader) [2]string {
	var arrayToReturn [2]string
	for i := 1; i >= 0; i-- {
		receivedText, _ := reader.ReadString('\n')
		receivedText = trimString(receivedText)
		arrayToReturn[i] = receivedText;
	}
	return arrayToReturn;
}

func detectIfCommonSubStringExists(firsString string, secondString string) {
	firstStringLetters := make(map[byte]bool)
	for i := 0; i < len(firsString); i++ {
		firstStringLetters[firsString[i]] = true
	}

	for i := 0; i < len(secondString); i++ {
		if firstStringLetters[secondString[i]] == true {
			fmt.Println("YES")
			return
		}
	}

	fmt.Println("NO")
}

