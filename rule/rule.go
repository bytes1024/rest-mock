package rule

const (
	NameSpace Mode = iota
	Project
	URL
)

type Mode int32

func (m Mode) String() string {
	switch m {
	case NameSpace:
		return "NS"
	case Project:
		return "PJ"
	case URL:
		return "RL"
	default:
		return ""
	}
}

//路由地址对应的规则信息
type Rule struct {
	Name string
	Mode Mode
}
