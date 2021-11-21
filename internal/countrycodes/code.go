package countrycodes

type Code int

const (
	UK      Code = 0
	Germany Code = 1
	USA     Code = 2
	Russia  Code = 3
)

func (c Code) String() string {
	switch c {
	case UK:
		return "UK"
	case Germany:
		return "Germany"
	case USA:
		return "USA"
	case Russia:
		return "Russia"
	}
	return ""
}
