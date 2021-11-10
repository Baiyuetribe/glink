package api

import (
	"regexp"
	"strings"

	"github.com/levigross/grequests"
)

func DouYin(url string) string {
	defer func() string { // 用来处理异常
		if err := recover(); err != nil { // 此处防止错误列表导致程序退出
			return ""
		}
		return ""
	}()
	// re.findall("(https://v.douyin.com/.*?/)",a) 改进匹配算法
	realUrl := regexp.MustCompile(`(http.*://[^\s]*)`).FindStringSubmatch(url)[1]
	// 先获取视频ID
	res, err := grequests.Get(realUrl, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	})
	if err != nil {
		return "非法请求"
	}
	requrl := res.RawResponse.Request.URL.RequestURI()
	if len(requrl) < 10 {
		return ""
	}
	Itemid := regexp.MustCompile(`video/(.*?)\?`).FindStringSubmatch(requrl)[1] ///share/video/6734643996347485448/?region=CN&mid=6734637731277851404&u_code=0&titleType=title&utm_source=copy_link&utm_campaign=client_share&utm_medium=android&app=aweme
	// reg := regexp.MustCompile(".*?video/(.*?)/")
	// Itemid := reg.FindString(res.RawResponse.Request.URL.RequestURI())	--》 /share/video/6734643996347485448/
	// 根据ID然后获取原始视频
	infourl := "https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids=" + Itemid
	res2, err := grequests.Get(infourl, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	})
	if err != nil {
		return "非法请求"
	}

	// fmt.Println(res2.String())
	// reg := regexp.MustCompile("video_id=(.*?)&")
	video_id := regexp.MustCompile(`vid":"(.*?)"`).FindStringSubmatch(res2.String())[1] // 匹配两种结果，一种前后参数都代，另外一种只带（）里的
	// 拼接最终地址
	return strings.Join([]string{"https://aweme.snssdk.com/aweme/v1/play/?video_id=", video_id, "&ratio=720p&line=0"}, "") // json返回时&符号被转换为unicode，暂无解
}
