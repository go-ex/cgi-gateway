package constants

import "testing"

func TestGetAppAgent(t *testing.T) {
	if got := GetAppAgent(1); got != "" {
		t.Errorf("GetAppAgent() = %v, want %v", got, "empty string")
	}

	AddAppAgent(1, "http://baidu.com")

	if got := GetAppAgent(1); got == "" {
		t.Errorf("GetAppAgent() = %v, want %v", got, "url string")
	}
}
