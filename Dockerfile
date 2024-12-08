FROM ubuntu:latest
LABEL authors="Lach"

ENTRYPOINT ["top", "-b"]

FROM golang:1.21.6
WORKDIR /app
COPY /templates ./templates
COPY /static ./static
COPY main.go go.mod go.sum ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /testApp

RUN chmod +x /testApp

EXPOSE 8080

CMD ["/testApp"]