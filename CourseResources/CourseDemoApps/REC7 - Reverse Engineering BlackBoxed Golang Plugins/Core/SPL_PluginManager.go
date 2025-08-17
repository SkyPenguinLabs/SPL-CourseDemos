//////////////////////////////////////////////////////////// README //////////////////////////////////////////////////////////////////
/*			/////////////////
//////////////
			_______ _     _ __   __  _____  _______ __   _  ______ _     _ _____ __   _        _______ ______  _______
			|______ |____/    \_/   |_____] |______ | \  | |  ____ |     |   |   | \  | |      |_____| |_____] |______
			______| |    \_    |    |       |______ |  \_| |_____| |_____| __|__ |  \_| |_____ |     | |_____] ______|
//////////
  [ + ] |
  [ + ] | 				> This was a fun part of the development. Just- fun. Plugins in Go need to be managed, and I have found, the
  [ + ] |  				  best way to prevent issues with dynamically loading in-process plugins is to build some sort of manager that
  [ + ] |				  is super simple and super easy to debug. I have always stuck with this design as a baseline for plugin systems
  [ + ] |				  used in go, and would build ontop.
  [ + ] |
  [ + ] |			      This one manages multiple plugins at once. This is not ideal, as go plugins, probably should only be called upon
  [ + ] |			      once they are used. I like this concept because it already exists in the blockchain landscape with smart contracts.
  [ + ] |
  [ + ] |				  smart contracts, are stateless programs, deployed onto a network using RPC known as 'blockchain nodes'
  [ + ] |				  which allows any off-node/off-chain application (such as a server, or computer like yours with a client like
  [ + ] | 				  curl) to make invokes/function calls in relation to that program. However, functions are called once.
  [ + ] |
  [ + ] |				  Plugins in go, should be treated similar. In this case, they arent. They are loaded and kept in a live state.
  [ + ] |				  The reason its a good practice is because it manages the landscape of issues that can oncur with a plugin living
  [ + ] |				  on throughout the process. Sometimes, I have noticed that the longer a plugin sits active, the more bound it is
  [ + ] |				  to a heisenbug. You can execute the ideal state by simply
  [ + ] |
  [ + ] |						Call_Open_Plugin		<- OPEN IT IMMEDIATELY before you call the plugin function
  [ + ] |						Call_PluginSymbol
  [ + ] |						Call_UnloadPlugin    <- IMMEDIATELY after the plugin results are fetched
  [ + ] |
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
*/
package Core

import (
	"fmt"
	"plugin"
	"reflect"
	"sort"
)

type SPL_PluginManager struct {
	Plugins       map[string]*plugin.Plugin
	PluginSymbols map[string]map[string]plugin.Symbol
}

func NewPluginSession() *SPL_PluginManager {
	return &SPL_PluginManager{
		Plugins:       make(map[string]*plugin.Plugin),
		PluginSymbols: make(map[string]map[string]plugin.Symbol),
	}
}

// //// Register a plugin by name specific to its path
func (Manager *SPL_PluginManager) RegisterPlugin(PluginName, PluginPath string) error {
	Plugin, x := plugin.Open(PluginPath)
	if x != nil {
		return x
	}
	Manager.Plugins[PluginName] = Plugin
	Manager.PluginSymbols[PluginName] = make(map[string]plugin.Symbol)
	return nil
}

// //// Registers a function in relation to a specific plugin name
func (Manager *SPL_PluginManager) RegisterSymbol(PluginName, SymName string) error {
	p, exists := Manager.Plugins[PluginName]
	if !exists {
		return fmt.Errorf("[-] plugin %s not registered", PluginName)
	}

	symbol, x := p.Lookup(SymName)
	if x != nil {
		return fmt.Errorf("[-] symbol %s not found in registered plugin %s: %v", SymName, PluginName, x)
	}

	Manager.PluginSymbols[PluginName][SymName] = symbol
	return nil
}

// //// Gets the symbol or function based on the plugins and the symbols name
func (Manager *SPL_PluginManager) GetSymByName(PluginName, Symname string) (plugin.Symbol, error) {
	if syms, exists := Manager.PluginSymbols[PluginName]; exists {
		if sym, exists := syms[Symname]; exists {
			return sym, nil
		}
	}
	return nil, fmt.Errorf("[-] symbol %s not found in plugin %s", Symname, PluginName)
}

// ///// Gets all of the REGISTERED plugins
func (Manager *SPL_PluginManager) GetRegisteredPlugins() []string {
	names := make([]string, 0, len(Manager.Plugins))
	for name := range Manager.Plugins {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// ///// Gets all of the REGISTERED plugins
func (Manager *SPL_PluginManager) GetRegisteredSymbols(PluginName string) []string {
	symbols, exists := Manager.PluginSymbols[PluginName]
	if !exists {
		return nil
	}

	names := make([]string, 0, len(symbols))
	for name := range symbols {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (Manager *SPL_PluginManager) UnloadPlugin(name string) {
	delete(Manager.Plugins, name)
	delete(Manager.PluginSymbols, name)
}

//////////////////
////
//// @CALLER  - Calls the plugin function
////
////

func (
	Manager *SPL_PluginManager,
) CallPluginFunction(PluginName string, Symname string, Symargs ...interface{}) ([]reflect.Value, error) {
	sym, x := Manager.GetSymByName(PluginName, Symname)
	if x != nil {
		return nil, x
	}
	Routine := reflect.ValueOf(sym)
	if Routine.Kind() != reflect.Func {
		return nil, fmt.Errorf("[-] Symbol type is not a function. Must be function")
	}

	///// Args
	argValues := make([]reflect.Value, len(Symargs))
	for i, arg := range Symargs {
		argValues[i] = reflect.ValueOf(arg)
	}

	results := Routine.Call(argValues)
	return results, nil
}
