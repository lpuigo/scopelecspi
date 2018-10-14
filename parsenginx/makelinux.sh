#!/bin/bash.exe
export GOOS=linux

go build -v -o ../linux/parsenginx parsenginx.go