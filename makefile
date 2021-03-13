all:
	cd client && npm install
	go run *.go

server:
	go run *.go

build:
	cd client && npm install && npm run build
	go build *.go