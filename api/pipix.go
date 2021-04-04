package api

import (
	"regexp"
	"strings"

	"github.com/levigross/grequests"
)

func PiPix(url string) string {
	defer func() string { // 用来处理异常
		if err := recover(); err != nil { // 此处防止错误列表导致程序退出
			return ""
		}
		return ""
	}()
	res, err := grequests.Get(url, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	})
	if err != nil {
		return "非法请求"
	}
	requrl := res.RawResponse.Request.URL.RequestURI() //
	// https://h5.pipix.com/item/6863294377570081027?app_id=1319&app=super&timestamp=1598011674&user_id=62108092335&carrier_region=cn&region=cn&language=zh&utm_source=weixin
	if len(requrl) < 10 {
		return ""
	}

	// Itemid := strings.Split(requrl, "/")[2] ///share/video/6734643996347485448/?region=CN&mid=6734637731277851404&u_code=0&titleType=title&utm_source=copy_link&utm_campaign=client_share&utm_medium=android&app=aweme
	// fmt.Println(res.RawResponse.Request.URL.RequestURI())
	reg := regexp.MustCompile(`item/(.*?)\?`) // ?号需要替
	Itemid := reg.FindStringSubmatch(res.RawResponse.Request.URL.RequestURI())[1]
	// fmt.Println(Itemid)
	// 根据ID然后获取原始视频

	infourl := "https://is.snssdk.com/bds/cell/detail/?cell_type=1&aid=1319&app_name=super&cell_id=" + Itemid
	res2, err := grequests.Get(infourl, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	})
	if err != nil {
		return "非法请求"
	}
	reg2 := regexp.MustCompile("origin_video_download.*?url_list.*?url.*?:\"(.*?)\"") // re.findall("origin_video_download.*?url_list.*?url.*?:\"(.*?)\"",r.text)
	video_url := reg2.FindStringSubmatch(res2.String())[1]                            // --》 /share/video/6734643996347485448/
	// fmt.Println(rr)
	return strings.ReplaceAll(video_url, `\u0026`, "&")
}
