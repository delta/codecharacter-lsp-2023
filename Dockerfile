## Build
FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod go.sum  ./
RUN go mod download

COPY . .

RUN go build -o server

## Dev
FROM build AS dev

RUN apk add --no-cache build-base ccls py3-lsp-server curl openjdk17 py-pip

RUN pip install pyflakes rope --ignore-installed

WORKDIR /jdt
RUN curl https://download.eclipse.org/jdtls/milestones/1.9.0/jdt-language-server-1.9.0-202203031534.tar.gz --output jdt.tar.gz
RUN tar -xvf ./jdt.tar.gz

WORKDIR /app

RUN apk add --no-cache make

RUN go install github.com/cespare/reflex@latest

ENTRYPOINT ["./entry.sh"]
CMD ["make watch"]


## Prod
FROM alpine:latest AS prod

RUN apk add --no-cache build-base ccls py3-lsp-server curl openjdk17 py-pip

RUN pip install pyflakes rope --ignore-installed

WORKDIR /jdt
RUN curl https://download.eclipse.org/jdtls/milestones/1.9.0/jdt-language-server-1.9.0-202203031534.tar.gz --output jdt.tar.gz
RUN tar -xvf ./jdt.tar.gz

WORKDIR /

COPY --from=build /app/server /app/entry.sh /app/.env ./
COPY --from=build /app/player_code ./player_code/

ENTRYPOINT ["/entry.sh"]
CMD ["./server"]