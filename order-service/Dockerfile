FROM golang:1.22

WORKDIR /APP

COPY . .

RUN go mod init order-service && go mod tidy

EXPOSE 8081

CMD ["go", "run", "main.go"]
