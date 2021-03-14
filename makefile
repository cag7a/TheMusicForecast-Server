all:
	make build
	make server

server:
	go run *.go

install:
	cd client && npm install

build:
	make install
	cd client && npm run build
	go build *.go