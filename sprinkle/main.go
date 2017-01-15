package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord + " app",
	otherWord + " site",
	otherWord + " time",
	otherWord + " hq",
	otherWord + " hero",
	"get " + otherWord,
	"go " + otherWord,
	"lets " + otherWord,
}

func main() {
	// used to increase the degree of randomness
	rand.Seed(time.Now().UTC().UnixNano())
	// listen to stdin
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		// grabs random entry in transforms slice
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
