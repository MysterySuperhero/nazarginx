all:
	rm -rf ./httpd
	go build -o ./httpd ./src/main.go