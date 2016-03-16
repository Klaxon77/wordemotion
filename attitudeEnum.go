package emotion

type Attitude int

const (
	Positive Attitude = 1
	Negative Attitude = 2
	Neutral Attitude = 0
)

var AllAttitudes = []Attitude{Positive, Negative, Neutral}

func (attitude Attitude) String() string {
	switch attitude {
	case Positive : return "positive"
	case Negative : return "negative"
	case Neutral : return "neutral"
	default: return "unkown"
	}
}

func ToAttitudeEnum(s string) Attitude {
	switch s {
	case Positive.String() : return Positive
	case Negative.String() : return Negative
	case Neutral.String() : return Neutral
	default: return -1
	}
}