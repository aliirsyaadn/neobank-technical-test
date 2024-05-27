FROM golang:1.19-alpine3.18 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o main .

FROM alpine:3.18 AS production
RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app .

CMD ["./main"]
