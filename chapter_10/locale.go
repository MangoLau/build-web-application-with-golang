package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var locales map[string]map[string]string

func main() {
	locales = make(map[string]map[string]string, 2)
	en := make(map[string]string, 10)
	en["pea"] = "pea"
	en["bean"] = "bean"
	locales["en"] = en
	cn := make(map[string]string)
	cn["pea"] = "豌豆"
	cn["bean"] = "毛豆"
	locales["zh-CN"] = cn
	lang := "zh-CN"
	fmt.Println(msg(lang, "pea"))
	fmt.Println(msg(lang, "bean"))

	// 本地化日期和时间
	en["time_zone"] = "America/Chicago"
	cn["time_zone"] = "Asia/Shanghai"

	loc, _ := time.LoadLocation(msg(lang, "time_zone"))
	t := time.Now()
	t = t.In(loc)
	fmt.Println(t.Format(time.RFC3339))

	en["date_format"] = "%Y-%m-%d %H:%M:%S"
	cn["date_format"] = "%Y年%m月%d日 %H时%M分%S秒"

	fmt.Println(date(msg(lang, "date_format"), t))

	// 本地化货币值
	en["money"] = "USD %d"
	cn["money"] = "￥%d元"
	fmt.Println(money_format(msg(lang, "money"), 100))
}

func msg(locale, key string) string {
	if v, ok := locales[locale]; ok {
		if v2, ok := v[key]; ok {
			return v2
		}
	}
	return ""
}

func date(fomate string, t time.Time) string {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	//解析相应的%Y %m %d %H %M %S然后返回信息
	//%Y 替换成2012
	//%m 替换成10
	//%d 替换成24
	res := fomate
	fmt.Printf("%v %v %v %v %v %v", year, month, day, hour, min, sec)
	fmt.Println()
	res = strings.Replace(res, "%Y", strconv.Itoa(year), -1)
	res = strings.Replace(res, "%m", strconv.Itoa(int(month)), -1)
	res = strings.Replace(res, "%d", strconv.Itoa(day), -1)
	res = strings.Replace(res, "%H", strconv.Itoa(hour), -1)
	res = strings.Replace(res, "%M", strconv.Itoa(min), -1)
	res = strings.Replace(res, "%S", strconv.Itoa(sec), -1)
	return res
}

func money_format(fomate string, money int64) string {
	return fmt.Sprintf(fomate, money)
}
