package submerge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseSubFile(file *os.File) []*subLine {
	var lines []*subLine

	sc := bufio.NewScanner(file)

	nextLine := true
	for nextLine {
		line, notEmpty := parseSubLine(sc)
		nextLine = notEmpty
		lines = append(lines, line)
	}
	return lines
}

func parseSubLine(sc *bufio.Scanner) (*subLine, bool) {
	counter := 0
	var currentLine *subLine
	for sc.Scan() {

		if err := sc.Err(); err != nil {
			panic(err)
		}
		line := strings.TrimSpace(sc.Text())
		line = strings.Replace(line, "\ufeff", "", -1)
		if line == "" {
			return currentLine, true
		}

		switch counter {
		case 0:
			num, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				fmt.Println(line)
				panic(err)
			}
			currentLine = &subLine{Num: int(num)}
		case 1:
			currentLine.Time = line
		case 2:
			currentLine.Text1 = line
		case 3:
			currentLine.Text2 = line
		}
		counter++
	}
	return currentLine, false
}
