package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type numPos struct {
	begin int
	end   int
	value int
}

func checksOut(character rune) bool {
	return unicode.IsSymbol(character) && character != '.'
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	var prevPositions []numPos
	prevSymbols := make([]bool, 1000)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		text = "." + text + "."
		symbols := make([]bool, len(text))
		var positions []numPos
		checkingNum := false
		currentPos := 0
	posLoop:
		for i, v := range text {
			if unicode.IsDigit(v) {
				if !checkingNum {
					currentPos = i
					checkingNum = true
				}

			} else {
				if v != '.' && !unicode.IsLetter(v) {
					symbols[i] = true
				}
				if checkingNum {
					end := i - 1
					checkingNum = false
					value, err := strconv.Atoi(text[currentPos:i])
					if err != nil {
						continue
					}
					if symbols[i] {
						sum += value
						continue
					}
					if symbols[currentPos-1] {
						sum += value
						continue
					}
					for j := currentPos - 1; j <= end+1; j++ {
						if prevSymbols[j] {
							sum += value
							continue posLoop
						}
					}
					positions = append(positions, numPos{begin: currentPos, end: end, value: value})
				}

			}

		}
	prevLoop:
		for _, pos := range prevPositions {
			for i := pos.begin - 1; i <= pos.end+1; i++ {
				if symbols[i] {
					sum += pos.value
					continue prevLoop
				}
			}
		}
		prevSymbols = symbols
		prevPositions = positions
	}

	fmt.Println(sum)
}
