package submerge

import (
	"fmt"
	"sort"
	"strings"
)

type subLine struct {
	Num   int
	Time  string
	Text1 string
	Text2 string
}

func (s *subLine) isAfter(sub2 *subLine) bool {
	if sub2 == nil {
		return false
	}
	if s == nil {
		return true
	}
	times := []string{s.Time, sub2.Time}
	sort.Strings(times)
	return times[1] == s.Time
}

func (s *subLine) String() string {
	return fmt.Sprintf("[%d] %s   ((  %s |  %s ))", s.Num, s.Time, s.Text1, s.Text2)
}

func (s *subLine) toFormat() string {
	wr := strings.Builder{}
	wr.Write([]byte(fmt.Sprintf("%d\n", s.Num)))
	wr.Write([]byte(fmt.Sprintf("%s\n", s.Time)))
	wr.Write([]byte(fmt.Sprintf("%s\n", s.Text1)))
	if s.Text2 != "" {
		wr.Write([]byte(fmt.Sprintf("%s\n", s.Text2)))
	}
	wr.Write([]byte("\n"))
	return wr.String()
}

func adjustNumsForSortedSlice(lines []*subLine) {
	for i, line := range lines {
		if line == nil {
			fmt.Printf("\n Missing: %d", i)
			continue
		}
		line.Num = i
	}
}
