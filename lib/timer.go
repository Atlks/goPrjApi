package lib

import "time"

/*

// 设置间隔时间为30秒
	interval := 30 * time.Second

	// 启动定时器，定时执行myFunction
	go startTimer(interval, myFunction)

*/
// 定时器函数，每隔指定时间间隔执行一次传入的函数
func StartTimer(interval time.Duration, function func()) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			function()
		}
	}
}
