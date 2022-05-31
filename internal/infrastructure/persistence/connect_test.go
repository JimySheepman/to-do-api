package persistence

import (
	"errors"
	"reflect"
	"testing"
)

func TestConnectDB(t *testing.T) {
	tests := []struct {
		testName   string
		driverName string
		dns        string
		err        error
	}{
		{"successful connection", "postgres", "postgres://postgres:root@localhost:5432/postgres?sslmode=disable", nil},
		{"dns error", "postgres", "postgres://postgres:root@localhost/postgres?sslmode=disable", nil},
		{"ping error", "postgres", "postgres://postgres:root@localhost:5432", errors.New(`pq: SSL is not enabled on the server`)},
		{"driver name error", "postg", "postgres://postgres:root@localhost:5432/postgres?sslmode=disable", errors.New(`sql: unknown driver "postg" (forgotten import?)`)},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			_, actual := ConnectDB(test.driverName, test.dns)
			if !reflect.DeepEqual(actual, test.err) {
				t.Errorf("got the %v value\nexpecting the %v value", actual, test.err)
			}
		})
	}
}
