FROM golang:1.15.3-alpine AS build_base

ENV CGO_ENABLED=1
ENV GO111MODULE=on
RUN apk add --no-cache git  git gcc g++

WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./out/app ./cmd/main.go

FROM alpine:3.12
COPY --from=build_base /src/out/app /app
ENTRYPOINT ["/app"]
