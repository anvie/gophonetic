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
	"regexp"
	//"fmt"
	"sre2"
)


/**
 * This is Caverphone algorithm version 2.0
 * based on paper: http://caversham.otago.ac.nz/files/working/ctp150804.pdf
 */
func Caverphone(text string) string {

	rv := ""
	
	length := len(text)
	
	if length == 0 {
		return rv
	}
	
	rv = strings.ToLower(text)
	
	re, e := regexp.Compile("[^a-z]")
	if e != nil {
		return rv
	}
	
	// remove non alphabet char
	rv = re.ReplaceAllString(rv, "")
	
	re2 := sre2.MustParse("^([crt]|(en)|(tr))ough")
	
	if caps := re2.Extract(rv, 3); len(caps) > 1 {
		d := caps[1] + "ou2f"
		l := len(d)
		rv = d + rv[l:]
	}
	
	re2 = sre2.MustParse("^gn")
	
	if match := re2.Match(rv); match {
		rv = "2n" + rv[2:]
	}
	
	re2 = sre2.MustParse("mb$")
	
	if match := re2.Match(rv); match {
		rv = rv[:-2] + "m2"
	}
	
	rv = strings.Replace(rv, "cq", "2q", -1)
	rv = strings.Replace(rv, "ci", "si", -1)
	rv = strings.Replace(rv, "ce", "se", -1)
	rv = strings.Replace(rv, "cy", "sy", -1)
	rv = strings.Replace(rv, "tch", "2ch", -1)
	rv = strings.Replace(rv, "c", "k", -1)
	rv = strings.Replace(rv, "q", "k", -1)
	rv = strings.Replace(rv, "x", "k", -1)
	rv = strings.Replace(rv, "v", "f", -1)
	rv = strings.Replace(rv, "dg", "2g", -1)
	rv = strings.Replace(rv, "tio", "sio", -1)
	rv = strings.Replace(rv, "tia", "sia", -1)
	rv = strings.Replace(rv, "d", "t", -1)
	rv = strings.Replace(rv, "ph", "fh", -1)
	rv = strings.Replace(rv, "b", "p", -1)
	rv = strings.Replace(rv, "sh", "s2", -1)
	rv = strings.Replace(rv, "z", "s", -1)
	
	
	re2 = sre2.MustParse("^[aiueo]")
	
	if match := re2.Match(rv); match {
		rv = "A" + rv[1:]
	}
	
	rv = strings.Replace(rv, "a", "3", -1)
	rv = strings.Replace(rv, "i", "3", -1)
	rv = strings.Replace(rv, "u", "3", -1)
	rv = strings.Replace(rv, "e", "3", -1)
	rv = strings.Replace(rv, "o", "3", -1)
	
	rv = strings.Replace(rv, "j", "y", -1)
	
	if rv[:2] == "y3" {
		rv = "Y3" + rv[2:]
	}
	
	if rv[0] == 'y' {
		rv = "A" + rv[1:]
	}
	
	rv = strings.Replace(rv, "y", "3", -1)
	rv = strings.Replace(rv, "3gh3", "3kh3", -1)
	rv = strings.Replace(rv, "gh", "22", -1)
	rv = strings.Replace(rv, "g", "k", -1)
	rv = strings.Replace(rv, "e", "3", -1)
	
	for _, sc := range []string{"s", "t", "p", "k", "f", "m", "n"} {
		re, _ = regexp.Compile(string(sc) + "+")
		rv = re.ReplaceAllString(rv, strings.ToUpper(sc))
	}
	
	rv = strings.Replace(rv, "w3", "W3", -1)
	rv = strings.Replace(rv, "wh3", "Wh3", -1)

	if rv[:len(rv)] == "w" {
		rv = rv[:len(rv)] + "3"
	}
	
	rv = strings.Replace(rv, "w", "2", -1)
	
	if rv[0] == 'h' {
		rv = "A" + rv[1:]
	}
	
	rv = strings.Replace(rv, "h", "2", -1)
	rv = strings.Replace(rv, "r3", "R3", -1)
	
	if rv[:len(rv)] == "r" {
		rv = rv[:len(rv)] + "3"
	}
	
	rv = strings.Replace(rv, "r", "2", -1)
	rv = strings.Replace(rv, "l3", "L3", -1)
	
	if rv[:len(rv)] == "l" {
		rv = rv[:len(rv)] + "3"
	}
	
	rv = strings.Replace(rv, "l", "2", -1)
	rv = strings.Replace(rv, "2", "", -1)
	
	if rv[len(rv)-1] == '3' {
		rv = rv[:len(rv)-1] + "A"
	}
	
	rv = strings.Replace(rv, "3", "", -1)
	
	rv = rv + "1111111111"

	return rv[0:10]
}



