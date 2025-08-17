# SkyPenguinLabs Demo VPN Application

In lieu of reverse engineering Go, I figured I would do something special. So I went absolutely HAM on this and built a small login page for the first time with Fyne in Go. 

This application is only designed to run on Linux, as it uses in-process plugins within the Go programming language, and uses X11 as a rendering backbone, using fyne as the wrapoer. In order to set this up
it did take some work, so I decided to give you a descriptive document explaining how to set it up and what it looks like!


---
### Setting Up
---

In order to setup, clone all of the contents within this directory, `cd` into it, and then run the following commands in order below.

> [!TIP]
> This application was compiled using Go v1.12.8 linux/amd64. Make sure the compiler version is this one or some version around; anything too recent has not been tested, such as 1.25.0 

#### Step 1 
* if you do not have LibX11 devkits installed, pls install them. The following commands are for a debian system, and were run on a parrot6 machine running inside of Vmware Workstation Pro v17. 

```
sudo apt-get install libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libgl1-mesa-dev xorg-dev
sudo apt install libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libgl1-mesa-dev
```

#### Step 2 
* Once you have those libraries installed...

```
go get . 
```

To install all of the libraries and requirements. Either or, they are all very simple.

```
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
```

#### Step 3

* Now you can build the plugins, and then the application

```
go build -buildmode=plugin -ldflags="-s -w" -o keygen.so ./plugins/keygen/keygenplugin.go
go build -buildmode=plugin -ldflags="-s -w" -o postlogin.so ./plugins/postlogin/postlogin.go
go build -o app main.go
```

---
### Running the application!
---

You are now set to run the demo application, which loads two separate in-process plugins developed to assist the login functionality!

```
./app
```

# Demo Screenshots 



  

  

