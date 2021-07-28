package print

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

func Print(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs,(*C.FILE)(C.stdout))
}

//运行不了，等絮叨 makefile的时候再解决
/*
include $(GOROOT)/src/Make.inc
TARG=rand
CGOFILES=\
c1.go\
include $(GOROOT)/src/Make.pkg*/