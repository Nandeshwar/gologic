# tba

This project is used to find TBA in SIMS, GN, and COSMOS and fix them.

### clone project in folder given below
```
~/go/src
```

### To build executable file
```
go build cmd/tba/main.go
```

### To run project
```
go run cmd/tba/main.go
```

### config set up
```
https://confluence.dtc.dish.corp/pages/viewpage.action?spaceKey=SUPAIR&title=j.+TBA+issue%3A+temporary+fix+and+prevention+with+TBA+tool+version+8
```

### To get build info about executable file
```
GOOS - Target Operating System	GOARCH - Target Platform
android	arm
darwin	386
darwin	amd64
darwin	arm
darwin	arm64
dragonfly	amd64
freebsd	386
freebsd	amd64
freebsd	arm
linux	386
linux	amd64
linux	arm
linux	arm64
linux	ppc64
linux	ppc64le
linux	mips
linux	mipsle
linux	mips64
linux	mips64le
netbsd	386
netbsd	amd64
netbsd	arm
openbsd	386
openbsd	amd64
openbsd	arm
plan9	386
plan9	amd64
solaris	amd64
windows	386
windows	amd64



[DishNetwork] nandeshwa.sah@che-imssked01: ~ # cat /etc/os-release
PRETTY_NAME="Debian GNU/Linux 10 (buster)"
NAME="Debian GNU/Linux"
VERSION_ID="10"
VERSION="10 (buster)"
VERSION_CODENAME=buster
ID=debian
HOME_URL="https://www.debian.org/"
SUPPORT_URL="https://www.debian.org/support"
BUG_REPORT_URL="https://bugs.debian.org/"
[DishNetwork] nandeshwa.sah@che-imssked01: ~ # uname -r
4.19.0-17-amd64
[DishNetwork] nandeshwa.sah@che-imssked01: ~ #
```

Based on above info cross compile the code che-imsked01
```
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build cmd/tba/main.go
```
