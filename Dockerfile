#build stage
FROM golang:latest
ENV GOROOT=/usr/local/go
ENV PATH=$PATH:/usr/local/go/bin
WORKDIR /api
EXPOSE 8080
# CMD go run main.go
