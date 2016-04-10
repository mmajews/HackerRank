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
	receivedText = strings.Trim(receivedText, string(10))
	receivedText = strings.Replace(receivedText, " ", "", -1)
	receivedText = strings.ToLower(receivedText);

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



