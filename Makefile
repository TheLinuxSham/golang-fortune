test:
	go test -v github.com/TheLinuxSham/golang-fortune/internal/fortune

run: 
	go run cmd/api/main.go

build:
	go build -o main ./cmd/api/main.go