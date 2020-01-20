package router

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"rest-mock/setting"
	"strings"
)

//访问地址默认定义
type PathDefine struct {
	Path    string
	Method  string
	Setting setting.Setting
}

//mock 两个地址
func getDefines() []PathDefine {
	return []PathDefine{
		{
			Path:   "v1/test1",
			Method: "get",
		},
		{
			Path:   "v1/test1",
			Method: "post",
		},
	}
}

func (p PathDefine) isGetMethod() bool {
	return strings.EqualFold(p.Method, "get")
}

func (p PathDefine) isPostMethod() bool {
	return strings.EqualFold(p.Method, "post")
}

func handler(p PathDefine) func(g *gin.Context) {
	return func(g *gin.Context) {
		//根据规则获取值
		s := p.Setting
		if s != nil {
			g.PureJSON(200, s.Handler(g))
			return
		}
		g.PureJSON(200, "not found rule")
	}
}

//初始化路由信息
//如果路由地址信息重复,会有异常情况
//TODO
func InitRouteVersion1(r gin.RouterGroup) {
	defines := getDefines()
	for _, define := range defines {
		log.Info("register define: ", define)
		if define.isGetMethod() {
			r.GET(define.Path, handler(define))
		} else if define.isPostMethod() {
			r.POST(define.Path, handler(define))
		}
	}
}
