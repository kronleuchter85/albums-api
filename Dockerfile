FROM golang:1.21.5
RUN mkdir /app
ADD src /app/
WORKDIR /app
RUN go mod download
RUN go build -o main
CMD ["/app/main"]