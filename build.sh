#!/usr/bin/env bash
#!/bin/bash
export GOPATH=$GOPATH:`pwd`
export GOARCH=amd64

export GOOS=linux
go build -ldflags "-w -s -linkmode internal" -o tcproxy-linux src/tcproxy.go
export GOOS=darwin
go build -ldflags "-w -s -linkmode internal" -o tcproxy-mac src/tcproxy.go
export GOOS=windows
go build -ldflags "-w -s -linkmode internal" -o tcproxy-win.exe src/tcproxy.go
