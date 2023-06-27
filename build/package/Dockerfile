FROM golang:alpine AS builder

RUN apk add git

WORKDIR /app

COPY go.sum go.mod  ./

RUN go mod download

COPY . .

RUN go build -o /simple-ci

FROM alpine:lastes AS runtime

LABEL maintainer="Zahra Ehghaghi <z.ehghaghi@gmail.com>"

WORKDIR /app


COPY --from=builder /simple-ci  .

ENTRYPOINT [ "./simple-ci" ]

