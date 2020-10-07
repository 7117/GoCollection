package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type SpiderController struct {
	beego.Controller
}

func (c *SpiderController) CrawlMovie() {

	sUrl := "https://movie.douban.com/subject/25827935/"
	sHtml := httplib.Get(sUrl)
	sHtmls, _ := sHtml.String()

	c.Ctx.WriteString(sHtmls)

}
