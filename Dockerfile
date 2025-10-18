FROM golang:1.25

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download
RUN apt update -y
RUN apt install ffmpeg -y

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/main.go

CMD ["app"]

