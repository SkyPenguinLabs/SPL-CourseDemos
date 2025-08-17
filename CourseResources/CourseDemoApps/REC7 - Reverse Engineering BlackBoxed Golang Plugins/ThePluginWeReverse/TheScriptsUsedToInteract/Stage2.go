package main

import (
	"fmt"
	"log"
	"os"
	"plugin"
	"reflect"
)

func CE(x error) {
	if x != nil {
		log.Fatal(x)
	}
}

/////// Now that we know the args are  | (string, string) <- OUT[string] |
/////// we can construct a program to call the function, throw some input through it
/////// and see what happens

func Stage2() {
	//// Plugin [OPEN]
	p, x := plugin.Open("../3423n53n6536kjnm45k45ly.so")
	CE(x)
	//// Plugin [LOOKUP->SYMBOL]
	sym, x := p.Lookup("A25410863kb318989a95307")
	CE(x)

	Routine := reflect.ValueOf(sym)
	if Routine.Kind() != reflect.Func {
		log.Fatal("[-] Symbol type is not a function. Must be function")
	}

	//// Intentionally attempt an unknown type conversion, which most likely faults if the plugins function
	//// takes some form of argument
	AttemptCall := sym.(func(string, string) string)
	fmt.Println(
		"[+] Result when running function -> [" +
			AttemptCall(
				os.Args[1],
				os.Args[2],
			) + "]",
	)
}

func main() {
	Stage2()
}
