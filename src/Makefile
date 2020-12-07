hw = "1"

run:
	go run ./homework_${hw}

test:
	go test -v ./homework_${hw}/...

cover:
	go test -coverprofile=./homework_${hw}/cover.out ./homework_${hw}
	go tool cover -html=./homework_${hw}/cover.out

