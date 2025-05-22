FROM mtr.devops.telekom.de/dpm/internal_tools/golang:latest
WORKDIR /app

COPY main.go /app/main.go
COPY helloworld.html /app/helloworld.html

RUN go mod init mytestserver
RUN go mod tidy
RUN go build -o /app/main /app/main.go

EXPOSE 8080
CMD ["/app/main"]