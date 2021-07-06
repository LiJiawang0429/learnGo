package main

import (
	Sql "Sql/MySql"
	"encoding/json"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
)

func TestGet() {
	r := gin.Default()
	
	r.GET("/get", func(c *gin.Context) { //struct\map\json可相互转化
		jsonStr := LinkSql()
		c.JSON(200, gin.H{
			"json": jsonStr,
		})
	})
	r.POST("/post", func(c *gin.Context) { //post在浏览器无法正常获取，因为浏览器默认get请求
		uname := c.DefaultPostForm("uname", "默认用户")
		ucontent := c.PostForm("ucontent")
		c.JSON(200, gin.H{
			"用户名":  uname,
			"评论内容": ucontent,
		})
		SendInfo(uname, ucontent)
		fmt.Println(uname)
	})
	if err := r.Run(":8090"); err != nil { //ctrl+c可终止服务
		log.Fatal(err.Error())
	}
}

func SendInfo(name string, content string) {
	if content != "" {
		Sql.SqlOpen()
		Sql.SqlInsert(name, content)
		fmt.Printf("%s的留言数据已入库！\n", name)
		Sql.SqlClose()
	} else {
		fmt.Println("评论内容不可为空！！")
	}
}

func LinkSql()(jsonStr string){
	Sql.SqlOpen()
	list := Sql.SqlSelect()
	jsonS, err := json.Marshal(list)

	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	jsonStr = string(jsonS)
	fmt.Println(jsonStr)
	Sql.SqlClose()
	return 
}

func main() {
	TestGet()
}
