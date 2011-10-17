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
	"strings"
)



var digits string = "01230120022455012623010202"

func isAlpha(ch int) bool {
	return ch <= 'z' && ch >= 'A'
}


func Soundex(name string, length int) string {
	
	sndx := ""
	var fc int = 0
	
	for _, c := range strings.ToUpper(name) {
		if isAlpha(c) {
			if fc == 0 {
				fc = c
			}
			d := digits[c - 'A']
			if sndx == "" || (d != sndx[len(sndx)-1]) {
				sndx += string(d)
			}
		}
	}
	
	if len(sndx) == 0 {
		return ""
	}
	
	sndx = string(fc) + sndx[1:]
	
	sndx = strings.Replace(sndx, "0", "", -1)
	
	zeros := ""
	
	for i := 0; i < length; i++ {
		zeros += "0"
	}
	
	return (sndx + zeros)[:length]
}
