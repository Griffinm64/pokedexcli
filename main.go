package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		 if scanner.Scan() != false {
			dirtyText := scanner.Text()
			cleandedtext := cleanInput(dirtyText)
			fmt.Println("Your command was: " + cleandedtext[0])
		 }

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
