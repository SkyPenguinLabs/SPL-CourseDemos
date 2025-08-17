package main

import (
	"log"
	"plugin"
	"reflect"
)

// // This is the code you build to test the outcomes
func main() {
	plugin, x := plugin.Open("./3423n53n6536kjnm45k45ly.so")
	if x != nil {
		log.Fatal(x)
	}

	//// getting the symbol, we use what we found in Ghidra
	/// 							<- plugin/unnamed-f866734334be937294303517e47490150358c1f9.A25410863kb318989a95307
	sym, x := plugin.Lookup("A25410863kb318989a95307")
	if x != nil {
		log.Fatal(x)
	}

	///// Verify the symbol is a function, we should expect it to be one before type converting
	Routine := reflect.ValueOf(sym)
	if Routine.Kind() != reflect.Func {
		log.Fatal("[-] Symbol type is not a function. Must be function")
	}

	///// Now we type convert it to what we discoverd about the arguments and the function in ghidra
	AttemptCall := sym.(func(string, int, int, int))
	//// Attempt a call
	AttemptCall("hello", 10, 20, 14)

}
