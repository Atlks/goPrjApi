package 类库包

import "goapiPrj/lib"

func F当前时间() string {
	return lib.CURRENT_TIME()
}

func F时间戳() int64 {
	return lib.Timestamp()
}
