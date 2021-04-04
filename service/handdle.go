package service

import (
	"Baiyuetribe/glink/api"
	"fmt"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ApiHandler(c *fiber.Ctx) error {
	url := c.Params("*") // 带有?的替换为&后进行，否则url无参数
	match, _ := regexp.MatchString("^([a-zA-Z]*)://([^ ]*)$", url)
	if !match {
		return c.JSON(fiber.Map{"msg": "请输入正确的url地址"})
	}
	if !strings.HasPrefix(url, "http") {
		return c.JSON(fiber.Map{"msg": "请输入合法的url地址"})
	}
	// 格局url匹配函数
	var r string
	if strings.Contains(url, "douyin") { // 没有匹配时，值为-1
		r = api.DouYin(url)
	} else if strings.Contains(url, "weishi") {
		r = api.WeiShi(url)
	} else if strings.Contains(url, "pipix") {
		r = api.PiPix(url)
	} else if strings.Contains(url, "kuaishou") {
		r = api.KuaiShou(url)
	} else if strings.Contains(url, "kg3.qq.com") {
		r = api.Kg3(url)
	} else if strings.Contains(url, "ixigua.com") {
		r = "截至2021-04-04，仅能通过下载合并的方式获取无无水印视频，可付费实现"
	} else if strings.Contains(url, "eyepetizer") {
		r = api.EyePetizer(url)
	} else {
		r = "暂未支持该接口，请提交issueue"
	}
	return c.SendString(r)
	// return c.JSON(&fiber.Map{"msg": r})

}

func PrintLogo() {
	fmt.Print(`

	██████╗  █████╗ ██╗██╗   ██╗██╗   ██╗███████╗
	██╔══██╗██╔══██╗██║╚██╗ ██╔╝██║   ██║██╔════╝
	██████╔╝███████║██║ ╚████╔╝ ██║   ██║█████╗  
	██╔══██╗██╔══██║██║  ╚██╔╝  ██║   ██║██╔══╝  
	██████╔╝██║  ██║██║   ██║   ╚██████╔╝███████╗
	╚═════╝ ╚═╝  ╚═╝╚═╝   ╚═╝    ╚═════╝ ╚══════╝
			欢迎使用Glink短视频去水印软件
			作者QQ：2894049053

`)
}
