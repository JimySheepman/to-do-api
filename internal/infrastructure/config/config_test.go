package config

import (
	"errors"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	var tests = []struct {
		testName        string
		envVariableName string
		envPath         string
		expected        error
	}{
		{"wrong path", "DB_HOST", "../.env", errors.New("wrong env path")},
		{"successful database host", "DB_HOST", "./.env", nil},
		{"database host load error", "DB_HOS", "./.env", errors.New("wrong env variable name")},
		{"successful database port", "DB_PORT", "./.env", nil},
		{"database port load error", "DB_POR", "./.env", errors.New("wrong env variable name")},
		{"successful database user", "DB_USER", "./.env", nil},
		{"database user load error", "DB_USE", "./.env", errors.New("wrong env variable name")},
		{"successful database password", "DB_PASSWORD", "./.env", nil},
		{"database password load error", "DB_PASSWOR", "./.env", errors.New("wrong env variable name")},
		{"successful database name", "DB_NAME", "./.env", nil},
		{"database name load error ", "DB_NAM", "./.env", errors.New("wrong env variable name")},
		{"successful kafka url", "KAFKA_URL", "./.env", nil},
		{"kafka url load error", "KAFKA_UR", "./.env", errors.New("wrong env variable name")},
		{"successful kafka topic", "KAFKA_TOPIC", "./.env", nil},
		{"kafka topic load error", "KAFKA_TOPI", "./.env", errors.New("wrong env variable name")},
		{"successful kafka groupId", "KAFKA_GROUP_ID", "./.env", nil},
		{"kafka groupId load error", "KAFKA_GROUP_I", "./.env", errors.New("wrong env variable name")},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			_, expected := Config(test.envVariableName, test.envPath)
			if !reflect.DeepEqual(expected, test.expected) {
				t.Errorf("got:%v  expect:%v", expected, test.expected)
			}
		})
	}
}
