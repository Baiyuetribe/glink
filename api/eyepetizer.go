package api

import (
	"regexp"
)

func EyePetizer(url string) string {
	defer func() string { // 用来处理异常
		if err := recover(); err != nil { // 此处防止错误列表导致程序退出
			return ""
		}
		return ""
	}()
	// video_id=209323 || resource_id=244837
	Itemid := regexp.MustCompile(`[vid,resource_id]=(.*)\d`).FindStringSubmatch(url)[1]
	return "https://baobab.kaiyanapp.com/api/v1/playUrl?vid=" + Itemid + "&resourceType=video&editionType=default&source=aliyun&playUrlType=url_oss&ptl=true"
}
