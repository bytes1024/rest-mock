package rule

import (
	"fmt"
	"testing"
)

//测试规则构建
func TestRuleBuilder(t *testing.T) {
	rule := Rule{
		Name: "testRule",
		Mode: 0,
	}
	fmt.Printf("%v \n", rule.Mode)
}
