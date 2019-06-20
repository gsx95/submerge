package submerge

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func MergeSubs(file1Path, file2Path string) string {
	file1 := openFile(file1Path)
	file2 := openFile(file2Path)
	defer closeFile(file1)
	defer closeFile(file2)

	lines := parseSubFile(file1)
	lines2 := parseSubFile(file2)


	lines = append(lines, lines2...)
	sort.Slice(lines, func(i, j int) bool {
		return lines[j].isAfter(lines[i])
	})

	adjustNumsForSortedSlice(lines)
	return writeLinesToString(lines)
}

func writeLinesToString(lines []*subLine) string {

	w := strings.Builder{}
	for _, line := range lines {
		//fmt.Printf("\n [%d]   %s  ", i, line.Time)
		s := line.toFormat()
		fmt.Println(s)
		_, err := w.WriteString(s)
		if err != nil {
			panic(err)
		}
	}
	return w.String()
}

func writeLinesToFile(lines []*subLine, outPath string) {
	f, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)
	for _, line := range lines {
		//fmt.Printf("\n [%d]   %s  ", i, line.Time)
		s := line.toFormat()
		fmt.Println(s)
		_, err := w.WriteString(s)
		if err != nil {
			panic(err)
		}
	}
	err = w.Flush()
	if err != nil {
		panic(err)
	}
}

func writeNums(lines []*subLine) {
	for _, line := range lines {
		fmt.Println(line.Num)
	}
}

func writeTimes(lines []*subLine) {
	for _, line := range lines {
		fmt.Println(line.Time)
	}
}

func printMissingNums(lines []*subLine) {
	lastNum := -1
	for _, line := range lines {
		if line == nil {
			continue
		}
		if lastNum == -1 {
			lastNum = line.Num
			continue
		}
		if line.Num != lastNum + 1 {
			fmt.Println(lastNum + 1)
		}
		lastNum = line.Num
	}
}



