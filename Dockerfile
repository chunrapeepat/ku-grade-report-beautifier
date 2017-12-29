FROM golang:1.8-alpine
ENV APP /go/src/github.com/chunza2542/ku-grade-report-beautifier
RUN mkdir -p $APP
COPY . $APP
WORKDIR $APP
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main src/main.go
CMD ["./main"]
