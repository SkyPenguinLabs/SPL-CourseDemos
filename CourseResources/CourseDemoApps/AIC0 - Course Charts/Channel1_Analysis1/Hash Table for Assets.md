
> The following table displays a list of SHA-1 hashes of all images in this asset directory. 

| Asset Name | Sha1 Hash                                |
| ---------- | ---------------------------------------- |
| A.png      | 37519FDCD40B32CB7B050AF6F3886536BF9ABF98 |
| B.png      | 03AAFAE1EC2B9BC6F2AA2BEEB162A07A28F91851 |
| C.png      | C4F6004AD49A8635286615FD6B4CC1E0198A43AB |
| D.png      | 200E0DC07AF529198E2EF9887485DD4BD52525B0 |
| E.png      | 80DB293191F47FF434E360A51FB4BF32FC4A575B |
| F.png      | 4E63DA2841011E1DFB8E5C6AE214ED729AAE0248 |

> Note, this was run inside of a Windows 10 Virtual Machine using a Windows Powershell interpreter command to do this was the following *(for users to understand the output format)*

```powershell
Get-FileHash -Path "*.png" -Algorithm SHA1
```

