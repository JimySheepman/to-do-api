package consumer

import (
	"testing"
)

func TestGetKafkaReader(t *testing.T) {
	var tests = []struct {
		testName string
		url      string
		topic    string
		groupID  string
		expect   error
	}{
		{"create successful kafka reader", "127.0.0.1", "my-topic", "0", nil},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			_, err := GetKafkaReader(test.url, test.topic, test.groupID)
			if err != test.expect {
				t.Errorf("got: %v, want: %v", err, test.expect)
			}

		})
	}
}
