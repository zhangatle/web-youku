package controllers

import (
	"github.com/astaxie/beego"
	"web-youku/models"
)

type AliyunController struct {
	beego.Controller
}

//获取上传凭证
func (c *AliyunController) CreateUploadVideo() {
	title := c.GetString("title")
	desc := c.GetString("desc")
	fileName := c.GetString("fileName")
	coverUrl := c.GetString("coverUrl")
	tags := c.GetString("tags")
	req := models.CreateUploadVideo(title, desc, fileName, coverUrl, tags)
	c.Ctx.WriteString(req)
}

//刷新上传凭证
func (c *AliyunController) RefreshUploadVideo() {
	videoId := c.GetString("videoId")
	req := models.RefreshUploadVideo(videoId)
	c.Ctx.WriteString(req)
}
