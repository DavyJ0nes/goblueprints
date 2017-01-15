package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

// tlds is a list of top level domains
var tlds = []string{"com", "net", "io"}

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789-_"

func main() {
	// randomness based on microsecond
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		// convert input to lowercase
		text := strings.ToLower(s.Text())
		var newText []rune
		for _, r := range text {
			// convert spaces to dashes
			if unicode.IsSpace(r) {
				r = '-'
			}
			// remove invalid chars
			if !strings.ContainsRune(allowedChars, r) {
				continue
			}
			newText = append(newText, r)
		}
		fmt.Println(string(newText) + "." + tlds[rand.Intn(len(tlds))])
	}
}
