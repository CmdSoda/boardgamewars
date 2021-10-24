package countrycodes

type Code int

const (
	UK               Code   = 0
	German           Code   = 1
	InvalidParameter string = "InvalidParameter"
)

func (c Code) String() string {
	switch c {
	case UK:
		return "UK"
	case German:
		return "German"
	}
	return ""
}
