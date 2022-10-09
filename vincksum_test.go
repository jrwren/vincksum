package vincksum

import (
	"testing"
)

func TestValidateInvalidChar(t *testing.T) {
	if Validate("0I23456789ABCDEFG") {
		t.Error("expected false validation because if invalid I in vin")
	}
	if Validate("O123456789ABCDEFG") {
		t.Error("expected false validation because if invalid O in vin")
	}
	if Validate("Q123456789ABCDEFG") {
		t.Error("expected false validation because if invalid O in vin")
	}
}

func TestValidateInvalidLength(t *testing.T) {
	if Validate("") {
		t.Error("expected false validation because if invalid length")
	}
}

func TestChecksum(t *testing.T) {
	a := Checksum("1HGCP2F35CA038278")
	if a != 5 {
		t.Errorf("actual %d expected 5", a)
	}
}

func TestValidate(t *testing.T) {
	if !Validate("1HGCP2F35CA038278") {
		t.Errorf("failed to validate a valid vin")
	}
}

func TestSanitycharToNum(t *testing.T) {
	if charToNum('1') != 1 {
		t.Errorf("1 != 1 boy you dun messed up somewhere")
	}
}
