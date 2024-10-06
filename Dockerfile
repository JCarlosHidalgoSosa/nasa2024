FROM golang:1.23.2-alpine3.20

WORKDIR /app

COPY . .

RUN go build .

EXPOSE 8000

ENTRYPOINT ["go","run","."]
