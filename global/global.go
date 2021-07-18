package global

import (
	"beego-admin/conf"
	"time"
)

// URL_CURRENT 不做任何操作
const URL_CURRENT = "url://current"

// URL_RELOAD 刷新页面
const URL_RELOAD = "url://reload"

// URL_BACK 返回上一个页面
const URL_BACK = "url://back"

// URL_CLOSE_LAYER 关闭当前layer弹窗
const URL_CLOSE_LAYER = "url://close-layer"

// URL_CLOSE_REFRESH 关闭当前弹窗并刷新父级
const URL_CLOSE_REFRESH = "url://close-refresh"

// LOGIN_USER 登录用户key
const LOGIN_USER = "loginUser"

// LOGIN_USER_ID 登录用户id
const LOGIN_USER_ID = "LoginUserId"

// LOGIN_USER_ID_SIGN 登录用户签名
const LOGIN_USER_ID_SIGN = "loginUserIdSign"

var (
	// BA_CONFIG conf.Server
	BA_CONFIG conf.Server
)

func GetNowTime()  int{
	return int(time.Now().Unix())
}

func GetFormatDate(timestamp int) string{
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	//时间戳转化为日期
	datetime := time.Unix(int64(timestamp), 0).Format(timeLayout)
	return datetime
}
