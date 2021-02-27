# /bin/bash

pwd
go run main.go
go run cmd/another/main.go
cd cmd
pwd
go run ../main.go
go run another/main.go 
