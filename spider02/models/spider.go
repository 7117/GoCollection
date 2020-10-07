package models

import "github.com/astaxie/beego/orm"

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
	orm.RegisterModel(new(Access))
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test1?charset=utf8", 30)
	db = orm.NewOrm()
}

func AddMovie(movie_info *MovieInfo)(int64,error)  {
	id,err := db.Insert(movie_info)
	return id,err
}

