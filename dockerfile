FROM golang:1.20.6-bullseye AS builder
RUN apt-get update  
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64\
    connectionString=root:bebes2023*(192.168.1.65:3306)/directorio
WORKDIR /go/src/github.com/gherbust/lab
COPY go.mod .
RUN go mod download
COPY . .
RUN go build ../lab/app/main.go

FROM scratch
COPY --from=builder /go/src/github.com/gherbust/lab .
ENTRYPOINT ["./main"]