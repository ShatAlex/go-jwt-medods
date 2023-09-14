FROM golang:1.18.1

RUN go version
ENV GOPATH=/

COPY . .

RUN go mod download
RUN go build -o go-jwt ./cmd/main.go

CMD ["./go-jwt"]