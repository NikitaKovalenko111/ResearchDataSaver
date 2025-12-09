FROM golang:ubuntu

WORKDIR /server

COPY go.mod .
COPY go.sum .

RUN go install

COPY . .

RUN go build -o server ./cmd/main.go

EXPOSE 3001

CMD ["./server"]