package main

import (
	"server/api"
	"server/router"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(api.Cors())
	router.SetRouter(r)
	r.Run(":8090")
}
