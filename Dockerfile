FROM golang:1.14 AS builder

WORKDIR /src/webhook-fanout

COPY go.mod .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
    go build -o webhook-fanout cmd/webhook-fanout/main.go

FROM alpine:latest 

COPY --from=builder /src/webhook-fanout/webhook-fanout /bin/webhook-fanout

USER 65534
ENTRYPOINT ["/bin/webhook-fanout"]
