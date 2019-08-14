FROM golang:alpine as builder
COPY . /data
WORKDIR /data
RUN go build -o tcproxy tcproxy.go

FROM alpine:latest
COPY --from=builder /data/tcproxy /tcproxy
ENTRYPOINT [ "/tcproxy" ]