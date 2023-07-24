package main

import (
	"fmt"
	"server/api"
	"server/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./dist")
	r.Use(api.Cors())
	fmt.Println("访问链接：http://localhost:8090/page")
	router.SetRouter(r)
	r.Run(":8090")
}
