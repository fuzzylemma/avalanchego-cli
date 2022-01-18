PORT=9648
tidy:
	go mod tidy

build: tidy
	go build -o goavax

test: build
	./goavax health healthy -p ${PORT} 
