package main

import (
	"gin_project/config"
	_ "gin_project/data_source"
	_ "gin_project/logs_source"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	config.Router(router)

	router.Run(":8090")
}
