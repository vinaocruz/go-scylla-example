FROM golang:latest

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go install github.com/air-verse/air@latest

RUN go mod download

CMD ["air", "-c", ".air.toml"]
