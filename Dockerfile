FROM alpine

WORKDIR /app

RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY asm-env /app

CMD ["/app/asm-env"]