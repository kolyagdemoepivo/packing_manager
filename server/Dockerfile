FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go test -v ./...

RUN go build -o server ./cmd

EXPOSE 8080

CMD ["./server"]