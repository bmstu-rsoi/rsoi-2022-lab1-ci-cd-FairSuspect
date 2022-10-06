# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY src/go.mod /
COPY src/go.sum /
RUN go mod download

COPY /src/ /

RUN go build -o build ./...

EXPOSE 8080
ENV PORT=8080

CMD [ "/build" ]