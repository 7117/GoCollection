package first

import (
	"fmt"
	_ "gin_project/dblink"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id   int
	Name string
	Age  int
	Addr string
}

type Getdata struct {
	Name string
	Age  int
}

type UserProfile struct {
	Id     int
	Pic    string
	CPic   string
	Phone  string
	UserID int // 默认关联字段为Id
	Name   string
	Age    int
	Addr   string
}

func Hello(context *gin.Context) {

	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	defer db.Close() // 关闭连接

	if nil != err {
		fmt.Println(err)
	}

	first(db)
	//firstOrCreatedb(db)
	//last(db)
	//take(db)
	//find(db)
	//where(db)
	//selectdata(db)
	//createdata(db)
	//batchcreatedata(db)
	//savedata(db)
	//updatedata(db)
	//del(db)
	//Unscoped(db)
	//notdata(db)
	//ordata(db)
	//orderdata(db)
	//page(db)
	//scandata(db)
	//countdata(db)
	//groupdata(db)
	//Joindata(db)
	//Firstinit(db)
	//Pluck(db)

	//var useaaars []User
	//db.Debug().Scopes(Getsatus).Find(&useaaars)
	//fmt.Println(useaaars)
	//var useaaars []User
	//db.Debug().Scopes(GetRowByName(11)).Find(&useaaars)
	//fmt.Println(useaaars)

	//var useaaars []User
	//db.Scopes(GetRowByName(11)).Find(&useaaars)
	//fmt.Println(useaaars)

}

func GetRowByName(age int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("age IN (?)", age)
	}
}

func Getsatus(db *gorm.DB) *gorm.DB {
	return db.Where("age IN (?)", 100)
}

func Pluck(db *gorm.DB) {
	var user []User
	var ages []int64
	var getdata Getdata
	db.Find(&user).Pluck("age", &ages)
	fmt.Println(ages)

	var names []string
	db.Model(&User{}).Pluck("name", &names)
	fmt.Println(names)

	// 超过一列的查询，应该使用 `Scan` 或者 `Find`，例如：
	db.Debug().Select("name,age").Find(&user).Scan(&getdata)
	fmt.Println(getdata)
}

func Firstinit(db *gorm.DB) {
	var user User

	//db.Debug().FirstOrInit(&user, User{Name: "hallen",Age: 19999})
	//db.Debug().FirstOrInit(&user, map[string]interface{}{"name":"hallen"})
	//fmt.Println(user)

	//db.Where(User{Name:"halsssslen"}).Attrs(User{Name:"hallen",Age:18}).FirstOrInit(&user)
	//fmt.Println(user)

	db.Where("id = ?", "1").Assign(User{Id: 15}).FirstOrInit(&user)
	fmt.Println(user)

}

func Joindata(db *gorm.DB) {

	var user User
	var userprofile UserProfile

	db.Debug().
		Select("users.*,user_profiles.*").Joins("left join user_profiles on user_profiles.user_id = users.id").
		Where("users.id=?", 1).
		Find(&user).Scan(&userprofile)

	fmt.Println(user)
	fmt.Println(userprofile)
}

func groupdata(db *gorm.DB) {
	var user User

	db.Debug().Select("name, count(*)").Group("name").Find(&user)
	// select name,count(*) FROM users GROUP BY `age`

	db.Debug().Select("name, count(*)").Group("name").Having("name = ?", "hallen").Find(&user)
	// select name,count(*) FROM users GROUP BY `age` 后面不能用where限制条件，只能使用having

}

func countdata(db *gorm.DB) {
	var users User
	var count int

	db.Where("name = ?", "hallen").Find(&users).Count(&count)
	fmt.Println(count)
	// SELECT count(*) FROM users WHERE name = 'jinzhu'

	db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
	// SELECT count(*) FROM users WHERE name = 'jinzhu'; (count)

	db.Table("deleted_users").Count(&count)
	// SELECT count(*) FROM deleted_users;
}

func scandata(db *gorm.DB) {

	var user []User

	type Result struct {
		Id   int64
		Name string
	}

	var results []Result
	db.Select("id,name").Where("id in (?)", []int{1, 2}).Find(&user).Scan(&results)
}

func page(db *gorm.DB) {
	var users []User

	db.Debug().Limit(3).Find(&users) // 三条
	// SELECT * FROM users LIMIT 3;

	db.Debug().Limit(10).Offset(5).Find(&users) // 从5开始的10条数据
	// SELECT * FROM users OFFSET 5 LIMIT 10;
}

