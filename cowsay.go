package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

const cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

func buildBalloon(lines []string, maxwidth int) string {
	var borders []string
	count := len(lines)
	var ret []string

	borders = []string{"/", "\\", "\\", "/", "|", "<", ">"}

	top := " " + strings.Repeat("_", maxwidth+2)
	bottom := " " + strings.Repeat("-", maxwidth+2)

	ret = append(ret, top)
	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[5], lines[0], borders[6])
		ret = append(ret, s)
	} else {
		s := fmt.Sprintf(`%s %s %s`, borders[0], lines[0], borders[1])
		ret = append(ret, s)
		i := 1
		for ; i < count-1; i++ {
			s = fmt.Sprintf(`%s %s %s`, borders[4], lines[i], borders[4])
			ret = append(ret, s)
		}
		s = fmt.Sprintf(`%s %s %s`, borders[2], lines[i], borders[3])
		ret = append(ret, s)
	}

	ret = append(ret, bottom)
	return strings.Join(ret, "\n")
}

func tabsToSpaces(lines []string) []string {
	var ret []string
	for _, l := range lines {
		l = strings.Replace(l, "\t", "    ", -1)
		ret = append(ret, l)
	}
	return ret
}

func calculateMaxWidth(lines []string) int {
	w := 0
	for _, l := range lines {
		size := utf8.RuneCountInString(l)
		if size > w {
			w = size
		}
	}
	return w
}

func normalizeStringsLength(lines []string, maxWidth int) []string {
	var ret []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", maxWidth-utf8.RuneCountInString(l))
		ret = append(ret, s)
	}
	return ret
}

func main() {
	info, _ := os.Stdin.Stat()
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | cowsay")
		return
	}

	var lines []string

	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}

	lines = tabsToSpaces(lines)
	maxWidth := calculateMaxWidth(lines)
	messages := normalizeStringsLength(lines, maxWidth)
	balloon := buildBalloon(messages, maxWidth)
	fmt.Println(balloon)
	fmt.Println(cow)
	fmt.Println()
}
