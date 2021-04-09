

run: main.go
	go run *.go

build: main.go **.go
	go build
