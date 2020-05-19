package controllers

import (
	"spider/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"fmt"
)

type CrawlController struct {
	beego.Controller
}

func (c *CrawlController) Crawl() {
	sUrl := "https://movie.douban.com/subject/25827935/"

	sHtml := httplib.Get(sUrl)
	sHtmls, _ := sHtml.String()

	var movieInfo models.MovieInfo

	movieInfo.Movie_name = models.GetMovieName(sHtmls)
	movieInfo.Movie_director = models.GetMovieDirector(sHtmls)
	movieInfo.Movie_main_character = models.GetMovieMainCharacters(sHtmls)
	movieInfo.Movie_type = models.GetMovieGenre(sHtmls)
	movieInfo.Movie_on_time = models.GetMovieOnTime(sHtmls)
	movieInfo.Movie_grade = models.GetMovieGrade(sHtmls)
	movieInfo.Movie_span = models.GetMovieRunningTime(sHtmls)

	id, _ := models.AddMovieInfo(&movieInfo)
	c.Ctx.WriteString(fmt.Sprintf("%v", id))

}
