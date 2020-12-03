package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"web-youku/models"
	"web-youku/utils"
)

type MainController struct {
	beego.Controller
}

// @router /index [get]
func (c *MainController) Get() {
	fmt.Println(1)
	//前后端分离
	//获取顶部视频推荐广告，调用接口
	advertVideos := models.GetChannelAdvert(1)
	c.Data["advertVideos"] = advertVideos
	//获取正在热播，调用接口
	hotVideos := models.GetChannelHotList(1)
	c.Data["hotVideos"] = hotVideos
	//获取日漫、国漫推荐，调用接口
	japanVideos := models.GetChannelRegionRecommend(1, 1)
	chinaVideos := models.GetChannelRegionRecommend(1, 2)
	c.Data["japanVideos"] = japanVideos
	c.Data["chinaVideos"] = chinaVideos
	//获取少女推荐，调用接口
	girlVideos := models.GetChannelTypeRecommend(1, 1)
	c.Data["girlVideos"] = girlVideos
	//获取动漫排行榜
	comicTop := models.GetChannelTop(1)
	c.Data["comicTop"] = comicTop
	//获取少女排行榜
	girlTop := models.GetTypeTop(1)
	c.Data["girlTop"] = girlTop
	c.TplName = "index.html"
}

//更多页面
//频道页
func (c *MainController) Channel() {
	//接受浏览器参数
	//地区
	regionId, _ := c.GetInt("regionId")
	c.Data["regionId"] = regionId
	//类型
	typeId, _ := c.GetInt("typeId")
	c.Data["typeId"] = typeId
	//状态
	end := c.GetString("end")
	c.Data["end"] = end
	//排序
	sort := c.GetString("sort")
	if sort == "" {
		sort = "episodesUpdateTime"
	}
	c.Data["sort"] = sort
	//获得地区列表
	regions := models.GetChannelRegion(1)
	c.Data["regions"] = regions
	//获得类型列表
	types := models.GetChannelType(1)
	c.Data["types"] = types

	//根据接受到的参数获取视频信息
	videos, _ := models.GetChannelVideoList(1, regionId, typeId, end, sort, 1)
	c.Data["videos"] = videos
	c.TplName = "channel.html"
}

//频道JS加载数据接口
func (c *MainController) ChannelVideoData() {
	//接受浏览器参数
	//地区
	regionId, _ := c.GetInt("regionId")
	//类型
	typeId, _ := c.GetInt("typeId")
	//状态
	end := c.GetString("end")
	//排序
	sort := c.GetString("sort")
	if sort == "" {
		sort = "episodesUpdateTime"
	}
	//页数
	page, _ := c.GetInt("page")
	//根据接受到的参数获取视频信息
	videos, count := models.GetChannelVideoList(1, regionId, typeId, end, sort, page)
	title := utils.ReturnSuccess(0, "success", videos, count)
	c.Ctx.WriteString(title)
}

//搜索页
func (c *MainController) Search() {
	keyword := c.GetString("keyword")
	c.Data["keyword"] = keyword
	c.TplName = "search.html"
}

//搜索接口
func (c *MainController) SearchData() {
	//接受浏览器参数
	//搜索关键词
	keyword := c.GetString("keyword")
	//页数
	page, _ := c.GetInt("page")
	//根据接受到的参数获取视频信息
	videos, count := models.GetSearchVideoList(keyword, page)
	title := utils.ReturnSuccess(0, "success", videos, count)
	c.Ctx.WriteString(title)
}

//排行榜
func (c *MainController) Top() {
	//获取动漫排行榜
	comicTop := models.GetChannelTop(1)
	c.Data["comicTop"] = comicTop
	//获取少女排行榜
	girlTop := models.GetTypeTop(1)
	c.Data["girlTop"] = girlTop
	c.TplName = "top.html"
}

//详情页
func (c *MainController) Show() {
	//获取视频ID
	videoId, _ := c.GetInt("id")
	c.Data["videoId"] = videoId
	episodesId, _ := c.GetInt("episodesId")
	//获取视频信息
	videoInfo := models.GetVideoInfo(videoId)
	c.Data["videoInfo"] = videoInfo
	//获取视频集数列表
	episodesList := models.GetVideoEpisodesList(videoId)
	c.Data["episodesList"] = episodesList
	var episodes models.Episodes
	if len(episodesList) > 0 {
		for i, v := range episodesList {
			if episodesId != 0 && episodesId == v.Id {
				episodes.Id = v.Id
				episodes.Title = v.Title
				episodes.AddTime = v.AddTime
				episodes.Num = v.Num
				episodes.PlayUrl = v.PlayUrl
				episodes.Comment = v.Comment
				episodes.AliyunVideoId = v.AliyunVideoId
			} else if episodesId == 0 && i == 0 {
				episodes.Id = v.Id
				episodes.Title = v.Title
				episodes.AddTime = v.AddTime
				episodes.Num = v.Num
				episodes.PlayUrl = v.PlayUrl
				episodes.Comment = v.Comment
				episodes.AliyunVideoId = v.AliyunVideoId
			}
		}
	}
	c.Data["episodes"] = episodes

	//获取播放凭证
	if episodes.AliyunVideoId != "" {
		c.Data["PlayAuth"] = models.GetPlayAuth(episodes.AliyunVideoId)
	} else {
		c.Data["PlayAuth"] = ""
	}
	//获取国漫推荐
	chinaVideos := models.GetChannelRegionRecommend(1, 2)
	c.Data["chinaVideos"] = chinaVideos
	//获取日漫推荐
	japanVideos := models.GetChannelRegionRecommend(1, 1)
	c.Data["japanVideos"] = japanVideos
	//获取少女推荐
	girlVideos := models.GetChannelTypeRecommend(1, 1)
	c.Data["girlVideos"] = girlVideos

	c.TplName = "show.html"
}

//获取评论信息
func (c *MainController) GetCommentList() {
	//集数ID
	episodesId, _ := c.GetInt("episodesId")
	//获取页码
	page, _ := c.GetInt("page")
	//获取评论ID
	title := models.GetCommentList(episodesId, page)
	c.Ctx.WriteString(title)
}

//保存评论
func (c *MainController) SaveComment() {
	//集数ID
	episodesId, _ := c.GetInt("episodesId")
	videoId, _ := c.GetInt("videoId")
	//获取用户ID
	uid, _ := c.GetInt("uid")
	//内容
	content := c.GetString("content")
	//获取评论ID
	title := models.SaveComment(content, uid, episodesId, videoId)
	c.Ctx.WriteString(title)
}

//保存弹幕
func (c *MainController) SaveBarrage() {
	//集数ID
	episodesId, _ := c.GetInt("episodesId")
	videoId, _ := c.GetInt("videoId")
	//获取用户ID
	uid, _ := c.GetInt("uid")
	//内容
	content := c.GetString("content")
	//视频当前播放时间
	currentTime, _ := c.GetInt("currentTime")
	//获取评论ID
	title := models.SaveBarrage(content, currentTime, uid, episodesId, videoId)
	c.Ctx.WriteString(title)
}
