FROM golang:1.24.5

WORKDIR /app
COPY ./src .
RUN go mod tidy
RUN go build -o app .

CMD ["./app"]