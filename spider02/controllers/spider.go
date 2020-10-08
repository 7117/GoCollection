package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"goPractice/spider02/models"
	"time"
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
		//url数量个数
		length := models.GetQueueLength()
		if length == 0 {
			break
		}

		/**
		只有没有访问过的才进行提取
		*/
		sUrl = models.PopfromQueue()
		c.Ctx.WriteString(sUrl)
		if models.IsVisit(sUrl) {
			continue
		}

		/**
		网页内容的处理后进行保存DB
		*/
		sHtml := httplib.Get(sUrl)
		sHtmls, err := sHtml.String()
		if err != nil {
			panic(err)
		}

		movieInfo.Movie_name = models.GetMovieName(sHtmls)

		if movieInfo.Movie_name != " " {

			c.Ctx.WriteString(movieInfo.Movie_name)
			movieInfo.Movie_director = models.GetMovieDirector(sHtmls)
			movieInfo.Movie_main_character = models.GetMovieMainCharacters(sHtmls)
			movieInfo.Movie_grade = models.GetMovieGrade(sHtmls)
			movieInfo.Movie_type = models.GetMovieGenre(sHtmls)
			movieInfo.Movie_on_time = models.GetMovieOnTime(sHtmls)
			movieInfo.Movie_span = models.GetMovieRunningTime(sHtmls)

			models.AddMovieInfo(&movieInfo)
		}

		/**
		进行添加至已经访问过的url
		*/
		models.AddToSet(sUrl)

		/**
		添加保存url队列中
		*/
		urls := models.GetMovieUrls(sHtmls)
		for _, url := range urls {
			models.PutinQueue(url)
		}

		time.Sleep(2 * time.Second)

	}

}
