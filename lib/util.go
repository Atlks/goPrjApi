package lib

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

// 日志函数
func setDbgFunEnter(methodName string, args ...interface{}) {
	log.Printf("Entering method: %s with args: %v\n", methodName, args)
}

func setDbgValRtval(methodName string, retval any) {
	log.Printf("Exiting method: %s with return value: %d\n", methodName, retval)
}

// 延时执行函数
func executeAfterDelay(delay time.Duration, action func()) {
	time.Sleep(delay)
	action()
}

func LoopForever() {
	for {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(5 * time.Second)
	}
}

// 播放MP3函数
func PlayMp3(mp3FilePath string, lenx int) {
	methodName := "PlayMp3"
	setDbgFunEnter(methodName, mp3FilePath)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in PlayMp3", r)
		}
	}()

	file, err := os.Open(mp3FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	streamer, format, err := mp3.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		fmt.Println("Playback finished")
	})))

	fmt.Println("Playing... Press any key to stop.")
	// 延时60秒
	time.Sleep(time.Duration(lenx) * time.Second)
	//
	//// 延时5秒后停止播放
	//executeAfterDelay(5*time.Second, func() {
	//	speaker.Clear()
	//	fmt.Println("Stopped playing.")
	//})

	setDbgValRtval(methodName, 0)
}

//func main() {
//	playMp3somesec("path_to_your_mp3_file.mp3")
//}
