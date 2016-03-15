package emotion

type Emotion int

const (
	Anger Emotion = 1
	Anticipation Emotion = 2
	Disgust Emotion = 4
	Fear Emotion = 8
	Joy Emotion = 16
	Sadness Emotion = 32
	Surprise Emotion = 64
	Trust Emotion = 128
)

func (emotion Emotion) String() string {
	switch emotion {
	case Anger : return "anger"
	case Anticipation : return "anticipation"
	case Disgust : return "disgust"
	case Fear : return "fear"
	case Joy : return "joy"
	case Sadness : return "sadness"
	case Surprise : return "surprise"
	case Trust : return "trust"
	default: return "unkown"
	}
}

func ToEmotion(s string) Emotion {
	switch s {
	case Anger.String() : return Anger
	case Anticipation.String() : return Anticipation
	case Disgust.String() : return Disgust
	case Fear.String() : return Fear
	case Joy.String() : return Joy
	case Sadness.String() : return Sadness
	case Surprise.String() : return Surprise
	case Trust.String() : return Trust
	default: return -1
	}
}