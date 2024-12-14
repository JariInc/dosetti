build:
    go build -o ./build/dosetti cmd/dosetti/main.go

run: build
    ./build/dosetti

lint:
    go fmt github.com/jariinc/...

watch:
    air
