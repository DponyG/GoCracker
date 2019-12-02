package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



var fileName = "./abc.txt"
var fileName2 = "./realhuman_phill.txt"



//1112114



//SOURCE: https://github.com/mxschmitt/golang-combinations/blob/master/combinations.go
// Package combinations provides a method to generate all combinations out of a given string array.

// All returns all combinations for a given string array.
// This is essentially a powerset of the given set except that the empty set is disregarded.
func All(set []string) (subsets [][]string) {
	length := uint(len(set))

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []string

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}

func brute(pass string) int {
	//combinations := All([]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","t","u","v","w","x","y","z",
	//		"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z",
	//		"!","@","#","$","%","^","&","*","(",")","-","=","_","+","`","~","[","]","{","}",";",":",",","<",".",">","/","?",
	//		"1","2","3","4","5","6","7","8","9","0"})
	combinations := All([]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","t","u","v","w","x","y","z",
		"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z",
		"1","2","3","4","5","6","7","8","9","0"})

	fmt.Println(combinations)
	for _, s := range combinations {
		justString := strings.Join(s,"")
		println(justString)
		if justString == pass {
			println("Hallelujah! the string is ", justString)
			break
		}

	}

	return 0

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
		//fmt.Println("Password Found On Big List:", Match(scanner.Text(), true))
		//fmt.Println("Password Found On Special List:", Match(scanner.Text(), false))

		fmt.Println("Password Found On Special List:", brute(scanner.Text()))
	}
}
