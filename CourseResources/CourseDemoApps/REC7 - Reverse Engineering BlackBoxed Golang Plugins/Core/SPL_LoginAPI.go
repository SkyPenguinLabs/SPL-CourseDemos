//////////////////////////////////////////////////////////// README //////////////////////////////////////////////////////////////////
/*			/////////////////
//////////////
			_______ _     _ __   __  _____  _______ __   _  ______ _     _ _____ __   _        _______ ______  _______
			|______ |____/    \_/   |_____] |______ | \  | |  ____ |     |   |   | \  | |      |_____| |_____] |______
			______| |    \_    |    |       |______ |  \_| |_____| |_____| __|__ |  \_| |_____ |     | |_____] ______|
//////////
  [ + ] |
  [ + ] | 				> This is essentially the 'API' for accessing the functions within the plugins. Both of these functions
  [ + ] | 				  rely on externally compiled plugins, which are defined in the root of this directory, @./plugins.
  [ + ] |
  [ + ] |				  The plugins include
  [ + ] |
  [ + ] |							> keygenplugin.go |> This is what the first API function, @API_FUNC1, below, uses to
  [ + ] | 												 invoke a sort of software license cycle for the login system.
  [ + ] |
  [ + ] |							> postlogin.go    |> This is a filler plugin to add some noise to the code, but primarily takes
  [ + ] |												 the username and password in plaintext, and makes an HTTP(s) post request
  [ + ] |												 with those credentials to httpbin.org and sends back the response
  [ + ] |
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
*/
package Core

import (
	"fmt"
	"log"
)

const (
	PLUGIN_PATH_KEYGEN         = "./keygen.so"               //// @FilePath  |> Filepath to the keygen system plugin
	PLUGIN_PATH_CONTROLLER     = "./postlogin.so"            //// @Filepath  |> Login functionality
	PLUGIN_KEYGEN_VNAME        = "loginhelper"               //// @VRef      |> This is the name of our keygen plugin in this programs environment (refer to ./SPL_PLuginManager.go)
	PLUGIN_CONTROLLER_VNAME    = "login_post_remote_request" //// @VRef		 |> This is the name of the postlogin plugin in this environment
	PLUGIN_KEYGEN_VFUNC_KEYGEN = "A25410863kb318989a95307"   //// @SymName   |> This is the symbol name of the function in the keygen plugin
	PLUGIN_POSTFUNC_VFUNC      = "PostLogin"                 //// @SymName	 |> This is the symbol name of the function that we need to call in the postlogin plugin which makes a call to httpbin
)

// //////
// /
// / @HELPER_FUNC1 | CE (CheckError):
// /			Takes an input 'x' of type 'error' and checks
// /			if it is not nil. If not, calls log.fatal
func CE(x error) {
	if x != nil {
		log.Fatal(x)
	}
}

// //////
// /
// / @API_FUNC1
// /
// /
func SPL_CoreAPI_Login(username, password string) {
	pm := NewPluginSession()
	x := pm.RegisterPlugin(PLUGIN_KEYGEN_VNAME, PLUGIN_PATH_KEYGEN)
	CE(x)
	pm.RegisterSymbol(PLUGIN_KEYGEN_VNAME, PLUGIN_KEYGEN_VFUNC_KEYGEN)
	result, x := pm.CallPluginFunction(PLUGIN_KEYGEN_VNAME, PLUGIN_KEYGEN_VFUNC_KEYGEN, username, password)
	CE(x)
	fmt.Println(result)
	fmt.Printf("[*] Login attempt | \n> Username    %s\n> Password    %s\n", username, password)
	fmt.Println("[+] Making a call to the login invoker....")
	SPL_CoreAPISendLogin(username, password)
}

func SPL_CoreAPISendLogin(username, password string) {
	pm := NewPluginSession()
	x := pm.RegisterPlugin(PLUGIN_CONTROLLER_VNAME, PLUGIN_PATH_CONTROLLER)
	CE(x)
	/////// A good safe runtime thought here is making sure that we actually have to register every single symbol
	/////// even if the program already knows it exists inside the plugin when it loads it. This is to make sure that
	/////// a developer such as ourselves, may feel is safe to allow into the program. Even if, developers are not the
	/////// best at knowing what is 'safe' versus what is 'secure'.
	pm.RegisterSymbol(PLUGIN_CONTROLLER_VNAME, PLUGIN_POSTFUNC_VFUNC)
	result, x := pm.CallPluginFunction(PLUGIN_CONTROLLER_VNAME, PLUGIN_POSTFUNC_VFUNC, username, password)
	CE(x)
	fmt.Println("OK? ", result)
}
