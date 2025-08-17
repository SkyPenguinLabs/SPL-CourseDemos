//////////////////////////////////////////////////////////// README //////////////////////////////////////////////////////////////////
/*			/////////////////
//////////////
			_______ _     _ __   __  _____  _______ __   _  ______ _     _ _____ __   _        _______ ______  _______
			|______ |____/    \_/   |_____] |______ | \  | |  ____ |     |   |   | \  | |      |_____| |_____] |______
			______| |    \_    |    |       |______ |  \_| |_____| |_____| __|__ |  \_| |_____ |     | |_____] ______|
//////////
  [ + ] |
  [ + ] |  This is designed for people to read! :D
  [ + ] |
  [ + ] |
  [ + ] | 				> Welcome to the SkyPenguinLabs Golang Demo/CTF thingy I developed in some short time. This was
  [ + ] | 				  for the REC7 course, which taught the idea of reverse engineering plugins to discover functionalities
  [ + ] | 				  and then write programs in Go to interact with them.
  [ + ] |
  [ + ] | 				  This application is not an actual VPN, but is built with fyne to mimic the idea of a login form, for a
  [ + ] | 				  theoretical VPN application which relies on two separate plugins, both written in Go, to hoist the login
  [ + ] | 				  functionality!
  [ + ] |
  [ + ] | 				  This is really not that hard of an application to understand as its the most minimalistic GUI, and
  [ + ] | 				  at most, some fancy plugin management logic which has notes like this documenting their purpose
  [ + ] | 				  and haxxers personal struggle with developing such hellish systems.
  [ + ] |
  [ + ] |
  [ + ] | 								~ Enjoy the ride ^_^ | @Totally_Not_A_Haxxer
  [ + ] |
  [ + ] |
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
*/
package main

import (
	"fmt"

	SplCore "main/Core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	AppWindow fyne.Window
)

// //////// @Initializer
func init() {
	App := app.New()
	App.Settings().SetTheme(&SplCore.DarkTheme{})
	AppWindow = App.NewWindow("SkyPenguinLabs VPN Login")
	AppWindow.Resize(
		fyne.NewSize(
			400, 500,
		),
	)
	AppWindow.CenterOnScreen()
}

func main() {
	///// @Section > Header
	title := widget.NewLabelWithStyle("Welcome Back!", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	subtitle := widget.NewLabelWithStyle("Sign in to continue", fyne.TextAlignCenter, fyne.TextStyle{})

	///// @Section > Username / Password Input
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Username or Email")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	///// @Widget
	loginBtn := widget.NewButton("Sign In", func() {
		SplCore.SPL_CoreAPI_Login(usernameEntry.Text, passwordEntry.Text)
	})
	loginBtn.Importance = widget.HighImportance

	///// @FormHandler
	usernameEntry.OnSubmitted = func(string) { passwordEntry.FocusGained() }
	passwordEntry.OnSubmitted = func(string) { loginBtn.OnTapped() }

	///// @Section > Remember me
	rememberCheck := widget.NewCheck("Remember me", nil)

	///// @Section > Forgot password
	forgotLink := widget.NewHyperlink("Forgot password?", nil)
	forgotLink.OnTapped = func() { fmt.Println("Forgot password clicked") }
	rememberRow := container.NewBorder(nil, nil, rememberCheck, forgotLink)

	///// @Section > Create account
	createLabel := widget.NewLabel("Don't have an account?")
	createLink := widget.NewHyperlink("Create one", nil)
	createLink.OnTapped = func() { fmt.Println("Create account clicked") }
	createRow := container.NewHBox(createLabel, createLink)

	////// @GUI - Content in order of the layout
	content := container.NewVBox(
		widget.NewSeparator(),
		title,
		subtitle,
		widget.NewSeparator(),
		container.NewPadded(container.NewVBox(
			usernameEntry,
			passwordEntry,
			rememberRow,
			widget.NewSeparator(),
			loginBtn,
			widget.NewSeparator(),
			container.NewCenter(createRow),
		)),
	)

	////// @GUI - Location & Sizing
	content.Resize(
		fyne.NewSize(
			350, 450,
		),
	)
	AppWindow.SetContent(container.NewBorder(
		nil, nil, nil, nil,
		container.NewCenter(container.NewWithoutLayout(content)),
	))
	AppWindow.ShowAndRun()
}
