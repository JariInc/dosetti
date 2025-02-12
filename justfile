set dotenv-load

cpus := if os() == "macos" { `sysctl -n hw.ncpu` } else { `getconf _NPROCESSORS_ONLN` }

build: build_css build_js build_go

run: build
    ./build/dosetti

lint: lint_go

test:
    go test ./test/... -v -parallel {{cpus}}

watch:
    trap 'kill 0' SIGINT; \
    just watch_go & \
    just watch_js & \
    just watch_css & \
    wait

clean:
    rm -rf build/*
    # delete odd OS X junk files from SMB share
    find . -name '._*' -delete
    find . -name '.smbdelete*' -delete

alias migrate := migrate_up

migrate_up:
    goose -dir ./migrations up

migrate_down:
    goose -dir ./migrations down

db_shell:
    sqlite3 dosetti.db

seed:
    just db_shell < seed.sql

install:
    npm install
    go mod download

lint_go:
    go fmt github.com/jariinc/...

build_go:
    go build -o ./build/dosetti cmd/dosetti/main.go

build_css:
    node_modules/.bin/tailwindcss -i ./web/css/tailwind.css -o web/assets/style.css --minify

build_js:
    node_modules/.bin/esbuild ./web/js/*.js --bundle --outfile=web/assets/bundle.js --minify

watch_go:
    air

watch_js:
    node_modules/.bin/esbuild ./web/js/*.js --bundle --outfile=web/assets/bundle.js --watch

watch_css:
    node_modules/.bin/tailwindcss -i ./web/css/tailwind.css -o web/assets/style.css --watch

build_docker:
    docker build -t dosetti:latest .

run_docker:
    docker run --rm -e DATABASE_URL="file:./dosetti.db" -p 8080:8080 dosetti:latest