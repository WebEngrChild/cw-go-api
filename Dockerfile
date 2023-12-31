FROM golang:1.20-alpine as builder

RUN apk --no-cache add gcc musl-dev

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -trimpath -ldflags "-w -s" -o ./main main.go

FROM alpine:3.17 as runner

EXPOSE 8080

RUN apk update && \
    apk add --no-cache shadow && \
    useradd -m appuser && \
    rm -f /usr/bin/gpasswd /usr/bin/passwd /usr/bin/chfn /sbin/unix_chkpwd /usr/bin/expiry /usr/bin/chage /usr/bin/chsh && \
    rm -rf /var/cache/apk/*

USER appuser

WORKDIR /app
COPY --from=builder /app/main .

CMD ["./main"]