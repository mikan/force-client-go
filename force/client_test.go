package force

import "testing"

const (
	sampleAPIVer    = "40.0"
	sampleSObjectID = "0037F00000Hc2GyQAJ"
)

func TestNewClientInvalidVersion(t *testing.T) {
	_, err := NewClient(UnitTest, "a.b", nil)
	if err == nil {
		t.Fatal("nil error returned")
	}
}
