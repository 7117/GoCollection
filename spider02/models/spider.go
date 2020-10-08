package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"strings"
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

var (
	db orm.Ormer
)

func init() {
	orm.Debug = true
	orm.RegisterModel(new(MovieInfo))
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/spider?charset=utf8", 30)
	db = orm.NewOrm()
}

func AddMovieInfo(movieInfo *MovieInfo) (id int64, err error) {
	movieInfo.Id = 0
	id, err = db.Insert(movieInfo)
	return id, err
}

func GetMovieDirector(movieHtml string) string {
	if movieHtml == "" {
		return " "
	}

	// <a href="/celebrity/1274534/" rel="v:directedBy">曾国祥</a>
	// <a .*?rel="v:directedBy">(.*)</a>

	reg := regexp.MustCompile(`<a .*?rel="v:directedBy">(.*)</a>`)

	res := reg.FindAllStringSubmatch(movieHtml, -1)

	if res == nil {
		return ""
	} else {
		return string(res[0][1])
	}
}

func GetMovieName(movieHtml string) string {
	//<span property="v:itemreviewed">七月与安生</span>

	if movieHtml == "" {
		return " "
	}

	reg := regexp.MustCompile(`<span property="v:itemreviewed">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if result == nil {
		return " "
	} else {
		return string(result[0][1])
	}

}

func GetMovieMainCharacters(movieHtml string) string {
	//<a href="/celebrity/1274224/" rel="v:starring">周冬雨</a>
	//MustCompile匹配规则
	//FindAllStringSubmatch
	reg := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	mainCharacters := ""
	for _, v := range result {
		//[<a href="/celebrity/1274224/" rel="v:starring">周冬雨</a> 周冬雨]
		mainCharacters += v[1] + "/"
	}

	return strings.Trim(mainCharacters, "/")

}

func GetMovieGrade(movieHtml string) string {
	reg := regexp.MustCompile(`<strong.*?property="v:average">(.*?)</strong>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if result == nil {
		return " "
	} else {
		return string(result[0][1])
	}
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

	if result == nil {
		return " "
	} else {
		return string(result[0][1])
	}
}

func GetMovieRunningTime(movieHtml string) string {

	reg := regexp.MustCompile(`<span.*?property="v:runtime".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if result == nil {
		return " "
	} else {
		return string(result[0][1])
	}

}

func GetMovieUrls(movieHtml string) []string {
	//<a href="https://movie.douban.com/subject/3319755/">

	reg := regexp.MustCompile(`<a.*?href="(https://movie.douban.com/subject/.*?)"`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	var movieSets []string
	for _, v := range result {
		movieSets = append(movieSets, v[1])
	}

	return movieSets

}
