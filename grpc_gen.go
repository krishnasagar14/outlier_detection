package main

//go:generate protoc --go_out=plugins=grpc:od_go --go_opt=paths=source_relative outliers.proto
