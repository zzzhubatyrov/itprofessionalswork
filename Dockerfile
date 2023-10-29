FROM golang:latest

WORKDIR /app

COPY . .

# RUN go mod download

RUN go get -u

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

EXPOSE 5000

CMD ["/main"]