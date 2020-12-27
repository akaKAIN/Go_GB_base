hw = "1"

run:
	go run ./homework_${hw}

test:
	go test -v ./...

cover:
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out

bench:
	go test -bench=. ./...