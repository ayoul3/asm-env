FROM alpine

WORKDIR /app

COPY asm-env /app

CMD ["/app/asm-env"]