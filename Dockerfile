FROM golang:1.19.3-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /salt

FROM alpine

COPY --from=builder /salt /
WORKDIR /
EXPOSE 8090
ENTRYPOINT ["/salt"]
