package api

import (
	"regexp"
	"strconv"

	"github.com/levigross/grequests"
)

type Req struct {
	Pid  int    `json:"pid"`
	Type string `json:"type"`
}

func IppZone(url string) string {
	defer func() string { // 用来处理异常
		if err := recover(); err != nil { // 此处防止错误列表导致程序退出
			return ""
		}
		return ""
	}()

	Itemid := regexp.MustCompile(`post/(\d{1,})`).FindStringSubmatch(url)[1]
	// 直接获取
	pid, _ := strconv.Atoi(Itemid)
	data := Req{Pid: pid, Type: "post"}
	// var data := Req{"pid":Itemid,"type":"post"}
	res, err := grequests.Post("http://share.ippzone.com/ppapi/share/fetch_content", &grequests.RequestOptions{
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		},
		// Cookies: []*http.Cookie,
		UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25",
		JSON:      data,
	})

	if err != nil {
		return "非法请求"
	}
	regs := regexp.MustCompile(`url": "(.*?)"`).FindStringSubmatch(res.String())
	if len(regs) != 2 {
		return ""
	}
	return regs[1]
}
