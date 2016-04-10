package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	receivedText, _ := reader.ReadString('\n')
	receivedText = deleteAllWhiteSpacesAndTrimString(receivedText)
	detectIfPangram(receivedText)
}

func deleteAllWhiteSpacesAndTrimString(stringToTrim string) string{
	stringToTrim = strings.Trim(stringToTrim, string(10))
	stringToTrim = strings.Replace(stringToTrim, " ", "", -1)
	stringToTrim = strings.ToLower(stringToTrim);
	return stringToTrim
}

func detectIfPangram(receivedText string){
	//Imitating set with map, golang lack of so
	allLetters := make(map[byte]bool)
	for i := 0; i < len(receivedText); i++ {
		allLetters[receivedText[i]] = true
	}

	if len(allLetters) == 26 {
		fmt.Println("pangram")
	} else {
		fmt.Println("not pangram")
	}
}





