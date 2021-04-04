package api

import (
	"regexp"

	"github.com/levigross/grequests"
)

// type Req struct {
// 	Feedids []string `json:"feedids"`
// }

func MonMo(url string) string {
	defer func() string { // 用来处理异常
		if err := recover(); err != nil { // 此处防止错误列表导致程序退出
			return ""
		}
		return ""
	}()
	// 直接获取
	s := regexp.MustCompile(`share-v2/(\w\w\d{1,})`).FindStringSubmatch(url)
	if len(s) != 2 {
		return ""
	}
	// data := Req{Feedids:[]string{s[1]}}
	res, err := grequests.Post("https://m.immomo.com/inc/microvideo/share/profiles", &grequests.RequestOptions{
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		},
		UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25",
		Params:    map[string]string{"feedids": s[1]},
	})

	if err != nil {
		return "无效请求"
	}
	// re.findall("playurl_video\":\"(.*?)\"",r.text)[0] 该接口已失效，videmo_url为空
	regs := regexp.MustCompile(`video_url":"(.*?)"`).FindStringSubmatch(res.String())
	if len(regs) != 2 {
		return ""
	}
	return regs[1]
}
