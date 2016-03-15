package emotion

import (
	"fmt"
	"strings"
)

var emotionMap map[string]int
var attitudeMap map[string]Attitude
func init() {
	emotionMap = make(map[string]int)
	attitudeMap = make(map[string]Attitude)

	lines := strings.Split(db(), "\n")
	for _, line := range lines {
		var word, emotionOrAttitudeStr, associated string
		fmt.Sscanf(line, "%s %s %s", &word, &emotionOrAttitudeStr, &associated)

		if (associated == "1") {
			emotion := ToEmotion(emotionOrAttitudeStr)
			if emotion != -1 {
				val := emotionMap[word]
				emotionMap[word] = (val | int(emotion))
			}

			attitude := ToAttitude(emotionOrAttitudeStr)
			if attitude != -1 {
				attitudeMap[word] = attitude
			}
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func EmotionFor(s string) (int, bool) {
	r, ok := emotionMap[strings.ToLower(s)]
	return r, ok
}

func AttitudeFor(s string) Attitude {
	a, ok := attitudeMap[strings.ToLower(s)]
	if ok {
		return a
	}

	return Neutral
}