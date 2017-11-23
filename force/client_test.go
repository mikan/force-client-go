package force

import "testing"

const (
	sampleAPIVer    = "40.0"
	sampleSObjectID = "0037F00000Hc2GyQAJ"
)

func TestNewClientInvalidInstance(t *testing.T) {
	_, err := NewClient("\n", UnitTest, sampleAPIVer, nil)
	if err == nil {
		t.Fatal("nil error returned")
	}
}

func TestNewClientInvalidVersion(t *testing.T) {
	_, err := NewClient("localhost", UnitTest, "a.b", nil)
	if err == nil {
		t.Fatal("nil error returned")
	}
}
