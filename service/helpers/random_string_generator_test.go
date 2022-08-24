package helpers

import "testing"

func TestRandomString(t *testing.T) {
	result := RandomString(10) // 10 is count of letters will generate

	if len(result) != 10 {
		t.Error("fail to generate random string")
	}
}
