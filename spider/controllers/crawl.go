package controllers

import (
	"spider/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"time"
)

type CrawlController struct {
	beego.Controller
}

func (c *CrawlController) Crawl() {
	var movieInfo models.MovieInfo
	models.ConnectRedis("127.0.0.1:6379")

	//init
	sUrl := "https://movie.douban.com/subject/25827935/"
	models.PutinQueue(sUrl)

	for {

		length := models.GetQueueLength()
		if length == 0 {
			break
		}

		sUrl = models.PopfromQueue()

		if models.IsVisit(sUrl) {
			continue
		}

		sHtml := httplib.Get(sUrl)
		sHtmls, _ := sHtml.String()

		movieInfo.Movie_name = models.GetMovieName(sHtmls)
		if movieInfo.Movie_name != "" {
			movieInfo.Movie_director = models.GetMovieDirector(sHtmls)
			movieInfo.Movie_main_character = models.GetMovieMainCharacters(sHtmls)
			movieInfo.Movie_type = models.GetMovieGenre(sHtmls)
			movieInfo.Movie_on_time = models.GetMovieOnTime(sHtmls)
			movieInfo.Movie_grade = models.GetMovieGrade(sHtmls)
			movieInfo.Movie_span = models.GetMovieRunningTime(sHtmls)

			models.AddMovieInfo(&movieInfo)
		}

		//进行添加url
		urls := models.GetMovieUrls(sHtmls)
		for _, url := range urls {
			models.PutinQueue(url)
			c.Ctx.WriteString(url + "   ")
		}

		models.AddToSet(sUrl)
		time.Sleep(3 * time.Second)
	}

}
