FROM golang:1.16-alpine
WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -o /go-fibre-es7
EXPOSE 5000
CMD ["/go-fibre-es7" ]