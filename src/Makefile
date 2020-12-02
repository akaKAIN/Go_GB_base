dir = "homework_1"
run:
	go run ./${dir}

test:
	go test -v ./${dir}

cover:
	go test -coverprofile=./${dir}/cover.out ./${dir}
	go tool cover -html=./${dir}/cover.out

