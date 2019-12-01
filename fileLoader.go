package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func testing() {
	file, err := os.Open("abc.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
