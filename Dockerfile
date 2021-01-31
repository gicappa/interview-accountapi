FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v github.com/gicappa/interview-accountapi/cmd/client_example

#CMD ["client_example"]
ENTRYPOINT ["go", "test", "-v", "./...", "-coverprofile", "cover.out"]