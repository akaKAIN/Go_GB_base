run:
	go run .

test:
	go test -v ./...

cover:
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out

bench:
	go test -bench=. ./...

benchstat:
	go test -v -bench=. -count 5 -benchmem ./bubble/ >first.txt
	go test -v -bench=. -count 5 -benchmem ./insertion/ >second.txt
	bechstat first.txt second.txt

