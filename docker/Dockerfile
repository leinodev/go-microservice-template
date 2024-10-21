FROM golang:1.23.2 as builder

ADD ./go.mod go.sum /modules/
RUN cd /modules && go mod download \
    && mkdir -p /service

ADD . /service
WORKDIR /service

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o ./bin/api ./cmd/all-in-one/main.go

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /service/bin/api /app/service

ADD https://github.com/mikefarah/yq/releases/download/v4.25.1/yq_linux_amd64 /bin/yq

COPY ./config/config.yaml ./config/

CMD ["./service"]