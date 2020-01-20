package setting

import (
	"github.com/gin-gonic/gin"
	"rest-mock/rule"
)

//设置
type Setting struct {
	NameSpace string
	Unique    string
	Rules     []rule.Rule
}

//根据空间+唯一标识获取对应的设置信息
func GetSettingByNameSpaceAndUnique(nameSpace, unique string) Setting {
	return Setting{
		NameSpace: "",
		Unique:    "",
		Rules:     nil,
	}
}

//根据当前的规则执行返回对应的文本信息
func (s Setting) Handler(c *gin.Context) string {
	rules := s.Rules

	if rules != nil && len(rules) > 0 {
		for _, r := range rules {
		}
	}

}
