## Build
FROM golang:1.18-alpine AS build

RUN apk add build-base

WORKDIR /app

COPY go.mod go.sum  ./
RUN go mod download

COPY . .

RUN go build -o server

## Dev
FROM build AS dev

WORKDIR /app

RUN apk add --no-cache make ccls

RUN go install github.com/cespare/reflex@latest

ENTRYPOINT ["./entry.sh"]
CMD ["make watch"]


## Prod
FROM alpine:latest AS prod

WORKDIR /

RUN apk add --no-cache build-base ccls

COPY --from=build /app/server /app/entry.sh /app/.env  /

ENTRYPOINT ["/entry.sh"]
CMD ["./server"]