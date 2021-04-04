package api

import (
	"regexp"
	"strings"

	"github.com/levigross/grequests"
)

func ZuiYou(url string) string {
	defer func() string { // 用来处理异常
		if err := recover(); err != nil { // 此处防止错误列表导致程序退出
			return ""
		}
		return ""
	}()
	// 直接获取
	res, err := grequests.Get(strings.Replace(url, "post&", "post?", -1), &grequests.RequestOptions{
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
		UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25",
	})

	if err != nil {
		return "无效请求"
	}
	// fmt.Println(res.String())
	regs := regexp.MustCompile(`playcnt":\d{1,},"url":"(.*?)"`).FindStringSubmatch(res.String())
	if len(regs) != 2 {
		return ""
	}
	return strings.ReplaceAll(regs[1], "\\u002F", "/")
}
