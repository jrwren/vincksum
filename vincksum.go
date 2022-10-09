package vincksum

import (
	"strings"
)

// some vin constant
var w = []int{8, 7, 6, 5, 4, 3, 2, 10, 0, 9, 8, 7, 6, 5, 4, 3, 2}

// Checksum returns the checksum for the vin or -1 on error.
func Checksum(vin string) int {
	t := 0
	for i := 0; i < len(vin); i++ {
		t += int(charToNum(vin[i])) * w[i]
	}
	t = t % 11
	if t == 10 {
		return -1
	}
	return t
}

// Validate returns true if the vin checksum matches the 8th byte.
func Validate(vin string) bool {
	if len(vin) != 17 || strings.ContainsAny(vin, "IOQ") {
		return false
	}
	return Checksum(vin) == int(vin[8])-48
}

func charToNum(n byte) byte {
	if n <= byte('9') {
		return n - byte('0')
	}
	if n < byte('I') {
		return n - byte('A') + 1
	}
	if n <= byte('R') {
		return n - byte('J') + 1
	}
	return n - byte('S') + 2
}
