package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"goPractice/spider02/models"
)

type SpiderController struct {
	beego.Controller
}

func (c *SpiderController) CrawlMovie() {

	sUrl := "https://movie.douban.com/subject/25827935/"
	sHtml := httplib.Get(sUrl)
	sHtmls, _ := sHtml.String()


	c.Ctx.WriteString(models.GetMovieDirector(sHtmls))
	c.Ctx.WriteString(models.GetMovieName(sHtmls)+ "|")
	c.Ctx.WriteString(models.GetMovieMainCharacters(sHtmls))

}






