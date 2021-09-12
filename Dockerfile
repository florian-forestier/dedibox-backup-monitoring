FROM golang:1.17-alpine as builder

WORKDIR /
COPY ./src ./src
ENV GO111MODULE=off
RUN go build -o dedibackup_poller ./src


FROM alpine:latest
COPY --from=builder /dedibackup_poller .

ENTRYPOINT "/dedibackup_poller"
EXPOSE 9101