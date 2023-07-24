package router

import (
	"server/controller"

	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	// 单词路由
	word := r.Group("/word")
	// 获取单词
	word.GET("/getword", controller.GetWord)
	// 获取单词组列表
	word.GET("/list", controller.List)
	// 后台路由
	admin := r.Group("/admin")
	// 获取列表详细信息
	admin.GET("/detail", controller.Detail)
	// 添加单词表
	admin.GET("/addTable", controller.AddTable)
	// 搜索
	admin.GET("/search", controller.Search)
	// 删除表格
	admin.POST("/delTable", controller.DeleteTable)
	// 更新表格数据
	admin.POST("/update", controller.Update)
	// 删除单词表中的特定单词
	admin.GET("delWord", controller.DeleteWord)
	// 添加特定的单词
	admin.POST("addWord", controller.AddWord)
	// 搜索特定的单词
	admin.GET("findWord", controller.FindWord)
}
