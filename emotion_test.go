package emotion

import "testing"

func TestEmotion(t *testing.T) {
	cases := []struct {
		in  string
		emo int
	}{
		{"anger", int(Anger)},
		{"love", int(Joy)},
	}

	for _, c := range cases {
		emo, ok := EmotionFor(c.in)
		if !ok || emo != c.emo {
			t.Errorf("EmotionFor(%q) == %q, want %q", c.in, emo, c.emo)
		}
	}
}

func TestAttitude(t *testing.T) {
	cases := []struct {
		in  string
		attitude Attitude
	}{
		{"anger", Negative},
		{"love", Positive},
	}

	for _, c := range cases {
		attitude := AttitudeFor(c.in)
		if attitude != c.attitude {
			t.Errorf("AttitudeFor(%q) == %q, want %q", c.in, attitude, c.attitude)
		}
	}
}