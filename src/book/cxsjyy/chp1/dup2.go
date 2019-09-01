package chp1

import (
	"bufio"
	"fmt"
	"os"
)

func FindFileMultiLine(fileName []string) {
	var (
		counts = make(map[string]int)
		num    = 0
	)
	if len(fileName) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range fileName {
			curFile, openErr := os.Open(arg)
			if openErr != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", openErr)
				continue
			}
			countLines(curFile, counts)
			curFile.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			num++
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	if num == 0 {
		fmt.Println("There is no multi line!")
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
