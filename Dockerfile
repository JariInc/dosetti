FROM node:current-slim AS build-node
WORKDIR /build

COPY package.json package-lock.json tailwind.config.js ./
RUN npm install

COPY web web
RUN npx esbuild ./web/js/*.js --bundle --outfile=web/assets/bundle.js --minify
RUN npx tailwindcss -i ./web/css/tailwind.css -o web/assets/style.css --minify

FROM golang:1.24 AS build-go
WORKDIR /build

COPY cmd ./cmd
COPY web ./web
COPY internal ./internal
COPY go.mod go.sum ./
COPY --from=build-node /build/web/assets web/assets

RUN go build -ldflags "-s -w -extldflags '-static'" -o ./build/dosetti cmd/dosetti/main.go

FROM scratch
COPY --from=build-go /build/build/dosetti /dosetti
COPY web/html /web/html
COPY migrations /migrations

ENTRYPOINT ["/dosetti"]
