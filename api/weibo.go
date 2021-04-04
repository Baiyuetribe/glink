package api

import (
	"regexp"
	"strings"

	"github.com/levigross/grequests"
)

func WeiBo(url string) string {
	defer func() string { // 用来处理异常
		if err := recover(); err != nil { // 此处防止错误列表导致程序退出
			return ""
		}
		return ""
	}()
	Itemid := regexp.MustCompile(`(\d{1,}:\d{1,})`).FindStringSubmatch(url)[1]
	// 直接获取
	res, err := grequests.Get("https://video.h5.weibo.cn/s/video/object?object_id="+Itemid, &grequests.RequestOptions{
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		},
		UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25",
	})

	if err != nil {
		return "非法请求"
	}
	regs := regexp.MustCompile(`hd_url":"(.*?)"`).FindStringSubmatch(res.String())
	if len(regs) != 2 {
		return ""
	}
	return strings.ReplaceAll(regs[1], "\\", "")
}
