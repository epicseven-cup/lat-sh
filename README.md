# lat-sh
cli tool for https://lat.sh, url shorter

```bash
lat-sh -s theapothecarydiariesopening3 -d https://www.youtube.com/watch?v=wOleNo7T6_4
```

# Install
1. You will need to install go on your machine: https://go.dev/doc/install
2. Setup GOPATH

Add the following to your shell config
```bash
export PATH=${PATH}:$HOME/go/bin
```
More information: https://go.dev/wiki/GOPATH#gopath-variable

3. Install the binary
```bash
go install github.com/epicseven-cup/lat-sh@latest
```

There could be delays between the Goproxy and GitHub binary, you can use the direct setup
```bash
GOPROXY=direct go install github.com/epicseven-cup/lat-sh@latest
```
