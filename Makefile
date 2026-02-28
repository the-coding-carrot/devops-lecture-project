GO_CMD = go

all: build test

build:
	$(GO_CMD) build -o ./$(service)/out/$(service) ./$(service)/cmd/main.go
	

test:
	$(GO_CMD) test -v ./$(service)/...

clean:
	rm -f ./$(service)/out/$(service)