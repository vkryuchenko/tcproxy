#!/usr/bin/env bash
#!/bin/bash
export GOPATH=$GOPATH:`pwd`
export GOARCH=amd64

export GOOS=linux
go build -ldflags "-w -s -linkmode internal" -o tcproxy-linux tcproxy.go
export GOOS=darwin
go build -ldflags "-w -s -linkmode internal" -o tcproxy-mac tcproxy.go
export GOOS=windows
go build -ldflags "-w -s -linkmode internal" -o tcproxy-win.exe tcproxy.go
