include $(GOROOT)/src/Make.inc

TARG=srv
GOFILES=\
	tersrvbackend.go\
	datastructs.go\
	rest.go\
	dbmongo.go\

include $(GOROOT)/src/Make.cmd

