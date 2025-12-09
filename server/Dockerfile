FROM golang:latest

WORKDIR /server

COPY . .

RUN go build -o server ./cmd/main.go

EXPOSE 3001

CMD ["./server"]