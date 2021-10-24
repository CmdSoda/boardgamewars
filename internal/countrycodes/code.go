package countrycodes

type Code int

const (
	UK               Code   = 0
	Germany          Code   = 1
	InvalidParameter string = "InvalidParameter"
)

func (c Code) String() string {
	switch c {
	case UK:
		return "UK"
	case Germany:
		return "Germany"
	}
	return ""
}
