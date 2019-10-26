package submerge

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Delay struct {
	Hours int64
	Mins int64
	Secs int64
	Ms int64
	TextTooEarly bool
}

type SubConfig struct {
	FilePath string
	*Delay
	Color string
}

type Config struct {
	Sub1 SubConfig
	Sub2 SubConfig
}

func MergeSubs(conf Config) (string, error) {
	file1, err := openFile(conf.Sub1.FilePath)
	if err != nil {
		return "", err
	}
	file2, err := openFile(conf.Sub2.FilePath)
	if err != nil {
		return "", err
	}
	defer closeFile(file1)
	defer closeFile(file2)

	lines, err := parseSubFile(file1)
	if err != nil {
		return "", err
	}
	lines2, err := parseSubFile(file2)
	if err != nil {
		return "", err
	}

	if conf.Sub1.Delay != nil {
		addDelay(lines, conf.Sub1.Delay)
	}

	if conf.Sub1.Color != "" {
		addColor(lines, conf.Sub1.Color)
	}

	if conf.Sub2.Delay != nil {
		addDelay(lines2, conf.Sub2.Delay)
	}

	if conf.Sub2.Color != "" {
		addColor(lines2, conf.Sub2.Color)
	}


	lines = append(lines, lines2...)
	sort.Slice(lines, func(i, j int) bool {
		return lines[j].isAfter(lines[i])
	})

	adjustNums(lines)
	return writeLinesToString(lines), nil
}

func writeLinesToString(lines []*subLine) string {

	w := strings.Builder{}
	for _, line := range lines {
		//fmt.Printf("\n [%d]   %s  ", i, line.Time)
		s := line.toFormat()
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



