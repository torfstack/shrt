FROM golang:1.23 AS builder

RUN mkdir /opt/shrt
WORKDIR /opt/shrt

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY pkg pkg

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/shrt-server cmd/server/main.go

FROM alpine:3.19.1

RUN mkdir /opt/shrt
WORKDIR /opt/shrt

COPY --from=builder /bin/shrt-server shrt-server

CMD ["./shrt-server"]
