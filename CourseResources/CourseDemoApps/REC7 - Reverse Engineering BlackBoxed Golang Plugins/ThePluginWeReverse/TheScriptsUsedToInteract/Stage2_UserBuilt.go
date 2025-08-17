package main

import (
	"fmt"
	"log"
	"plugin"
	"reflect"
)

func main() {
	plugin, x := plugin.Open("./3423n53n6536kjnm45k45ly.so")
	if x != nil {
		log.Fatal(x)
	}

	sym, x := plugin.Lookup("A25410863kb318989a95307")
	if x != nil {
		log.Fatal(x)
	}

	Routine := reflect.ValueOf(sym)
	if Routine.Kind() != reflect.Func {
		log.Fatal("[-] Symbol type is not a function. Must be function")
	}

	Unknown := sym.(func(string, string) string)
	fmt.Println(
		Unknown("hello", "world"),
	)
	///// Output - helldb3b7b61-orldcbd6
	/////					/\\__ looks like sha256 blended with plaintext?
	////// We now use this output to determine the contents of the function
	//////
}
