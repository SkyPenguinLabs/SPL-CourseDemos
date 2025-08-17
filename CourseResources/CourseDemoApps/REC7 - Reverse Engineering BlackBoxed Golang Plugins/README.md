### SkyPenguinLabs Demo VPN Application

In lieu of reverse engineering Go, I figured I would do something special. So I went absolutely HAM on this and built a small login page for the first time with Fyne in Go. 

This application is only designed to run on Linux, as it uses in-process plugins within the Go programming language, and uses X11 as a rendering backbone, using fyne as the wrapoer. In order to set this up
it did take some work, so I decided to give you a descriptive document explaining how to set it up and what it looks like!


#### Setting Up

In order to setup, clone all of the contents within this directory, `cd` into it, and then run the following commands in order below.

> [!INFO] This application was compiled using Go v1.12.8 linux/amd64, make sure the compiler version is this one or some version around, anything too recent has not been tested, such as 1.25.0 


* if you do not have X11 devkits installed, pls install them.
