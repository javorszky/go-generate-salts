FROM golang:1.19.3-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /salt

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /salt /salt

EXPOSE 8090

USER nonroot:nonroot

ENTRYPOINT ["/salt"]
