package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"unicode"
)

type numPos struct {
	begin int
	end   int
}

func checksOut(character rune) bool {
	return unicode.IsSymbol(character) && character != '.'
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	prev := []rune(".......................................................................................................................................................................................................................")
	var prevPositions []numPos

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		text = "." + text + "."
		runes := []rune(text)
		checkingNum := false
		currentPos := 0
		var positions []numPos
		var foundSims []int

		for i := 0; i < len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				if !checkingNum {
					currentPos = i
					checkingNum = true
				}

			} else {
				if checkingNum {
					positions = append(positions, numPos{begin: currentPos, end: i - 1})
					checkingNum = false
				}
				if checksOut(runes[i]) {
					foundSims = append(foundSims, i)
				}
			}

		}
		for _, s := range foundSims {
			fmt.Println(runes[s])
		}
	prevPos:
		for _, n := range prevPositions {
			num, err := strconv.Atoi(string(prev[n.begin : n.end+1]))
			if err != nil {
				continue
			}
			for i := n.begin - 1; i <= n.end+1; i++ {
				if slices.Contains(foundSims, i) {
					sum += num
					continue prevPos
				}
			}
		}
	posIter:
		for _, n := range positions {
			num, err := strconv.Atoi(string(runes[n.begin : n.end+1]))
			if err != nil {
				continue
			}
			for i := n.begin; i <= n.end; i++ {
				if checksOut(prev[i]) {
					sum += num
					continue posIter
				}
			}
			if checksOut(runes[n.begin-1]) || checksOut(prev[n.begin-1]) || checksOut(runes[n.end+1]) || checksOut(prev[n.end+1]) {
				sum += num
				continue posIter
			}
			prevPositions = append(prevPositions, n)

		}
		prev = runes
	}

	fmt.Println(sum)
}
