package main

// #include <stdio.h>
// #include <stdlib.h>
/*
void print(char *str) {
    printf("%s\n", str);
}
*/
import "C"

import "unsafe"

func main() {
	s := "hello cgo"
	cs := C.CString(s)
	C.print(cs)
	C.free(unsafe.Pointer(cs))
	C.puts(C.CString("hello, world\n"))
}
