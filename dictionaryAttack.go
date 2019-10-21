package main

import (
	"bufio"
	"fmt"
	"os"
)

var fileName = "./realhuman_phill.txt"

func Match(pass string) bool {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if line == pass {
			return true
		}
	}
	return false
}

func printPasswords() {
	f, _ := os.Open(fileName)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func main() {
	fmt.Println(Match("~typhlosion~"))
	//printPasswords()
}
