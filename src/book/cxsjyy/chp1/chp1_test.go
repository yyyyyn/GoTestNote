package chp1

import "testing"

func TestOsArgs(t *testing.T) {
	OsArgs()
}

func TestForRange(t *testing.T) {
	ForRange()
}

//func TestFindMultiLine(t *testing.T) {
//	FindMultiLine()
//}

func TestFindFileMultiLine(t *testing.T) {
	file1 := "C:\\Users\\Administrator\\go\\src\\GoTestNote\\note.txt"
	files := []string{}
	files = append(files, file1)
	FindFileMultiLine(files)
}
