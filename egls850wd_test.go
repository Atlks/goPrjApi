package main

import (
	"fmt"
	"goapiPrj/lib"
	"testing"
)

func TestHandler100(t *testing.T) {

	f := "c:/w/egls850wd.txt"

	f2 := "egls850wdAbbr.txt"
	lib.ReadLine(f, func(line string) {
		fmt.Println(line)
		if lib.Len(line) < 3 {
			return
		}
		firstCh := lib.Left(line, 1)
		other := lib.Substr(line, 1)
		lib.Write(f2, line+"\n")
	})
	//readFile, _ := os.Open(f)
	//
	//fileScanner := bufio.NewScanner(readFile)
	//
	//fileScanner.Split(bufio.ScanLines)
	//
	//for fileScanner.Scan() {
	//	fmt.Println(fileScanner.Text())
	//}

}
