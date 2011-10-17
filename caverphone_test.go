// @todo{robin}: code this

package phonetic

import (
	"testing"
	"fmt"
)


// Check the given state to be true.
func checkState(t *testing.T, state bool, err string) {
	if !state {
		t.Error(err)
	}
}

// Check string match
func checkString(t *testing.T, result, expected, err string) {
	checkState(t, result == expected, fmt.Sprintf("%s: got %s, expected %s", err, result, expected))
}


func checkCapture(t *testing.T, expected []string, result []string, err string) {
	match := true
	if (expected == nil || result == nil) && (expected != nil && result != nil) {
		match = false
	} else if len(expected) != len(result) {
		match = false
	} else {
		for i := 0; i < len(expected); i++ {
			if expected[i] != result[i] {
				match = false
			}
		}
	}
	checkState(t, match, fmt.Sprintf("%s: got %s, expected %s", err, result, expected))
}


func TestCaverphone(t *testing.T) {
	checkString(t, Caverphone("mayer"), "MA11111111", "should match")
	checkString(t, Caverphone("meier"), "MA11111111", "should match")
	checkString(t, Caverphone("Henrichsen"), "ANRKSN1111", "should match")
	checkString(t, Caverphone("Henricsson"), "ANRKSN1111", "should match")
	checkString(t, Caverphone("Henriksson"), "ANRKSN1111", "should match")
	checkString(t, Caverphone("Hinrichsen"), "ANRKSN1111", "should match")
	checkString(t, Caverphone("Stevenson"), "STFNSN1111", "should match")
	checkString(t, Caverphone("Peter"), "PTA1111111", "should match")
	checkString(t, Caverphone("Karleen,"), "KLN1111111", "should match")
	checkString(t, Caverphone("Thompson"), "TMPSN11111", "should match")
	checkString(t, Caverphone("Whitlam"), "WTLM111111", "should match")
}
