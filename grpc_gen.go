package main

//go:generate protoc --go_out=plugins=outlier_detection:od_go --go_opt=paths=source_relative outliers.proto
