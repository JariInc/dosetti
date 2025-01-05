FROM node:23 AS build-node
WORKDIR /build

COPY package.json package-lock.json ./
COPY web ./web

RUN npm install
RUN npx esbuild ./web/js/*.js --bundle --outfile=web/assets/bundle.js --minify
RUN npx tailwindcss -i ./web/css/tailwind.css -o web/assets/style.css --minify

FROM golang:1.23 AS build-go
WORKDIR /build

COPY cmd ./cmd
COPY web ./web
COPY internal ./internal
COPY go.mod go.sum ./

COPY --from=build-node /build/web/assets web/assets

RUN go build -ldflags "-s -w -extldflags '-static'" -o ./build/dosetti cmd/dosetti/main.go

FROM scratch
COPY --from=build-go /build/build/dosetti /dosetti
COPY --from=build-go /build/web/html /web/html

ENTRYPOINT ["/dosetti"]
