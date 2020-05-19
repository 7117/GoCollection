package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"strings"
)

var (
	db orm.Ormer
)

type MovieInfo struct {
	Id                   int64
	Movie_id             int64
	Movie_name           string
	Movie_pic            string
	Movie_director       string
	Movie_writer         string
	Movie_country        string
	Movie_language       string
	Movie_main_character string
	Movie_type           string
	Movie_on_time        string
	Movie_span           string
	Movie_grade          string
}

func init() {
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/spider?charset=utf8", 30)
	orm.RegisterModel(new(MovieInfo))
	db = orm.NewOrm()
}

func AddMovieInfo(movieInfo *MovieInfo) (id int64, err error) {
	id, err = db.Insert(movieInfo)
	return id, err
}

func GetMovieDirector(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	reg := regexp.MustCompile(`<a.*?rel="v:directedBy">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	return string(result[0][1])
}

func GetMovieName(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	reg := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	return string(result[0][1])
}

func GetMovieMainCharacters(movieHtml string) string {
	reg := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	mainCharacters := ""
	for _, v := range result {
		mainCharacters += v[1] + "/"
	}

	return strings.Trim(mainCharacters, "/")
}

func GetMovieGrade(movieHtml string) string {
	reg := regexp.MustCompile(`<strong.*?property="v:average">(.*?)</strong>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	return string(result[0][1])
}

func GetMovieGenre(movieHtml string) string {
	reg := regexp.MustCompile(`<span.*?property="v:genre">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	movieGenre := ""
	for _, v := range result {
		movieGenre += v[1] + "/"
	}
	return strings.Trim(movieGenre, "/")
}

func GetMovieOnTime(movieHtml string) string {
	reg := regexp.MustCompile(`<span.*?property="v:initialReleaseDate".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	return string(result[0][1])
}

func GetMovieRunningTime(movieHtml string) string {
	reg := regexp.MustCompile(`<span.*?property="v:runtime".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	return string(result[0][1])
}
