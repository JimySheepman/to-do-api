package producer

import "testing"

func TestSend(t *testing.T) {

	var tests = []struct {
		testName string
		key      string
		value    interface{}
		want     error
	}{
		{"successful send message", "producer", map[string]string{"Title": "producer Test"}, nil},
		{"successful send message without value", "producer", "", nil},
		{"successful send message without key and value", "", "", nil},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			expected := Send(test.key, test.value)

			if expected != test.want {
				t.Errorf("got:%v  want:%v ", expected, test.want)
			}
		})
	}

}
