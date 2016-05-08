package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin);
	receivedText, _ := reader.ReadString('\n')
	allPowerSets := generateAllPowerSets(trimString(receivedText))
	for _, value := range allPowerSets{
		fmt.Println(value)
	}
}

func generateAllPowerSets(input string) map[string]bool{
	powerSets := make(map[string]bool)
	characters := make([]rune,0,0);
	//lenOfString := len(input)
	for _, r := range input{
		characters = append(characters,r)
	}

	for _, value := range characters{
		fmt.Println(string(value))
	}

	return powerSets
}


func trimString(stringToTrim string) string {
	var trimmedString string
	trimmedString = strings.TrimSpace(stringToTrim)
	return trimmedString
}
