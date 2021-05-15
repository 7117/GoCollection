package api11

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Create(context *gin.Context) {

	books := []User{
		{Id: 1, Name: "啊啊", Age: 12},
		{Id: 2, Name: "安安", Age: 13},
		{Id: 3, Name: "方法", Age: 14},
		{Id: 4, Name: "丰富", Age: 15},
	}

	context.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"books": books,
	})

}

func Detail(context *gin.Context) {
	id := context.Query("id")
	fmt.Println(id)

	books := []User{
		{Id: 1, Name: "啊啊", Age: 12},
		{Id: 2, Name: "安安", Age: 13},
		{Id: 3, Name: "方法", Age: 14},
		{Id: 4, Name: "丰富", Age: 15},
	}
	context.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"books": books,
	})

}
