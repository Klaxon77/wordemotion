package emotion

import (
	"fmt"
	"strings"
)

var (
	emotionMap map[string]int
	attitudeMap map[string]Attitude
)

func init() {
	emotionMap = make(map[string]int)
	attitudeMap = make(map[string]Attitude)

	lines := strings.Split(db(), "\n")
	for _, line := range lines {
		var word, emotionOrAttitudeStr, associated string
		fmt.Sscanf(line, "%s %s %s", &word, &emotionOrAttitudeStr, &associated)

		if associated == "1" {
			emotion := ToEmotion(emotionOrAttitudeStr)
			if emotion != -1 {
				val := emotionMap[word]
				emotionMap[word] = (val | int(emotion))
			}

			attitude := ToAttitudeEnum(emotionOrAttitudeStr)
			if attitude != -1 {
				attitudeMap[word] = attitude
			}
		}
	}
}

//return emotions aggregated in int, 0 otherwise
func EmotionsIntFor(s string) int {
	return emotionMap[strings.ToLower(s)]
}

func EmotionsFor(s string) []Emotion {
	emotionsInt := EmotionsIntFor(s)

	emotions := make([]Emotion, 0, len(AllEmotions))
	for _, emo := range AllEmotions {
		if emotionsInt & int(emo) != 0 {
			emotions = append(emotions, emo)
		}
	}

	return emotions
}

func AttitudeFor(s string) Attitude {
	a, ok := attitudeMap[strings.ToLower(s)]
	if ok {
		return a
	}

	return Neutral
}