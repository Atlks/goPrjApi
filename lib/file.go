package lib

import (
	"bufio"
	"os"
)

func ReadLine(f string, lineHdlr func(line string)) {
	readFile, _ := os.Open(f)

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lineHdlr(fileScanner.Text())
	}
}

func Write(f string, data string) {

	file, err := os.OpenFile(f, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		panic(err)
	}
}
