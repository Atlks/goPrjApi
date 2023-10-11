package lib

import (
	"time"
)

func Timestamp() int64 {
	now := time.Now()
	timestamp := now.Unix()
	return timestamp
}

func CURRENT_TIME_forID() string {

	now := time.Now()

	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan

	return (now.Format("20060102_150405"))
	//获取当前时间
	//时间戳
	//	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式

	//fmt.Println(timeObj)
	//year := timeObj.Year()     //年
	//month := timeObj.Month()   //月
	//day := timeObj.Day()       //日
	//hour := timeObj.Hour()     //小时
	//minute := timeObj.Minute() //分钟
	//second := timeObj.Second() //秒
	////return ("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)

}

func CURRENT_TIME() string {

	now := time.Now()

	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan

	return (now.Format("2006-01-02 15:04:05"))
	//获取当前时间
	//时间戳
	//	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式

	//fmt.Println(timeObj)
	//year := timeObj.Year()     //年
	//month := timeObj.Month()   //月
	//day := timeObj.Day()       //日
	//hour := timeObj.Hour()     //小时
	//minute := timeObj.Minute() //分钟
	//second := timeObj.Second() //秒
	////return ("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)

}
