# SPL Scripts used for reversing the golang plugin

All of the code within this repository is from the REC7, paid course, released here, by SkyPenguinLabs. 

The goal of this source code was completed in two stages, which involved using stage 1 to see if we could call a proprietary function within a black-boxed Go plugin, and stage 2, which took the results of stage 1 and used them to create results that were further taken into consideration
during analysis. To know more, correlate this repository with the course.


> [!TIP]
> This code has two copies. `StageX*.go` and `StageX_UserBuilt.go`. All `UserBuilt.go` files are ones that you are taught to build in the course. These are the ones you follow along with to create and use. On the other hand, I spent my own time refining the code, with some better, more
> advanced handling to express how you could extend the code. This happens in all files `StageX,go`, such as `Stage1.go`


---
### Scenario
---

When working with this plugin, the ideal scenario was that we stumbled across a file called `3423n53n6536kjnm45k45ly.so` which we have 0 information on, and 0 app to correlate with. This was found on an open directory, just sitting in Belarus. So, we decided to open it up. 

The goal of opening it up is to figure out what type of object file it is, what it does, how it works, and reverse engineer it to find all the requirements necessary to build a golang program for interfacing with it.

---
### Why no setup?
---

Ultimately, this code in the plugin `3423n53n6536kjnm45k45ly.so` runs in-process, which means you need to execute a golang program to interface with this code. It's not designed as an independent executable and was not compiled like one.

There are no environments or special setups for loading these go scripts because they all use standard libraries.

**THE ONLY REQUIREMENT IS YOU HAVE A LINUX SYSTEM** which almost every reverse engineer has. This is to make sure the plugin system and this main application can run. 
