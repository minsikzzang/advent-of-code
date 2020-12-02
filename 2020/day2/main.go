package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", err))
	}

	input := string(data)
	lines := strings.Split(input, "\n")

	// numValid := numValidPasswordPart1(lines)
	numValid := numValidPasswordPart2(lines)
	fmt.Printf("%d number passwords are valid\n", numValid)
	return
}

func numValidPasswordPart2(lines []string) int {
	numValid := 0

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		r := strings.Split(tokens[0], "-")
		f, _ := strconv.Atoi(r[0])
		s, _ := strconv.Atoi(r[1])

		letter := tokens[1][:1]
		password := tokens[2]
		matched := 0

		if password[f-1:f] == letter {
			matched++
		}

		if password[s-1:s] == letter {
			matched++
		}

		if matched == 1 {
			numValid++
		}
	}

	return numValid
}

func numValidPasswordPart1(lines []string) int {
	numValid := 0

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		r := strings.Split(tokens[0], "-")
		min, _ := strconv.Atoi(r[0])
		max, _ := strconv.Atoi(r[1])

		letter := tokens[1][:1]
		password := tokens[2]

		matched := 0

		for _, c := range password {
			if c == rune(letter[0]) {
				matched++
			}
		}

		if max >= matched && min <= matched {
			numValid++
		}
	}

	return numValid
}
