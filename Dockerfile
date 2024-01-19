FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build -o ./bin/server ./cmd/main.go

CMD [ "./bin/server" ]