func orderdata(db *gorm.DB) {
	var users []User
	db.Order("age desc").Find(&users) // 注意这里的order要在find前面，否则不生效
	fmt.Println(users)

	// SELECT * FROM users ORDER BY age desc;
}

func ordata(db *gorm.DB) {
	var users []User

	db.Debug().Where("name = 'hallen'").Or(User{Name: "hallen2", Age: 18}).Find(&users)
}

func notdata(db *gorm.DB) {
	var user User

	db.Debug().Not(User{Name: "hallen", Age: 18}).First(&user)
}

func Unscoped(db *gorm.DB) {

	var user User
	var users []User

	// 在查询时会忽略被软删除的记录
	db.Where("age = 20").Find(&user)
	// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;

	// 查询逻辑删除的数据   包含所有的
	db.Unscoped().Where("age = 20").Find(&users)
	// SELECT * FROM users WHERE age = 20;

	// 想要物理删除的办法   忽略软删除  进行彻底的删除
	db.Unscoped().Delete(&user)
}

func del(db *gorm.DB) {
	var user User

	db.Delete(&user, 1)
	db.Where("age = ?", 20).Delete(&user)
}

func updatedata(db *gorm.DB) {
	var users []User
	db.Where("active = ?", true).Find(&users).Update("name", "hello")

	db.Where("active = ?", true).Find(&users).Updates(User{Name: "hello", Age: 18})
}

func savedata(db *gorm.DB) {
	var user User

	db.First(&user)

	user.Name = "jinzhu 2"
	user.Age = 100
	db.Save(&user)
}

func batchcreatedata(db *gorm.DB) {
	user4 := []User{
		{
			Name: "hallen8",
			Age:  18,
			Addr: "xxx",
		},
		{
			Name: "hallen9",
			Age:  18,
			Addr: "xxx",
		},
	}

	db.Create(&user4) // 这种方式不支持
}

func createdata(db *gorm.DB) {
	user := User{Name: "李四", Age: 18, Addr: "xxx"}
	result := db.Create(&user)

	fmt.Println(user.Id)             // 返回插入数据的主键
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数
}

func selectdata(db *gorm.DB) {
	var users User
	db.Debug().Select("name, age").Find(&users)
	// SELECT name, age FROM users;

	db.Debug().Select([]string{"name", "age"}).Find(&users)
	// SELECT name, age FROM users;

	db.Debug().Table("users").Select("COALESCE(age,?)", 42).Rows() // COALESCE:聚合
	// SELECT COALESCE(age,'42') FROM users;
}

func where(db *gorm.DB) {
	var user User

	// 根据条件查询得到满足条件的第一条记录
	db.Where("role_id = ?", "2").First(&user)
	fmt.Println(user)

	var users []User

	// 根据条件查询得到满足条件的所有记录
	db.Where("user_id = ?", "1").Find(&users)
	fmt.Println(users)

	// like模糊查询
	db.Where("role_id like ?", "%2").Find(&users)
	fmt.Println(users)

	db.Where("updated_at > ?", "2019-02-08 18:08:27").Find(&users)
	fmt.Println(users)
}

func find(db *gorm.DB) {
	var user User

	// 所有记录
	db.Find(&user, []int{1, 2, 3})
	// sql语句：// SELECT * FROM users WHERE id IN (1,2,3);

	db.Find(&user)
	// sql语句：SELECT * FROM users;

	// 根据指定条件查询
	db.Find(&user, "name = ?", "hallen")

	//或者结合where
	db.Where("name = ?", "hallen").Find(&user)
	// sql语句：SELECT * FROM users WHERE name = 'hallen';

	db.Where("name LIKE ?", "%ha%").Find(&user)
	// sql语句：SELECT * FROM users WHERE name LIKE '%hal%';
}

func take(db *gorm.DB) {
	var user User
	db.Take(&user)
	fmt.Println(user)
}

func last(db *gorm.DB) {
	// 获取最后一条记录（主键降序）
	var user User
	// 按照主键顺序的最后一条记录
	db.Last(&user)
	fmt.Println(user)
}

func first(Db *gorm.DB) {
	var user User

	result := Db.First(&user)
	fmt.Println(user)

	fmt.Println(result.RowsAffected) // 返回找到的记录数
	fmt.Println(result.Error)        // returns error
}

func firstOrCreatedb(db *gorm.DB) {
	// 未找到 user，则根据给定条件创建一条新纪录
	var user User
	db.FirstOrCreate(&user, User{Name: "hallen"})

}
