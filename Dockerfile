FROM golang:1.19-alpine as builder

WORKDIR /
COPY . .
RUN go build -o dedibackup_poller github.com/Artheriom/dedibox-backup-monitoring/cmd


FROM alpine:latest
COPY --from=builder /dedibackup_poller .

ENTRYPOINT "/dedibackup_poller"
EXPOSE 9101