all:
	make build
	make server

server:
	go run main.go net.go

install:
	cd client && npm install

build:
	make install
	cd client && npm run build
	go build main.go net.go