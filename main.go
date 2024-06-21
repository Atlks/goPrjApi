package main

import (
	"awesomeProject/lib"
	"awesomeProject/libx"
	"awesomeProject/testpkg"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"path/filepath"
	"time"
)

func main() {
	defer lib.HandlePanic()

	//path := "C:\\Users\\Administrator\\OneDrive\\mklv song lst\\Kim Chiu-Wala Man Sa'yo Ang Lahat.mp3"
	//lib.PlayMp3(path, 30)
	//time.Sleep(60 * time.Second)
	//fmt.Printf("....finish...")
	//return

	lib.EvtBoot(func() {})
	// 获取当前时间
	currentTime := time.Now()

	// 打印当前时间
	fmt.Println("Current time:", currentTime.Format("2006-01-02 15:04:05"))
	// Example usage of the save function

	//-------------------- 设置间隔时间为 20min
	interval := 18 * 60 * time.Second

	// 使用匿名函数作为定时器函数
	timerFunc := func() {
		defer lib.HandlePanic()
		//	counter++
		path := "C:\\Users\\Administrator\\OneDrive\\mklv song lst\\Kim Chiu-Wala Man Sa'yo Ang Lahat.mp3"
		lib.PlayMp3(path, 60)
		time.Sleep(60 * time.Second)

		//	fmt.Printf("Function executed %d times at: %s\n", counter, time.Now().Format("2006-01-02 15:04:05"))
	}

	// 启动定时器，定时执行myFunction
	go lib.StartTimer(interval, timerFunc)

	// 保持主函数运行
	select {}
	//	lib.BuyEth()
	//botMsgRcvrHdlr()
	lib.LoopForever()
	//	funcName222()
}

func funcName222() {
	testpkg.Mmm()
	testSave()

	fmt.Print(111)
	libx.MthdFrmPkg1()

	// 创建 session 数据
	chtsSesss := map[string]interface{}{
		"id": "1",
		"nm": "tommm",
	}

	dbFileName := "example2024.db"

	// 调用 save 方法

	libx.Save(chtsSesss, dbFileName)

	dbFileName = "example2024.json"
	//libx.SaveJson(chtsSesss, dbFileName)
	list := libx.QryJson(dbFileName)
	fmt.Print(libx.JsonEncode(list))
	//m()
}

func botMsgRcvrHdlr() {
	botToken := "6999501721:AAFNqa2YZ-lLZMfN8T2tYscKBi33noXhdJA" // 替换为你的Bot Token
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		//	msgJsonStr := lib.Json_encode(update)
		lib.FwrtLgTypeDate("msgrcvDir2024", update)
		//lib.Ffile_put_contents("msgrcv", msgJsonStr, false)

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "欢迎光临")
		//_, err := bot.Send(msg)
		//if err != nil {
		//	log.Printf("Failed to send message: %v", err)
		//}

	}
}

func testSave() {
	// 初始化变量
	saveDataDir := "测试数据表目录"
	dbg := make(map[string]interface{})
	sortedListNew := map[string]interface{}{
		"国家":   "China",
		"name": "tomm", "id": "007",
	}

	// 调用函数

	strEngr := func(row map[string]interface{}) int {

		prtnKey := "国家"
		wrtFile := filepath.Join(saveDataDir, fmt.Sprintf("%v.json", row[prtnKey]))
		//使用增量还是全量模式，都是由存储引擎决定的，orm框架与查询引擎是不管的，主管业务层面
		libx.SaveJson(row, wrtFile)
		return 0
	}
	str := libx.Qe_save(sortedListNew, saveDataDir, strEngr, dbg)
	fmt.Print(str)
}
