package rule

const (
	NameSpace Mode = iota
	Project
	URL
)

type Mode int32

func (m Mode) string() string {
	switch m {
	case NameSpace:
		return "namespace"
	case Project:
		return "project"
	case URL:
		return "url"
	default:
		return ""
	}
}

//路由地址对应的规则信息
type Rule struct {
	Name string
	Mode Mode
}
