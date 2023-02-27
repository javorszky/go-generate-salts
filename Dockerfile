# syntax=docker/dockerfile:1

## Build
FROM golang:1.20-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /salt

## Deploy
FROM gcr.io/distroless/static-debian11

WORKDIR /

COPY --from=build /salt /salt

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/salt"]
