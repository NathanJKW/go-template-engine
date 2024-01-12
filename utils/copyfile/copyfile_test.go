package copyfile

import (
	"testing"
)

func TestPass(t *testing.T) {
	if output := BasicTest(1, 2); output != 3 {
		t.Errorf("error")
	}
}

func TestFail(t *testing.T) {
	if output := BasicTest(1, 2); output != 3 {
		t.Errorf("error")
	}
}
