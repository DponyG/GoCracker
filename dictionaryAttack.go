package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var fileName = "./abc.txt"
var fileName2 = "./realhuman_phill.txt"

//GenerateCombinations generates the combinations and uses channels
//Source:https://stackoverflow.com/questions/19249588/go-programming-generating-combinations
//Because my implementations resulted in 16GB of memory going byebye in the blink of an eye
func GenerateCombinations(alphabet string, length int) <-chan string {
	c := make(chan string)

	// Starting a separate goroutine that will create all the combinations,
	// feeding them to the channel c
	go func(c chan string) {
		defer close(c) // Once the iteration function is finished, we close the channel

		AddLetter(c, "", alphabet, length) // We start by feeding it an empty string
	}(c)

	return c // Return the channel to the calling function
}

// AddLetter adds a letter to the combination to create a new combination.
// This new combination is passed on to the channel before we call AddLetter once again
// to add yet another letter to the new combination in case length allows it
func AddLetter(c chan string, combo string, alphabet string, length int) {
	// Check if we reached the length limit
	// If so, we just return without adding anything
	if length <= 0 {
		return
	}

	var newCombo string
	for _, ch := range alphabet {
		newCombo = combo + string(ch)
		c <- newCombo
		AddLetter(c, newCombo, alphabet, length-1)
	}
}

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
		start := time.Now()
		fmt.Println("Password Found On Big List:", Match(scanner.Text(), true))
		fmt.Println("This took: ", time.Since(start))
		start = time.Now()
		fmt.Println("Password Found On Special List:", Match(scanner.Text(), false))
		fmt.Println("This took: ", time.Since(start))
		start = time.Now()
		for combination := range GenerateCombinations("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()1234567890[]{}\\|;':\",./<>?", len(scanner.Text())) {
			if scanner.Text() == combination {
				fmt.Println("Password found through brute force: ", scanner.Text())
				fmt.Println("This took: ", time.Since(start))
				break
			}
		}
		fmt.Println("+--------------------------------------------------+")
		fmt.Println("Enter some more passwords for us to crack, please :)")
	}
}
