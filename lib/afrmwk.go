package lib

import (
	"fmt"
	"io/ioutil"
	"log"
)

// evtBoot sets up error handling and starts a new goroutine to play an MP3 file
func EvtBoot(actBiz func()) {
	// Set up error handler (use recover in Go)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Print the logo (dummy function for demonstration)
	printLogo()

	// Start a new goroutine to play the MP3 file
	go func() {
		fmt.Println("New goroutine started")
		PlayMp3("C:\\Users\\Administrator\\OneDrive\\song cn\\新疆美丽公主组合 - 欢乐地跳吧.mp3", 5)

		fmt.Println("playMp3somesec goroutine finished")
	}()

	actBiz()
}

// playMp3somesec plays an MP3 file for a given duration (in seconds)
//func playMp3somesec(filePath string, lenx int) error {
//	PlayMp3(filePath, lenx)
//	time.Sleep(time.Duration(lenx) * time.Second)
//	return nil
//}

func printLogo() {
	fmt.Println(`
      ,--./,-.
     / #      \
    |          |
     \        /    
      ` + "`" + `._,._,'
         ]
      ,--'
      |
      ` + "`" + `.___.
      `)

	content, err := ioutil.ReadFile("D:\\0prj\\mdsj\\mdsjprj\\bin\\Debug\\net8.0\\logo.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	fmt.Println(string(content))
}
