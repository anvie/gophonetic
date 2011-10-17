include $(GOROOT)/src/Make.inc

TARG=phonetic
GOFILES=\
	soundex.go \
	caverphone.go

include $(GOROOT)/src/Make.pkg
