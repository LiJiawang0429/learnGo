package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "nice！",
		})
	})
	r.POST("/post", func(c *gin.Context) { //post在浏览器无法正常获取，因为浏览器默认get请求
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	if err := r.Run(":8090"); err != nil { //ctrl+c可终止服务
		log.Fatal(err.Error())
	}
	// r.Run() // listen and serve on 0.0.0.0:8080
}
