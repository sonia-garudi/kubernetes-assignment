FROM golang:latest
RUN mkdir /go-webapp
WORKDIR /go-webapp
COPY . .
RUN chmod 777 /go-webapp
RUN ls
RUN go get -u github.com/go-sql-driver/mysql
RUN go build template.go
RUN pwd
RUN ls -al /root
RUN ls -al /go-webapp
ENV PORT 8080
CMD ["./template"]