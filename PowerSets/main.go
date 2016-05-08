package PowerSets

import (
	"bufio"
	"os"
	"fmt"
)

func main(){
	reader := bufio.NewReader(os.Stdin);
	receivedText, _ := reader.ReadString("\n")
	allPowerSets := generateAllPowerSets(receivedText)
	for _, value := range allPowerSets{
		fmt.Println(value)
	}
}

func generateAllPowerSets(input string) map[string]bool{
	powerSets := make(map[string]bool)

}
