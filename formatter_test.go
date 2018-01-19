package logback

import (
	"testing"

	"github.com/sirupsen/logrus"
)

//  yyyy-MM-dd'T'HH:mm:ss.SSSZ [thread] %-5level logger(36) - msg\n</pattern>
func TestFormatting(t *testing.T) {
	tf := &Formatter{}

	testCases := []struct {
		key      string
		value    string
		expected string
	}{
		{`foo`, `bar`, "0001-01-01T00:00:00Z [main] ERROR unknown - {foo=bar}\n"},
		{`level`, `one`, "0001-01-01T00:00:00Z [main] ERROR unknown - {level=one}\n"},
		{`thread`, `turkey`, "0001-01-01T00:00:00Z [turkey] unknown - {}\n"},
		{`logger`, `turkey`, "0001-01-01T00:00:00Z [main] ERROR turkey - {}\n"},
	}

	for _, tc := range testCases {
		b, _ := tf.Format(logrus.WithField(tc.key, tc.value))

		if string(b) != tc.expected {
			t.Errorf("formatting expected for %q - result was:\n\n%s\ninstead of:\n\n%s", tc.value, string(b), tc.expected)
		}
	}
}
