package main

import (
	"Sql/Sql"
	"log"

	"github.com/gin-gonic/gin"
)

func TestGet(id int, name string, content string) {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"编号":   id,
			"用户名":  name,
			"评论内容": content,
		})
	})
	r.POST("/post", func(c *gin.Context) { //post在浏览器无法正常获取，因为浏览器默认get请求
		uname := c.DefaultPostForm("uname", "默认用户")
		ucontent := c.PostForm("ucontent")
		c.JSON(200, gin.H{
			"用户名":  uname,
			"评论内容": ucontent,
		})
	})
	if err := r.Run(":8090"); err != nil { //ctrl+c可终止服务
		log.Fatal(err.Error())
	}
}

func TestPost() {
	r := gin.Default()
	r.POST("/post", func(c *gin.Context) { //post在浏览器无法正常获取，因为浏览器默认get请求
		uname := c.DefaultQuery("uname", "默认用户")
		ucontent := c.Query("ucontent")
		c.JSON(200, gin.H{
			"用户名":  uname,
			"评论内容": ucontent,
		})
	})
	if err := r.Run(":8090"); err != nil { //ctrl+c可终止服务
		log.Fatal(err.Error())
	}

}

func main() {
	Sql.SqlOpen()
	uid, uname, ucontent := Sql.SqlSelect()
	Sql.SqlClose()
	//fmt.Println(uid, uname, ucontent)
	TestGet(uid, uname, ucontent)

	// TestPost()

	// r.Run() // listen and serve on 0.0.0.0:8080
}
