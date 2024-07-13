FROM golang:1.22.5-alpine AS build

WORKDIR /app

COPY . /app

RUN go build -o jackboxtv-room-finder

FROM busybox:1.36.1 AS run

# need these to make HTTPS calls in busybox
COPY --from=build /etc/ssl/certs /etc/ssl/certs

COPY --from=build /app/jackboxtv-room-finder /usr/bin/jackboxtv-room-finder
RUN chmod +x /usr/bin/jackboxtv-room-finder

CMD ["/usr/bin/jackboxtv-room-finder"]
