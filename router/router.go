package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//初始设置v1版本接口
func InitRouteVersion1(group gin.RouterGroup) {
	v1:=group.Group("v1")
	{
		v1.GET("/getPath",getPath)
		v1.GET("push",push)
	}
}

//获取集合
func getPath(c *gin.Context)  {
	v,_:=c.GetQuery("k")
	log.Info("获取参数:",v)
	c.PureJSON(200,fmt.Sprintf("getPath:%s",v))
}

func push(c *gin.Context) {
	//if pusher := c.Writer.Pusher(); pusher != nil {
	//	// 使用 pusher.Push() 做服务器推送
	//	if err := pusher.Push("/assets/app.js", nil); err != nil {
	//		log.Printf("Failed to push: %v", err)
	//	}
	//}
	//c.HTML(200, "https", gin.H{
	//	"status": "success",
	//})
}


