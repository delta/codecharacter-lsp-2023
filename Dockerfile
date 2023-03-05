## Base
FROM golang:1.18-alpine AS base

RUN apk add --no-cache build-base ccls py3-lsp-server

## Build
FROM base AS build

WORKDIR /app

COPY go.mod go.sum  ./
RUN go mod download

COPY . .

RUN go build -o server

## Dev
FROM build AS dev

WORKDIR /app

RUN apk add --no-cache make

RUN go install github.com/cespare/reflex@latest

ENTRYPOINT ["./entry.sh"]
CMD ["make watch"]


## Prod
FROM base AS prod

WORKDIR /

COPY --from=build /app/server /app/entry.sh /app/.env  /

ENTRYPOINT ["/entry.sh"]
CMD ["./server"]