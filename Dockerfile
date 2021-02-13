FROM alpine

WORKDIR /app

RUN apk add -U --no-cache ca-certificates

COPY asm-env /app

CMD ["/app/asm-env"]