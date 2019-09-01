package chp1

import (
	"bufio"
	"fmt"
	"os"
)

func SayHello() {
	fmt.Println("hello world!")
}

func OsArgs() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("Args", i, "is:", os.Args[i])
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func ForRange() {
	for num, arg := range os.Args {
		fmt.Printf("os.Args[%d] is:%s\n", num, arg)
	}
}

func FindMultiLine() {
	// 找出标准输入中出现次数大于1的行
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t %s\n", n, line)
		}
	}
}
