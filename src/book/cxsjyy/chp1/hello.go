package chp1

import (
	"fmt"
	"os"
)

func SayHello() {
	fmt.Println("hello world!")
}

func OsArgs() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
