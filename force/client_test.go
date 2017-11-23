package force

import "testing"

func TestNewClientInvalidInstance(t *testing.T) {
	_, err := NewClient("\n", UnitTest, "40.0", nil)
	if err == nil {
		t.Fatal("nil error returned")
	}
}

func TestNewClientInvalidVersion(t *testing.T) {
	_, err := NewClient("cs58.salesforce.com", UnitTest, "a.b", nil)
	if err == nil {
		t.Fatal("nil error returned")
	}
}
