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
	var prevMap map[int][]int

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		text = "." + text + "."
		symbols := make([]bool, len(text))
		curMap := make(map[int][]int)
		var positions []numPos
		checkingNum := false
		currentPos := 0
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
						curMap[i] = append(curMap[i], value)
					}
					if symbols[currentPos-1] {
						curMap[currentPos-1] = append(curMap[currentPos-1], value)
					}
					for j := currentPos - 1; j <= end+1; j++ {
						if prevSymbols[j] {
							prevMap[j] = append(prevMap[j], value)
						}
					}
					positions = append(positions, numPos{begin: currentPos, end: end, value: value})
				}

			}

		}
		for _, pos := range prevPositions {
			for i := pos.begin - 1; i <= pos.end+1; i++ {
				if symbols[i] {
					curMap[i] = append(curMap[i], pos.value)
				}
			}
		}
		for _, l := range prevMap {
			if len(l) == 2 {
				sum += l[0] * l[1]
			}
		}
		prevSymbols = symbols
		prevPositions = positions
		prevMap = curMap
	}
	for _, l := range prevMap {
		if len(l) == 2 {
			sum += l[0] * l[1]
		}
	}

	fmt.Println(sum)
}
