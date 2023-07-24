package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Static("/", "./dist")
	r.Run(":5173")
}
