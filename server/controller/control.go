package controller

import (
	"encoding/json"
	"fmt"
	"os"
	"server/api"
	"server/model"
	"strings"

	"github.com/gin-gonic/gin"
)

// 获取单词
func GetWord(c *gin.Context) {
	str := c.Query("item")
	list := api.Read(str)
	c.JSON(200, list)
}

// 获取单词列表
func List(c *gin.Context) {
	arr := api.GetList()
	c.JSON(200, arr)
}

// 获取列表的详细信息
func Detail(c *gin.Context) {
	var details []model.Detail = api.GetDetail()
	c.JSON(200, details)
}

// 添加单词表
func AddTable(c *gin.Context) {
	name := c.Query("name")
	file, err := os.Create("./json/" + name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file)
		c.JSON(200, gin.H{
			"flag": true,
		})
	}
}

// 搜索表
func Search(c *gin.Context) {
	var name string = c.Query("name")
	var arr []model.Detail = api.GetDetail()
	var result []model.Detail
	for _, v := range arr {
		if strings.Contains(v.Name, name) {
			result = append(result, v)
		}

	}
	c.JSON(200, result)
}

// 删除表
func DeleteTable(c *gin.Context) {
	// 获取数据
	str := c.PostForm("item")
	var detail model.Detail
	json.Unmarshal([]byte(str), &detail)
	// 操作
	err := os.Remove("./json/" + detail.Name)
	if err != nil {
		c.String(200, "删除失败")
	} else {
		c.String(200, "删除成功")
	}

}

// 更新表内容
func Update(c *gin.Context) {
	name := c.PostForm("name")
	strWords := c.PostForm("words")
	var words []model.Word
	json.Unmarshal([]byte(strWords), &words)
	api.Write(words, name)
}

// 删除特定的单词
func DeleteWord(c *gin.Context) {
	name := c.Query("name")
	strWord := c.Query("word")
	var word model.Word
	list := api.Read(name)
	json.Unmarshal([]byte(strWord), &word)
	new := model.DeleteItem(list, word)
	api.Write(new, name)
}

// 添加特定的单词
func AddWord(c *gin.Context) {
	name := c.PostForm("name")
	strList := c.PostForm("list")
	var list []model.Word
	json.Unmarshal([]byte(strList), &list)
	api.Write(list, name)
}

// 查找特定的单词
func FindWord(c *gin.Context) {
	item := c.Query("item")
	table := c.Query("table")
	list := api.Read(table)
	var resp []model.Word
	for _, v := range list {
		if strings.Contains(v.Eng, item) {
			resp = append(resp, v)
		}
	}
	fmt.Println(resp)
	c.JSON(200, resp)
}
