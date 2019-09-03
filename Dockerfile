FROM golang:1.12

LABEL maintainer="clooooode<jackey8616@gmail.com>"

EXPOSE 8000

WORKDIR /app

COPY . /app

RUN go mod download

ENTRYPOINT ["go", "run", "main.go"]

