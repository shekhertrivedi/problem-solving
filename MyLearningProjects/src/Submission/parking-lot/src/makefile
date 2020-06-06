build:
	go build main.go

test:
	go test -timeout 30s github.com/ContinuumLLC/MyLearningProjects/Submission/parking-lot/src/local/service -coverprofile=/tmp/vscode-goUBHZhX/go-code-cover
	go test -timeout 30s github.com/ContinuumLLC/MyLearningProjects/Submission/parking-lot/src/local/processor -coverprofile=/tmp/vscode-goRih24K/go-code-cover

all: build test run

run:
	./main

runService:
	go run main.go

runWithCommands:
	go build main.go
	./main ${ARGS}

package:
	rm -rf build
	mkdir build
	mv main build

