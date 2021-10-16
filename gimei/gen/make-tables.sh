#!/bin/bash
go run gen.go > ../tables.txt
pushd ..
rm tables.go
mv tables.txt tables.go
