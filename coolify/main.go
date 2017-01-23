package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// adding constants here for readabilty
const (
	duplicateVowel bool = true
	removeVowel    bool = false
)

// randBool returns either true or false.
// There is 50/50 chance for either
func randBool() bool {
	return rand.Intn(2) == 0
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		// convert input into byte slice
		word := []byte(s.Text())
		// are we going to mutate Scanned line or not?
		if randBool() {
			var vI int = -1
			// iterate over all characters
			for i, char := range word {
				switch char {
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					// choose which vowel to change, rather than always the same one
					if randBool() {
						vI = i
					}
				}
			}
			if vI >= 0 {
				// another random decision to duplicate or remove the vowel
				switch randBool() {
				case duplicateVowel:
					word = append(word[:vI+1], word[vI:]...)
				case removeVowel:
					word = append(word[:vI], word[vI+1:]...)
				}
			}
		}
		// simple output
		fmt.Println(string(word))
	}
}
