/**
*
* 	Phonetic library by Robin Syihab (r [at] nosql.asia)
*
*	License: MIT
*
*	Copyright (c) 2009 The Go Authors. All rights reserved.
*
**/

package phonetic

import (
	"testing"
)

type soundexTest struct {
	in, out string
}

var soundexTests = []soundexTest {
	soundexTest{"robin", "R1500"},
	soundexTest{"anis", "A5200"},
	soundexTest{"YouKnowYouAllRight", "Y2546"},
}

func TestSoundex(t *testing.T) {
	for _, dt := range soundexTests {
		rv := Soundex(dt.in, 5)
		if rv != dt.out {
			t.Errorf("Get(%s) = `%s`, want `%s`", dt.in, rv, dt.out)
		}
	}
}


