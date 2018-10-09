#!/bin/bash.exe
export GOOS=linux

go build -v -o ../linux/parsetop parsetop.go