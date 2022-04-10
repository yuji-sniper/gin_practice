FROM golang:latest

WORKDIR /app

RUN go mod init app/main && go mod tidy

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]
