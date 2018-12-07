package main

import (
	"fmt"
	"os"
)

func main() {

	n, err := fmt.Printf("Hello, World!")
	fmt.Printf("\n\n")

	n = 0
	switch {
	case err != nil:
		os.Exit(1)
	case n == 0:
		fmt.Printf("No bytes output")
		fallthrough
	case n != 13:
		fmt.Printf("Wrong number of characters: %d", n)

	default:
		fmt.Printf("OK!")
	}
	fmt.Printf("\n")

	numOfVowel := 0
	numOfConstant := 0
	zeds := 0
	atoz := "the quick brown fox jumps over the lazy dog"
	for _ , r := range atoz {
		switch r {
		case 'a', 'e','i','o','u':
			numOfVowel++
		case 'z':
			zeds++
			fallthrough
		default:
			numOfConstant++
		}
	}

	fmt.Printf("Number of vowels: %d\tNumber of Constants: %d\t Zeds: %d \n",
		numOfVowel, numOfConstant, zeds)
}

