package api

import (
	"regexp"

	"github.com/levigross/grequests"
)

// https://www.bilibili.com/video/BV1CE411H7bQ?t=351
func Bilibili(url string) string {
	defer func() string { // 用来处理异常
		if err := recover(); err != nil { // 此处防止错误列表导致程序退出
			return ""
		}
		return ""
	}()
	// 直接获取

	s := regexp.MustCompile(`(BV[0-9a-zA-Z]*)`).FindStringSubmatch(url)
	if len(s) != 2 {
		return ""
	}

	res, err := grequests.Get("https://m.bilibili.com/video/"+s[1], &grequests.RequestOptions{
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		},
		UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25",
	})

	if err != nil {
		return "无效请求"
	}
	regs := regexp.MustCompile(`readyVideoUrl: '(.*?)',`).FindStringSubmatch(res.String())
	if len(regs) != 2 {
		return ""
	}
	return regs[1]
}
