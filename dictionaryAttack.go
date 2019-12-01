package main

import (
	"bufio"
	"fmt"
	"os"
)

var fileName = "./abc.txt"
var fileName2 = "./realhuman_phill.txt"

//Match returns true if the given pass is found in the Big text file
func Match(pass string, bigMatch bool) bool {
	if bigMatch {
		f, _ := os.Open(fileName)
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()
			if line == pass {
				return true
			}
		}
	}

	if !bigMatch {
		f, _ := os.Open(fileName2)
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()
			if line == pass {
				return true
			}
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

func consoleInput() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter some text: ")
	for scanner.Scan() {
		fmt.Println("Password Found On Big List:", Match(scanner.Text(), true))
		fmt.Println("Password Found On Special List:", Match(scanner.Text(), false))
	}
}
