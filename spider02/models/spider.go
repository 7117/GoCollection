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
	orm.Debug = true
	orm.RegisterModel(new(MovieInfo))
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test1?charset=utf8", 30)
	db = orm.NewOrm()
}

func AddMovie(movie_info *MovieInfo)(int64,error)  {
	id,err := db.Insert(movie_info)
	return id,err
}

func GetMovieDirector(movieHtml string)string{
	if movieHtml == "" {
		return " "
	}

	// <a href="/celebrity/1274534/" rel="v:directedBy">曾国祥</a>
	// <a .*?rel="v:directedBy">(.*)</a>

	reg := regexp.MustCompile(`<a .*?rel="v:directedBy">(.*)</a>`)

	res := reg.FindAllStringSubmatch(movieHtml,-1)

	return string(res[0][1]);
}

func GetMovieName(movieHtml string)string{
	//<span property="v:itemreviewed">七月与安生</span>

	if movieHtml == "" {
		return " "
	}

	reg := regexp.MustCompile(`<span property="v:itemreviewed">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	return string(result[0][1])
}

func GetMovieMainCharacters(movieHtml string) string {
	//<a href="/celebrity/1274224/" rel="v:starring">周冬雨</a>

	if movieHtml == "" {
		return " "
	}

	reg := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)

	result := reg.FindAllStringSubmatch(movieHtml, -1)

	mainCharacters := ""
	for _, v := range result {
		mainCharacters += v[1] + "/"
	}

	return strings.Trim(mainCharacters, "/")


}



