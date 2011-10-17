include $(GOROOT)/src/Make.inc

TARG=phonetic
DEPS=sre2
GOFILES=\
	soundex.go \
	caverphone.go

include $(GOROOT)/src/Make.pkg
