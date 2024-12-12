package main

import (
	"fmt"
	"github.com/faiface/beep"
	"log"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {

	txt := "我需是谁需要指定格式化占位符，而你的代码直接传递了字符变量"
	//	[]string arr:=split2025(txt)

	for _, r := range txt {
		cn := string(r)
		fmt.Println(cn)
		// 打开 MP3 文件
		fname := "C:\\tts3500cnchar\\" + cn + ".mp3"

		// 示例: 使用 try-catch 的方式执行可能会失败的操作
		tryCatch(func() {
			// 这里是需要执行的代码
			play(fname)

		})

	}
	// 打开 MP3 文件
	fname := "C:\\tts3500cnchar\\阿.mp3"

	play(fname)

	// 等待音频播放结束
	select {}

	fmt.Printf("fff")
}

//func split2025(txt string) interface{} {
//
//}

func play(fname string) {
	defer func() {
		// 捕获 panic
		if r := recover(); r != nil {
			fmt.Errorf("play（） panic occurred: %v", r)
		}
	}()

	fmt.Printf("start..")
	fmt.Printf("fname=" + fname)
	f, err := os.Open(fname)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	// 解码 MP3 文件
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Println(err)
	}
	defer streamer.Close()

	// 初始化音频播放设备
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// 播放音频
	//	speaker.Play(streamer)

	// 播放音频
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	// 等待音频播放完成
	<-done
	fmt.Println("Audio playback finished.")
}

// tryCatch 模拟 try-catch 的功能
func tryCatchRt(f func() (interface{}, error)) (result interface{}, err error) {
	defer func() {
		// 捕获 panic
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()
	// 调用目标函数
	return f()
}

// tryCatch 模拟 try-catch 的功能
func tryCatch(f func()) {
	defer func() {
		// 捕获 panic
		if r := recover(); r != nil {
			//err = fmt.Errorf("panic occurred: %v", r)
			fmt.Errorf("panic occurred: %v", r)
		}
	}()
	// 调用目标函数
	f()
}
