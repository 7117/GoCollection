package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"goPractice/spider02/models"
)

type SpiderController struct {
	beego.Controller
}

func (c *SpiderController) CrawlMovie() {

	var movieInfo models.MovieInfo
	models.ConnectRedis("127.0.0.1:6379")

	sUrl := "https://movie.douban.com/subject/25827935/"
	models.PutinQueue(sUrl)

	for {
		//url数量
		length := models.GetQueueLength()
		if length == 0 {
			break
		}

		//进行获取url
		sHtml := httplib.Get(sUrl)
		sHtmls, _ := sHtml.String()
		movieInfo.Movie_name = models.GetMovieDirector(sHtmls)
		movieInfo.Movie_director = models.GetMovieName(sHtmls)
		movieInfo.Movie_main_character = models.GetMovieMainCharacters(sHtmls)
		movieInfo.Movie_grade = models.GetMovieGrade(sHtmls)
		movieInfo.Movie_type = models.GetMovieGenre(sHtmls)
		movieInfo.Movie_on_time = models.GetMovieOnTime(sHtmls)
		movieInfo.Movie_span = models.GetMovieRunningTime(sHtmls)

		id, _ := models.AddMovie(&movieInfo)

		c.Ctx.WriteString(fmt.Sprintf("%v", id))
	}

}
