FROM golang:1.22

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod download

RUN go build -o main .

CMD ["air"]
