//////// Use this to specify an unknown or unreversed function and quickly obtain args
/*
///
/// Example run:
///
///		)> go run CallPluginFunc.go ../3423n53n6536kjnm45k45ly.so <function_to_call>
///
///Then use the error output to determine the argument types necessary for the function
///
///------- OUTPUT example

													 /_____ Function takes two arguments of type string, not ()
													/
panic: interface conversion: plugin.Symbol is func(string, string) string, not func()

goroutine 1 [running]:
main.Caller()
	/home/xxxxxxxxx/Desktop/GolangCtf/RETools/CallPluginFunc.go:28 +0x268
main.main()
	/home/xxxxxxxxx/Desktop/GolangCtf/RETools/CallPluginFunc.go:33 +0xf
exit status 2


///
///
*/
package main

var (
	Hello string = ""
)

import (
	"fmt"
	"log"
	"os"
	"plugin"
	"reflect"
	"regexp"
	"runtime"
)

func Stage1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovery caught! ") // This is a message syntax
			///// we detect by recovery					/
			//// interface conversion: plugin.Symbol is func(string, string) string, not func()
			re := regexp.MustCompile(`is\s+(func\([^)]*\)(?:\s+\S+)?),\s+not\s+(func\([^)]*\)(?:\s+\S+)?)`)
			matches := re.FindStringSubmatch(r.(*runtime.TypeAssertionError).Error())
			if len(matches) > 2 {
				fmt.Println("Expected |> ", matches[1])
				fmt.Println("Test     |> ", matches[2])
			}
		}
	}()

	///// File specified by user
	p, err := plugin.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	///// Symbol specified by user
	sym, err := p.Lookup(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	Routine := reflect.ValueOf(sym)
	if Routine.Kind() != reflect.Func {
		log.Fatal("[-] Symbol type is not a function. Must be function")
	}

	//// Intentionally attempt an unknown type conversion, which most likely faults if the plugins function
	//// takes some form of argument
	AttemptCall := sym.(func())
	AttemptCall()
}

func main() {
	Stage1()
}
