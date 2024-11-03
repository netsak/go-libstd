package text_test

import (
	"testing"

	"github.com/netsak/go-libstd/text"
)

func TestToSnakeCase(t *testing.T) {
	testCases := []struct {
		str  string
		want string
	}{
		{"AccessToken", "access_token"},
		{"Tomcat", "tomcat"},
		{"HTMLElement", "html_element"},
		{"testingT", "testing_t"},
		{"text", "text"},
	}
	for _, tc := range testCases {
		t.Run(tc.str, func(t *testing.T) {
			got := text.ToSnakeCase(tc.str)
			if tc.want != got {
				t.Errorf("want: %q, got: %q", tc.want, got)
			}
		})
	}
}

func TestToEnvVar(t *testing.T) {
	testCases := []struct {
		str  string
		want string
	}{
		{"AccessToken", "ACCESS_TOKEN"},
		{"Tomcat", "TOMCAT"},
		{"HTMLElement", "HTML_ELEMENT"},
		{"testingT", "TESTING_T"},
		{"text", "TEXT"},
	}
	for _, tc := range testCases {
		t.Run(tc.str, func(t *testing.T) {
			got := text.ToEnvVar(tc.str)
			if tc.want != got {
				t.Errorf("want: %q, got: %q", tc.want, got)
			}
		})
	}
}
