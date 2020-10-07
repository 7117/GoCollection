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

	sUrl := "https://movie.douban.com/subject/25827935/"
	sHtml := httplib.Get(sUrl)
	sHtmls, _ := sHtml.String()

	models.ConnectRedis("127.0.0.1:6379")
	urls := models.GetMovieUrls(sHtmls)

	for _, url := range urls {
		models.PutinQueue(url)
	}

	c.Ctx.WriteString(fmt.Sprintf("%v", urls))

	var movieInfo models.MovieInfo

	movieInfo.Movie_name = models.GetMovieDirector(sHtmls)
	movieInfo.Movie_director = models.GetMovieName(sHtmls)
	movieInfo.Movie_main_character = models.GetMovieMainCharacters(sHtmls)
	movieInfo.Movie_grade = models.GetMovieGrade(sHtmls)
	movieInfo.Movie_type = models.GetMovieGenre(sHtmls)
	movieInfo.Movie_on_time = models.GetMovieOnTime(sHtmls)
	movieInfo.Movie_span = models.GetMovieRunningTime(sHtmls)

	id, err := models.AddMovie(&movieInfo)

	if err != nil {
		panic(err)
	}

	c.Ctx.WriteString(fmt.Sprintf("%v", id))

}
