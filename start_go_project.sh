#!/bin/bash

# run go mod init with the name of the project as argument
mkdir $1;
cd $1;
go mod init $1;

mkdir cmd cmd/main internal tests

touch cmd/main/main.go internal/internal.go tests/tests.go

echo "package main" >> cmd/main/main.go
echo "package internal" >> internal/internal.go
echo "package tests" >> tests/tests.go

