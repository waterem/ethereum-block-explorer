package functions

import (
	"testing"
)

func TestPadLeftEven1(t *testing.T) {
	actual := padLeftEven("1")
	expected := "01"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestPadLeftEven2(t *testing.T) {
	actual := padLeftEven("01")
	expected := "01"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestSanitizeHex1(t *testing.T) {
	actual := sanitizeHex("01")
	expected := "0x01"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestSanitizeHex2(t *testing.T) {
	actual := sanitizeHex("0x01")
	expected := "0x01"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestTrimHexZero1(t *testing.T) {
	actual := trimHexZero("0000000000000000000000008f16172f2c2c61dbbd52fcf1773add")
	expected := "0x8f16172f2c2c61dbbd52fcf1773add"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
