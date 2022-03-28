FROM golang:1.18-alpine AS builder
RUN mkdir /work
WORKDIR /work
COPY . /work
RUN go build main.go

FROM alpine:latest
COPY --from=builder /work/main .
EXPOSE 5000
ENTRYPOINT ./main

